package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "modernc.org/sqlite"

	scorer "github.com/siimsams/zendesk-homework/proto"
)

func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}

func CalculateOverallScore(db *sql.DB, start, end time.Time) (float64, error) {
	rows, err := db.Query(`
		SELECT r.rating, rc.weight
		FROM ratings r
		JOIN rating_categories rc ON r.rating_category_id = rc.id
		WHERE r.created_at BETWEEN ? AND ?
	`, start.Format(time.RFC3339), end.Format(time.RFC3339))
	if err != nil {
		return 0, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var totalWeightedScore float64
	var totalWeight float64

	for rows.Next() {
		var rating int
		var weight float64
		if err := rows.Scan(&rating, &weight); err != nil {
			return 0, err
		}

		score := (float64(rating) / 5.0) * weight
		totalWeightedScore += score
		totalWeight += weight
	}

	if totalWeight == 0 {
		return 0, nil
	}

	return (totalWeightedScore / totalWeight) * 100.0, nil
}

func CalculateScoresByTicket(db *sql.DB, start, end time.Time) ([]*scorer.TicketScore, error) {
	rows, err := db.Query(`
		SELECT t.id, t.subject, r.rating, rc.weight
		FROM tickets t
		JOIN ratings r ON t.id = r.ticket_id
		JOIN rating_categories rc ON r.rating_category_id = rc.id
		WHERE t.created_at BETWEEN ? AND ?
	`, start.Format(time.RFC3339), end.Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	ticketScores := []*scorer.TicketScore{}

	for rows.Next() {
		var ticketID string
		var subject string
		var rating int
		var weight float64
		if err := rows.Scan(&ticketID, &subject, &rating, &weight); err != nil {
			return nil, err
		}

		score := (float64(rating) / 5.0) * weight
		ticketIDInt64, err := strconv.ParseInt(ticketID, 10, 64)
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

	return ticketScores, nil
}

func CalculateAggregatedCategoryScores(db *sql.DB, start, end time.Time) ([]*scorer.CategoryScore, error) {
	periodLength := end.Sub(start)

	var groupBy string
	if periodLength.Hours() > 24*31 {
		groupBy = "strftime('%Y-W%W', r.created_at)"
	} else {
		groupBy = "date(r.created_at)"
	}

	query := fmt.Sprintf(`
		SELECT
			rc.name,
			rc.weight,
			%s as period,
			COUNT(*) as count,
			SUM(r.rating) as total_rating
		FROM ratings r
		JOIN rating_categories rc ON r.rating_category_id = rc.id
		WHERE r.created_at BETWEEN ? AND ?
		GROUP BY rc.id, period
		ORDER BY rc.name, period
	`, groupBy)

	rows, err := db.Query(query, start.Format(time.RFC3339), end.Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	type tempAgg struct {
		weight       float64
		totalRatings int
		totalScore   float64
		scoresByDate map[string]float64
	}

	result := map[string]*tempAgg{}

	for rows.Next() {
		var name, period string
		var count, sumRating int
		var weight float64
		if err := rows.Scan(&name, &weight, &period, &count, &sumRating); err != nil {
			return nil, err
		}

		weightedScore := (float64(sumRating) / float64(count*5)) * 100

		if _, exists := result[name]; !exists {
			result[name] = &tempAgg{
				weight:       weight,
				totalRatings: 0,
				totalScore:   0,
				scoresByDate: make(map[string]float64),
			}
		}

		result[name].totalRatings += count
		result[name].totalScore += weightedScore * float64(count)
		result[name].scoresByDate[period] = weightedScore
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

	return output, nil
}

func CalculatePeriodOverPeriodScoreChange(db *sql.DB, start, end time.Time) (*scorer.PeriodOverPeriodScoreChangeResponse, error) {
	currentScore, err := CalculateOverallScore(db, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate current score: %w", err)
	}

	periodDuration := end.Sub(start)
	previousStart := start.Add(-periodDuration)
	previousEnd := end.Add(-periodDuration)

	previousScore, err := CalculateOverallScore(db, previousStart, previousEnd)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate previous score: %w", err)
	}

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
