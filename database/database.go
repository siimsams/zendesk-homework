package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Rating struct {
	Rating int
	Weight float64
}

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
