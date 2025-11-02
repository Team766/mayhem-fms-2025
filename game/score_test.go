// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreSummary(t *testing.T) {
	redScore := TestScore1()
	blueScore := TestScore2()

	redSummary := redScore.Summarize(blueScore)
	assert.Equal(t, 12, redSummary.LeavePoints)
	assert.Equal(t, 68, redSummary.AutoPoints)
	assert.Equal(t, 6, redSummary.ParkPoints)
	assert.Equal(t, 30, redSummary.KrakenLairPoints)
	assert.Equal(t, 144, redSummary.MatchPoints)
	assert.Equal(t, 0, redSummary.FoulPoints)
	assert.Equal(t, 144, redSummary.Score)
	assert.Equal(t, true, redSummary.AutonRankingPoint)
	assert.Equal(t, true, redSummary.ScoringRankingPoint)
	assert.Equal(t, true, redSummary.EndgameRankingPoint)
	assert.Equal(t, 3, redSummary.BonusRankingPoints)

	blueSummary := blueScore.Summarize(redScore)
	assert.Equal(t, 12, blueSummary.LeavePoints)
	assert.Equal(t, 62, blueSummary.AutoPoints)
	assert.Equal(t, 9, blueSummary.ParkPoints)
	assert.Equal(t, 40, blueSummary.KrakenLairPoints)
	assert.Equal(t, 171, blueSummary.MatchPoints) // 30 (auto) + 10 (kraken) + 5 (park)
	assert.Equal(t, 60, blueSummary.FoulPoints)
	assert.Equal(t, 231, blueSummary.Score) // 45 (match) + 34 (fouls)
	assert.Equal(t, true, blueSummary.AutonRankingPoint)
	assert.Equal(t, true, blueSummary.ScoringRankingPoint)
	assert.Equal(t, true, blueSummary.EndgameRankingPoint)
	assert.Equal(t, 3, blueSummary.BonusRankingPoints)
}

func TestScoreEquals(t *testing.T) {
	score1 := TestScore1()
	score2 := TestScore1()
	assert.True(t, score1.Equals(score2))
	assert.True(t, score2.Equals(score1))

	score3 := TestScore2()
	assert.False(t, score1.Equals(score3))
	assert.False(t, score3.Equals(score1))

	score2 = TestScore1()
	score2.Mayhem.LeaveStatuses[0] = false
	assert.False(t, score1.Equals(score2))

	score2 = TestScore1()
	score2.Mayhem.AutoHullCount++
	assert.False(t, score1.Equals(score2))

	score2 = TestScore1()
	score2.Mayhem.ParkStatuses[0] = false
	assert.False(t, score1.Equals(score2))

	score2 = TestScore1()
	score2.Fouls = []Foul{}
	assert.False(t, score1.Equals(score2))
}

func TestAutonRankingPoint(t *testing.T) {
	score := Score{
		Mayhem: Mayhem{
			AutoHullCount: 0,
			AutoDeckCount: 2,
		},
	}
	summary := score.Summarize(&Score{})
	assert.True(t, summary.AutonRankingPoint)

	score = Score{
		Mayhem: Mayhem{
			AutoHullCount: 1,
			AutoDeckCount: 1,
		},
	}
	summary = score.Summarize(&Score{})
	assert.False(t, summary.AutonRankingPoint)
}

func TestScoringRankingPoint(t *testing.T) {
	score := Score{Mayhem: Mayhem{TeleopHullCount: 10, TeleopDeckCount: 3}}
	summary := score.Summarize(&Score{})
	assert.False(t, summary.ScoringRankingPoint)

	score = Score{Mayhem: Mayhem{TeleopHullCount: 10, TeleopDeckCount: 4}}
	summary = score.Summarize(&Score{})
	assert.True(t, summary.ScoringRankingPoint)
}

func TestEndgameRankingPoint(t *testing.T) {
	score := Score{Mayhem: Mayhem{EndgameKrakenLairCount: 2}}
	summary := score.Summarize(&Score{})
	assert.False(t, summary.EndgameRankingPoint)

	score = Score{Mayhem: Mayhem{EndgameKrakenLairCount: 3}}
	summary = score.Summarize(&Score{})
	assert.True(t, summary.EndgameRankingPoint)
}
