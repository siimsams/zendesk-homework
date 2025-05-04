package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/siimsams/zendesk-homework/env"
	"github.com/siimsams/zendesk-homework/observability/logging"
	scorer "github.com/siimsams/zendesk-homework/proto"
	"github.com/siimsams/zendesk-homework/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	grpcClient scorer.ScorerServiceClient
	grpcConn   *grpc.ClientConn
)

func TestMain(m *testing.M) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logging.SetLogLevel("debug")

	go func() {
		config := env.Config{
			Port:     "50051",
			DbPath:   "database.db",
			LogLevel: "debug",
		}
		if err := startServer(config); err != nil {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	time.Sleep(2 * time.Second)

	var err error
	grpcConn, err = grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to gRPC server")
	}
	grpcClient = scorer.NewScorerServiceClient(grpcConn)

	code := m.Run()

	grpcConn.Close()
	os.Exit(code)
}

func withAuth(ctx context.Context, t *testing.T) context.Context {
	md := metadata.New(map[string]string{
		"authorization": test.GenerateValidToken(t),
	})
	return metadata.NewOutgoingContext(ctx, md)
}

func TestGetCategoryScores(t *testing.T) {
	ctx := withAuth(context.Background(), t)

	req := &scorer.ScoreRequest{
		StartDate: "2019-04-01",
		EndDate:   "2019-04-30",
	}

	resp, err := grpcClient.GetCategoryScores(ctx, req)
	require.NoError(t, err, "GetCategoryScores failed")
	require.NotNil(t, resp, "Response should not be nil")

	// Define expected data for each category
	expectedCategories := map[string]struct {
		overallScore float64
		ratingCount  int32
		dailyScores  map[string]float64
	}{
		"Spelling": {
			overallScore: 50.99386503067485,
			ratingCount:  815,
			dailyScores: map[string]float64{
				"2019-04-01": 36.666666666666664,
				"2019-04-02": 39.35483870967742,
				"2019-04-03": 48.275862068965516,
				"2019-04-04": 57.77777777777777,
				"2019-04-05": 51.66666666666667,
				"2019-04-06": 56.00000000000001,
				"2019-04-07": 46.89655172413793,
				"2019-04-08": 51.891891891891895,
				"2019-04-09": 53.68421052631579,
				"2019-04-10": 50.66666666666667,
				"2019-04-11": 54.074074074074076,
				"2019-04-12": 60.76923076923077,
				"2019-04-13": 49.09090909090909,
				"2019-04-14": 52.307692307692314,
				"2019-04-15": 49.333333333333336,
				"2019-04-16": 44.666666666666664,
				"2019-04-17": 40,
				"2019-04-18": 65.18518518518519,
				"2019-04-19": 40.714285714285715,
				"2019-04-20": 49.6551724137931,
				"2019-04-21": 61.66666666666667,
				"2019-04-22": 62.42424242424243,
				"2019-04-23": 55.86206896551724,
				"2019-04-24": 52.142857142857146,
				"2019-04-25": 48.75,
				"2019-04-26": 58.94736842105262,
				"2019-04-27": 50,
				"2019-04-28": 49.473684210526315,
				"2019-04-29": 50.4,
			},
		},
		"Grammar": {
			overallScore: 49.32515337423313,
			ratingCount:  815,
			dailyScores: map[string]float64{
				"2019-04-01": 50,
				"2019-04-02": 53.5483870967742,
				"2019-04-03": 42.06896551724138,
				"2019-04-04": 45.18518518518518,
				"2019-04-05": 50,
				"2019-04-06": 54,
				"2019-04-07": 46.89655172413793,
				"2019-04-08": 45.40540540540541,
				"2019-04-09": 60,
				"2019-04-10": 52.44444444444445,
				"2019-04-11": 54.81481481481482,
				"2019-04-12": 53.84615384615385,
				"2019-04-13": 48.484848484848484,
				"2019-04-14": 53.84615384615385,
				"2019-04-15": 46.666666666666664,
				"2019-04-16": 55.333333333333336,
				"2019-04-17": 38.37837837837838,
				"2019-04-18": 59.25925925925925,
				"2019-04-19": 37.857142857142854,
				"2019-04-20": 53.103448275862064,
				"2019-04-21": 51.66666666666667,
				"2019-04-22": 50.90909090909091,
				"2019-04-23": 55.172413793103445,
				"2019-04-24": 46.42857142857143,
				"2019-04-25": 55.625,
				"2019-04-26": 46.31578947368421,
				"2019-04-27": 48.57142857142857,
				"2019-04-28": 29.47368421052631,
				"2019-04-29": 44.800000000000004,
			},
		},
		"GDPR": {
			overallScore: 49.61963190184049,
			ratingCount:  815,
			dailyScores: map[string]float64{
				"2019-04-01": 33.33333333333333,
				"2019-04-02": 52.25806451612903,
				"2019-04-03": 46.206896551724135,
				"2019-04-04": 51.11111111111111,
				"2019-04-05": 55.833333333333336,
				"2019-04-06": 48,
				"2019-04-07": 54.48275862068965,
				"2019-04-08": 63.78378378378379,
				"2019-04-09": 41.05263157894737,
				"2019-04-10": 52.44444444444445,
				"2019-04-11": 45.925925925925924,
				"2019-04-12": 52.307692307692314,
				"2019-04-13": 42.42424242424242,
				"2019-04-14": 57.692307692307686,
				"2019-04-15": 43.333333333333336,
				"2019-04-16": 47.333333333333336,
				"2019-04-17": 61.62162162162163,
				"2019-04-18": 48.888888888888886,
				"2019-04-19": 49.28571428571429,
				"2019-04-20": 46.89655172413793,
				"2019-04-21": 50.83333333333333,
				"2019-04-22": 47.27272727272727,
				"2019-04-23": 53.79310344827586,
				"2019-04-24": 42.857142857142854,
				"2019-04-25": 50,
				"2019-04-26": 47.368421052631575,
				"2019-04-27": 44.285714285714285,
				"2019-04-28": 45.26315789473684,
				"2019-04-29": 48.8,
			},
		},
		"Randomness": {
			overallScore: 52.147239263803684,
			ratingCount:  815,
			dailyScores: map[string]float64{
				"2019-04-01": 58.333333333333336,
				"2019-04-02": 49.03225806451613,
				"2019-04-03": 50.3448275862069,
				"2019-04-04": 59.25925925925925,
				"2019-04-05": 64.16666666666667,
				"2019-04-06": 60,
				"2019-04-07": 51.724137931034484,
				"2019-04-08": 48.10810810810811,
				"2019-04-09": 62.10526315789474,
				"2019-04-10": 49.333333333333336,
				"2019-04-11": 41.48148148148148,
				"2019-04-12": 48.46153846153846,
				"2019-04-13": 53.333333333333336,
				"2019-04-14": 40.76923076923077,
				"2019-04-15": 62.66666666666667,
				"2019-04-16": 54,
				"2019-04-17": 53.51351351351351,
				"2019-04-18": 48.888888888888886,
				"2019-04-19": 48.57142857142857,
				"2019-04-20": 53.103448275862064,
				"2019-04-21": 56.666666666666664,
				"2019-04-22": 48.484848484848484,
				"2019-04-23": 53.103448275862064,
				"2019-04-24": 56.42857142857143,
				"2019-04-25": 59.375,
				"2019-04-26": 48.421052631578945,
				"2019-04-27": 49.28571428571429,
				"2019-04-28": 46.31578947368421,
				"2019-04-29": 41.6,
			},
		},
	}

	assert.Len(t, resp.Categories, 4, "Expected 4 categories")

	for _, category := range resp.Categories {
		expected, exists := expectedCategories[category.Category]
		assert.True(t, exists, "Unexpected category found: %s", category.Category)
		if !exists {
			continue
		}

		assert.Equal(t, expected.overallScore, category.OverallScore,
			"Category %s: overall score mismatch", category.Category)

		assert.Equal(t, expected.ratingCount, category.RatingCount,
			"Category %s: rating count mismatch", category.Category)

		assert.Len(t, category.DateToScore, len(expected.dailyScores),
			"Category %s: number of daily scores mismatch", category.Category)

		for _, dailyScore := range category.DateToScore {
			expectedScore, exists := expected.dailyScores[dailyScore.Date]
			assert.True(t, exists, "Category %s: unexpected date found: %s",
				category.Category, dailyScore.Date)
			if exists {
				assert.Equal(t, expectedScore, dailyScore.Score,
					"Category %s, date %s: score mismatch",
					category.Category, dailyScore.Date)
			}
		}

		for i := 1; i < len(category.DateToScore); i++ {
			assert.True(t, category.DateToScore[i].Date > category.DateToScore[i-1].Date,
				"Category %s: dates are not in chronological order at index %d",
				category.Category, i)
		}
	}
}

func TestGetCategoryScoresByTicket(t *testing.T) {
	ctx := withAuth(context.Background(), t)

	req := &scorer.ScoreRequest{
		StartDate: "2019-04-01",
		EndDate:   "2019-04-05",
	}

	resp, err := grpcClient.GetCategoryScoresByTicket(ctx, req)
	require.NoError(t, err, "GetCategoryScoresByTicket failed")
	require.NotNil(t, resp, "Response should not be nil")

	// Define expected ticket scores
	expectedTickets := map[int64]map[string]float64{
		16583:  {"1": 60, "2": 40, "3": 80, "4": 0},
		28534:  {"1": 60, "2": 40, "3": 20, "4": 0},
		37917:  {"1": 80, "2": 80, "3": 100, "4": 0},
		42067:  {"1": 80, "2": 20, "3": 20, "4": 0},
		54540:  {"1": 60, "2": 20, "3": 100, "4": 0},
		61414:  {"1": 0, "2": 80, "3": 0, "4": 0},
		70851:  {"1": 80, "2": 20, "3": 20, "4": 0},
		72749:  {"1": 100, "2": 0, "3": 20, "4": 0},
		75202:  {"1": 100, "2": 0, "3": 0, "4": 0},
		96239:  {"1": 80, "2": 40, "3": 60, "4": 0},
		98232:  {"1": 100, "2": 59.999999999999986, "3": 40, "4": 0},
		109840: {"1": 100, "2": 80, "3": 0, "4": 0},
		128263: {"1": 40, "2": 100, "3": 100, "4": 0},
		132627: {"1": 80, "2": 40, "3": 60, "4": 0},
		136115: {"1": 60, "2": 40, "3": 100, "4": 0},
		150753: {"1": 20, "2": 20, "3": 40, "4": 0},
		150759: {"1": 80, "2": 0, "3": 40, "4": 0},
		155633: {"1": 40, "2": 59.999999999999986, "3": 0, "4": 0},
		157231: {"1": 0, "2": 59.999999999999986, "3": 100, "4": 0},
		158241: {"1": 60, "2": 20, "3": 20, "4": 0},
		181750: {"1": 0, "2": 20, "3": 80, "4": 0},
		187046: {"1": 0, "2": 0, "3": 0, "4": 0},
		188110: {"1": 40, "2": 20, "3": 100, "4": 0},
		199261: {"1": 100, "2": 0, "3": 80, "4": 0},
		205172: {"1": 60, "2": 100, "3": 20, "4": 0},
		218630: {"1": 40, "2": 59.999999999999986, "3": 60, "4": 0},
		255100: {"1": 80, "2": 80, "3": 80, "4": 0},
		262112: {"1": 80, "2": 80, "3": 40, "4": 0},
		274207: {"1": 0, "2": 40, "3": 60, "4": 0},
		283292: {"1": 60, "2": 40, "3": 0, "4": 0},
		287833: {"1": 40, "2": 0, "3": 80, "4": 0},
		302558: {"1": 80, "2": 40, "3": 20, "4": 0},
		303866: {"1": 100, "2": 80, "3": 20, "4": 0},
		315557: {"1": 60, "2": 0, "3": 20, "4": 0},
		319558: {"1": 60, "2": 20, "3": 40, "4": 0},
		321828: {"1": 80, "2": 59.999999999999986, "3": 20, "4": 0},
		332498: {"1": 80, "2": 59.999999999999986, "3": 80, "4": 0},
		333078: {"1": 40, "2": 40, "3": 80, "4": 0},
		348189: {"1": 80, "2": 100, "3": 20, "4": 0},
		349323: {"1": 100, "2": 100, "3": 20, "4": 0},
		364474: {"1": 0, "2": 40, "3": 80, "4": 0},
		379849: {"1": 20, "2": 80, "3": 0, "4": 0},
		394576: {"1": 20, "2": 59.999999999999986, "3": 0, "4": 0},
		404903: {"1": 20, "2": 100, "3": 0, "4": 0},
		411688: {"1": 100, "2": 0, "3": 20, "4": 0},
		418569: {"1": 0, "2": 40, "3": 80, "4": 0},
		423312: {"1": 40, "2": 80, "3": 100, "4": 0},
		430749: {"1": 40, "2": 59.999999999999986, "3": 60, "4": 0},
		431415: {"1": 0, "2": 20, "3": 40, "4": 0},
		444907: {"1": 20, "2": 100, "3": 0, "4": 0},
		461744: {"1": 60, "2": 40, "3": 80, "4": 0},
		463977: {"1": 100, "2": 40, "3": 40, "4": 0},
		465199: {"1": 80, "2": 100, "3": 0, "4": 0},
		488530: {"1": 20, "2": 100, "3": 60, "4": 0},
		493004: {"1": 80, "2": 40, "3": 0, "4": 0},
		505809: {"1": 40, "2": 59.999999999999986, "3": 100, "4": 0},
		510233: {"1": 60, "2": 20, "3": 80, "4": 0},
		529071: {"1": 20, "2": 59.999999999999986, "3": 80, "4": 0},
		532423: {"1": 0, "2": 0, "3": 40, "4": 0},
		542922: {"1": 80, "2": 100, "3": 100, "4": 0},
		562513: {"1": 60, "2": 0, "3": 100, "4": 0},
		576390: {"1": 100, "2": 59.999999999999986, "3": 60, "4": 0},
		580495: {"1": 0, "2": 40, "3": 60, "4": 0},
		584182: {"1": 40, "2": 59.999999999999986, "3": 0, "4": 0},
		600526: {"1": 100, "2": 59.999999999999986, "3": 40, "4": 0},
		639033: {"1": 60, "2": 80, "3": 0, "4": 0},
		639625: {"1": 60, "2": 100, "3": 0, "4": 0},
		643682: {"1": 80, "2": 100, "3": 0, "4": 0},
		644068: {"1": 80, "2": 80, "3": 0, "4": 0},
		668572: {"1": 40, "2": 40, "3": 80, "4": 0},
		673316: {"1": 80, "2": 59.999999999999986, "3": 0, "4": 0},
		675540: {"1": 20, "2": 40, "3": 100, "4": 0},
		675666: {"1": 80, "2": 0, "3": 0, "4": 0},
		678163: {"1": 20, "2": 40, "3": 100, "4": 0},
		678982: {"1": 40, "2": 90, "3": 40, "4": 0},
		681828: {"1": 60, "2": 40, "3": 100, "4": 0},
		682933: {"1": 20, "2": 59.999999999999986, "3": 40, "4": 0},
		686573: {"1": 80, "2": 100, "3": 60, "4": 0},
		694207: {"1": 100, "2": 80, "3": 0, "4": 0},
		701675: {"1": 40, "2": 20, "3": 80, "4": 0},
		721553: {"1": 80, "2": 59.999999999999986, "3": 80, "4": 0},
		739407: {"1": 40, "2": 59.999999999999986, "3": 0, "4": 0},
		741503: {"1": 0, "2": 100, "3": 60, "4": 0},
		756517: {"1": 80, "2": 40, "3": 0, "4": 0},
		764513: {"1": 20, "2": 40, "3": 20, "4": 0},
		773577: {"1": 100, "2": 100, "3": 60, "4": 0},
		781980: {"1": 20, "2": 59.999999999999986, "3": 20, "4": 0},
		798054: {"1": 60, "2": 80, "3": 80, "4": 0},
		811150: {"1": 60, "2": 20, "3": 0, "4": 0},
		837069: {"1": 40, "2": 59.999999999999986, "3": 80, "4": 0},
		854919: {"1": 0, "2": 20, "3": 0, "4": 0},
		889209: {"1": 80, "2": 0, "3": 100, "4": 0},
		901290: {"1": 60, "2": 20, "3": 60, "4": 0},
		902961: {"1": 100, "2": 0, "3": 80, "4": 0},
		920045: {"1": 0, "2": 80, "3": 60, "4": 0},
		929019: {"1": 80, "2": 100, "3": 60, "4": 0},
		937246: {"1": 20, "2": 20, "3": 40, "4": 0},
		979882: {"1": 40, "2": 80, "3": 40, "4": 0},
	}

	assert.Len(t, resp.TicketScores, len(expectedTickets), "Number of tickets mismatch")

	for _, ticketScore := range resp.TicketScores {
		expectedScores, exists := expectedTickets[ticketScore.TicketId]
		assert.True(t, exists, "Unexpected ticket ID found: %d", ticketScore.TicketId)
		if !exists {
			continue
		}

		assert.Len(t, ticketScore.CategoryScores, len(expectedScores),
			"Ticket %d: number of category scores mismatch", ticketScore.TicketId)

		for category, score := range ticketScore.CategoryScores {
			expectedScore, exists := expectedScores[category]
			assert.True(t, exists, "Ticket %d: unexpected category found: %s",
				ticketScore.TicketId, category)
			if exists {
				assert.Equal(t, expectedScore, score,
					"Ticket %d, category %s: score mismatch",
					ticketScore.TicketId, category)
			}
		}
	}
}

func TestGetOverallScore(t *testing.T) {
	ctx := withAuth(context.Background(), t)

	req := &scorer.ScoreRequest{
		StartDate: "2019-04-01",
		EndDate:   "2019-04-02",
	}

	resp, err := grpcClient.GetOverallScore(ctx, req)
	require.NoError(t, err, "GetOverallScore failed")
	require.NotNil(t, resp, "Response should not be nil")

	expectedScore := 38.50574712643679
	assert.Equal(t, expectedScore, resp.ScorePercentage,
		"Overall score percentage mismatch. Expected: %f, Got: %f",
		expectedScore, resp.ScorePercentage)
}

func TestGetPeriodOverPeriodChange(t *testing.T) {
	ctx := withAuth(context.Background(), t)

	req := &scorer.ScoreRequest{
		StartDate: "2019-04-01",
		EndDate:   "2019-04-02",
	}

	resp, err := grpcClient.GetPeriodOverPeriodChange(ctx, req)
	require.NoError(t, err, "GetPeriodOverPeriodChange failed")
	require.NotNil(t, resp, "Response should not be nil")

	assert.Equal(t, "2019-03-31T00:00:00Z", resp.PreviousPeriodStart,
		"Previous period start date mismatch")
	assert.Equal(t, "2019-04-01T00:00:00Z", resp.PreviousPeriodEnd,
		"Previous period end date mismatch")
	assert.Equal(t, 46.73563218390804, resp.PreviousPeriodScore,
		"Previous period score mismatch")
	assert.Equal(t, "2019-04-01T00:00:00Z", resp.CurrentPeriodStart,
		"Current period start date mismatch")
	assert.Equal(t, "2019-04-02T00:00:00Z", resp.CurrentPeriodEnd,
		"Current period end date mismatch")
	assert.Equal(t, 38.50574712643679, resp.CurrentPeriodScore,
		"Current period score mismatch")
	assert.Equal(t, -17.609444171175575, resp.ChangePercentage,
		"Change percentage mismatch")
}
