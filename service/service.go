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

	score, err := database.CalculateOverallScore(sqliteDB, start, end)
	if err != nil {
		log.Printf("score calculation failed: %v", err)
		return nil, err
	}

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

	scores, err := database.CalculateScoresByTicket(sqliteDB, start, end)
	if err != nil {
		log.Printf("score calculation failed: %v", err)
		return nil, err
	}

	return &scorer.ScoreByTicketResponse{
		TicketScores: scores,
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

	scores, err := database.CalculateAggregatedCategoryScores(sqliteDB, start, end)
	if err != nil {
		log.Printf("score calculation failed: %v", err)
		return nil, err
	}

	return &scorer.AggregatedCategoryScoresResponse{
		Categories: scores,
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

	change, err := database.CalculatePeriodOverPeriodScoreChange(sqliteDB, start, end)
	if err != nil {
		log.Printf("score calculation failed: %v", err)
		return nil, err
	}

	return &scorer.PeriodOverPeriodScoreChangeResponse{
		PreviousPeriodScore: change.PreviousPeriodScore,
		CurrentPeriodScore:  change.CurrentPeriodScore,
		ChangePercentage:    change.ChangePercentage,
	}, nil
}
