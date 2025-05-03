package service

import (
	"context"
	"log"
	"time"

	"github.com/siimsams/zendesk-homework/database"
	scorer "github.com/siimsams/zendesk-homework/proto"
)

type ScorerServer struct {
	scorer.UnimplementedScorerServiceServer
	DBPath string
}

func parseDateRange(req *scorer.ScoreRequest) (time.Time, time.Time, error) {
	start, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	end, err := time.Parse(time.DateOnly, req.EndDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return start, end, nil
}

func (s *ScorerServer) GetOverallScore(ctx context.Context, req *scorer.ScoreRequest) (*scorer.ScoreResponse, error) {
	start, end, err := parseDateRange(req)
	if err != nil {
		return nil, err
	}

	sqliteDB, err := database.OpenDB(s.DBPath)
	if err != nil {
		log.Printf("failed to open DB: %v", err)
		return nil, err
	}
	defer sqliteDB.Close()

	score, err := database.GetCombinedScore(sqliteDB, start, end)
	if err != nil {
		log.Printf("failed to get score: %v", err)
		return nil, err
	}

	return &scorer.ScoreResponse{
		ScorePercentage: score,
	}, nil
}

func groupTicketScores(ticketRatingsForEachCategory []database.TicketRating) ([]*scorer.TicketScore, error) {
	ticketScores := make([]*scorer.TicketScore, 0)
	ticketScoresMap := make(map[int64]map[string]float64)

	for _, tr := range ticketRatingsForEachCategory {
		if err := processSingleTicketRating(tr, ticketScoresMap, &ticketScores); err != nil {
			return nil, err
		}
	}

	return ticketScores, nil
}

func processSingleTicketRating(tr database.TicketRating,
	ticketScoresMap map[int64]map[string]float64,
	ticketScores *[]*scorer.TicketScore) error {

	if _, exists := ticketScoresMap[tr.TicketID]; exists {
		ticketScoresMap[tr.TicketID][tr.RatingCategoryName] = tr.CategoryScorePercent
		return nil
	}

	ticketScoresMap[tr.TicketID] = make(map[string]float64)
	*ticketScores = append(*ticketScores, &scorer.TicketScore{
		TicketId:       tr.TicketID,
		CategoryScores: ticketScoresMap[tr.TicketID],
	})
	ticketScoresMap[tr.TicketID][tr.RatingCategoryName] = tr.CategoryScorePercent
	return nil
}

func (s *ScorerServer) GetScoresByTicket(ctx context.Context, req *scorer.ScoreRequest) (*scorer.ScoreByTicketResponse, error) {
	start, end, err := parseDateRange(req)
	if err != nil {
		return nil, err
	}

	sqliteDB, err := database.OpenDB(s.DBPath)
	if err != nil {
		log.Printf("failed to open DB: %v", err)
		return nil, err
	}
	defer sqliteDB.Close()

	ticketRatingsForEachCategory, err := database.GetTicketRatingForEachCategory(sqliteDB, start, end)
	if err != nil {
		log.Printf("failed to get ticket ratings: %v", err)
		return nil, err
	}

	ticketScores, err := groupTicketScores(ticketRatingsForEachCategory)
	if err != nil {
		return nil, err
	}

	return &scorer.ScoreByTicketResponse{
		TicketScores: ticketScores,
	}, nil
}

func (s *ScorerServer) GetAggregatedCategoryScores(ctx context.Context, req *scorer.ScoreRequest) (*scorer.AggregatedCategoryScoresResponse, error) {
	start, end, err := parseDateRange(req)
	if err != nil {
		return nil, err
	}

	sqliteDB, err := database.OpenDB(s.DBPath)
	if err != nil {
		log.Printf("failed to open DB: %v", err)
		return nil, err
	}
	defer sqliteDB.Close()

	aggregations, err := database.GetCategoryAggregations(sqliteDB, start, end)
	if err != nil {
		log.Printf("failed to get category aggregations: %v", err)
		return nil, err
	}

	type tempAgg struct {
		weight       float64
		totalRatings int
		totalScore   float64
		scoresByDate map[string]float64
	}

	result := map[string]*tempAgg{}

	for _, agg := range aggregations {
		weightedScore := (float64(agg.TotalRating) / float64(agg.Count*5)) * 100

		if _, exists := result[agg.Name]; !exists {
			result[agg.Name] = &tempAgg{
				weight:       agg.Weight,
				totalRatings: 0,
				totalScore:   0,
				scoresByDate: make(map[string]float64),
			}
		}

		result[agg.Name].totalRatings += agg.Count
		result[agg.Name].totalScore += weightedScore * float64(agg.Count)
		result[agg.Name].scoresByDate[agg.Period] = weightedScore
	}

	var output []*scorer.CategoryScore
	for name, agg := range result {
		var avgScore float64
		if agg.totalRatings > 0 {
			avgScore = agg.totalScore / float64(agg.totalRatings)
		}
		output = append(output, &scorer.CategoryScore{
			Category:     name,
			RatingCount:  int32(agg.totalRatings),
			DateToScore:  agg.scoresByDate,
			OverallScore: avgScore,
		})
	}

	return &scorer.AggregatedCategoryScoresResponse{
		Categories: output,
	}, nil
}

func (s *ScorerServer) GetPeriodOverPeriodScoreChange(ctx context.Context, req *scorer.ScoreRequest) (*scorer.PeriodOverPeriodScoreChangeResponse, error) {
	start, end, err := parseDateRange(req)
	if err != nil {
		return nil, err
	}

	sqliteDB, err := database.OpenDB(s.DBPath)
	if err != nil {
		return nil, err
	}
	defer sqliteDB.Close()

	periodDuration := end.Sub(start)
	previousStart := start.Add(-periodDuration)
	previousEnd := end.Add(-periodDuration)

	currentScore, err := database.GetCombinedScore(sqliteDB, start, end)
	if err != nil {
		return nil, err
	}

	previousScore, err := database.GetCombinedScore(sqliteDB, previousStart, previousEnd)
	if err != nil {
		return nil, err
	}

	change := 0.0
	if previousScore != 0 {
		change = ((currentScore - previousScore) / previousScore) * 100
	} else if currentScore != 0 {
		change = 100
	}

	return &scorer.PeriodOverPeriodScoreChangeResponse{
		CurrentPeriodScore:  currentScore,
		CurrentPeriodStart:  start.Format(time.RFC3339),
		CurrentPeriodEnd:    end.Format(time.RFC3339),
		PreviousPeriodScore: previousScore,
		PreviousPeriodStart: previousStart.Format(time.RFC3339),
		PreviousPeriodEnd:   previousEnd.Format(time.RFC3339),
		ChangePercentage:    change,
	}, nil
}
