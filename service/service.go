package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

func calculateOverallScore(ratings []database.Rating) float64 {
	var totalWeightedScore float64
	var totalWeight float64

	for _, rating := range ratings {
		score := (float64(rating.Rating) / 5.0) * rating.Weight
		totalWeightedScore += score
		totalWeight += rating.Weight
	}

	if totalWeight == 0 {
		return 0
	}

	return (totalWeightedScore / totalWeight) * 100.0
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

	ratings, err := database.GetRatings(sqliteDB, start, end)
	if err != nil {
		log.Printf("failed to get ratings: %v", err)
		return nil, err
	}

	score := calculateOverallScore(ratings)
	return &scorer.ScoreResponse{
		ScorePercentage: score,
	}, nil
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

	ticketRatings, err := database.GetTicketRatings(sqliteDB, start, end)
	if err != nil {
		log.Printf("failed to get ticket ratings: %v", err)
		return nil, err
	}

	ticketScores := []*scorer.TicketScore{}
	for _, tr := range ticketRatings {
		score := (float64(tr.Rating) / 5.0) * tr.Weight
		ticketIDInt64, err := strconv.ParseInt(tr.TicketID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid ticket ID: %w", err)
		}
		ticketScores = append(ticketScores, &scorer.TicketScore{
			TicketId: ticketIDInt64,
			CategoryScores: map[string]float64{
				"Overall": score,
			},
		})
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
		log.Printf("failed to open DB: %v", err)
		return nil, err
	}
	defer sqliteDB.Close()

	currentRatings, err := database.GetRatings(sqliteDB, start, end)
	if err != nil {
		log.Printf("failed to get current ratings: %v", err)
		return nil, err
	}

	periodDuration := end.Sub(start)
	previousStart := start.Add(-periodDuration)
	previousEnd := end.Add(-periodDuration)

	previousRatings, err := database.GetRatings(sqliteDB, previousStart, previousEnd)
	if err != nil {
		log.Printf("failed to get previous ratings: %v", err)
		return nil, err
	}

	currentScore := calculateOverallScore(currentRatings)
	previousScore := calculateOverallScore(previousRatings)

	var change float64
	if previousScore != 0 {
		change = ((currentScore - previousScore) / previousScore) * 100
	} else if currentScore != 0 {
		change = 100
	} else {
		change = 0
	}

	return &scorer.PeriodOverPeriodScoreChangeResponse{
		CurrentPeriodScore:  currentScore,
		PreviousPeriodScore: previousScore,
		ChangePercentage:    change,
	}, nil
}
