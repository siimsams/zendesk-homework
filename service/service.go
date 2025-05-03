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

func (s *ScorerServer) GetOverallScore(ctx context.Context, req *scorer.ScoreRequest) (*scorer.ScoreResponse, error) {
	start, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(time.DateOnly, req.EndDate)
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
