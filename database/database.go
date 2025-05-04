package database

import (
	"database/sql"
	"fmt"
	"time"

	"go.nhat.io/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	_ "modernc.org/sqlite"
)

func OpenDB(path string) (*sql.DB, error) {
	driverName, err := otelsql.Register("sqlite",
		otelsql.AllowRoot(),
		otelsql.TraceQueryWithoutArgs(),
		otelsql.TraceRowsClose(),
		otelsql.TraceRowsAffected(),
		otelsql.WithDatabaseName("my_database"),
		otelsql.WithSystem(semconv.DBSystemSqlite),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register otelsql driver: %w", err)
	}

	db, err := sql.Open(driverName, path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}

type Score struct {
	ScorePercent float64
}

func GetCombinedScore(db *sql.DB, start, end time.Time) (float64, error) {
	row := db.QueryRow(`
		SELECT 
			COALESCE(SUM(r.rating * rc.weight) / NULLIF(SUM(rc.weight), 0) / 5.0 * 100.0, 0)
		FROM ratings r
		JOIN rating_categories rc ON r.rating_category_id = rc.id
		WHERE r.created_at BETWEEN ? AND ?
	`, start, end)

	var score float64
	err := row.Scan(&score)
	if err != nil {
		return -1, fmt.Errorf("scan failed: %w", err)
	}

	return score, nil
}

type TicketRating struct {
	TicketID             int64
	RatingCategoryId     string
	CategoryScorePercent float64
}

func GetTicketRatingForEachCategory(db *sql.DB, start, end time.Time) ([]TicketRating, error) {
	rows, err := db.Query(`
		SELECT
			t.id AS ticket_id,
			rc.id AS category_id,
			COALESCE(SUM(r.rating * rc.weight) / NULLIF(SUM(rc.weight), 0) / 5.0 * 100.0, 0) AS category_score_percent
		FROM tickets t
		JOIN ratings r ON t.id = r.ticket_id
		JOIN rating_categories rc ON r.rating_category_id = rc.id
		WHERE t.created_at BETWEEN ? AND ?
		GROUP BY t.id, rc.id
	`, start, end)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var ticketRatings []TicketRating
	for rows.Next() {
		var tr TicketRating
		if err := rows.Scan(&tr.TicketID, &tr.RatingCategoryId, &tr.CategoryScorePercent); err != nil {
			return nil, err
		}
		ticketRatings = append(ticketRatings, tr)
	}
	return ticketRatings, nil
}

type CategoryAggregation struct {
	Name        string
	Weight      float64
	Period      string
	Count       int
	TotalRating int
}

func GetCategoryAggregations(db *sql.DB, start, end time.Time) ([]CategoryAggregation, error) {
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
		ORDER BY period
	`, groupBy)

	rows, err := db.Query(query, start, end)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var aggregations []CategoryAggregation
	for rows.Next() {
		var agg CategoryAggregation
		if err := rows.Scan(&agg.Name, &agg.Weight, &agg.Period, &agg.Count, &agg.TotalRating); err != nil {
			return nil, err
		}
		aggregations = append(aggregations, agg)
	}
	return aggregations, nil
}
