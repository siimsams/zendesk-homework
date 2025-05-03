package service

import (
	"testing"

	"github.com/siimsams/zendesk-homework/database"
)

func TestCalculateOverallScore(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []database.Rating
		expected float64
	}{
		{
			name:     "empty ratings should return zero",
			ratings:  []database.Rating{},
			expected: 0,
		},
		{
			name: "single rating should return normalized score",
			ratings: []database.Rating{
				{Rating: 4, Weight: 1.0},
			},
			expected: 80.0,
		},
		{
			name: "multiple ratings with different weights",
			ratings: []database.Rating{
				{Rating: 5, Weight: 0.6},
				{Rating: 3, Weight: 0.4},
			},
			expected: 84.0,
		},
		{
			name: "perfect ratings should return 100",
			ratings: []database.Rating{
				{Rating: 5, Weight: 0.5},
				{Rating: 5, Weight: 0.5},
			},
			expected: 100.0,
		},
		{
			name: "zero ratings should return zero",
			ratings: []database.Rating{
				{Rating: 0, Weight: 1.0},
			},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateOverallScore(tt.ratings)

			if result != tt.expected {
				t.Errorf("calculateOverallScore() = %.2f, want %.2f", result, tt.expected)
			}
		})
	}
}
