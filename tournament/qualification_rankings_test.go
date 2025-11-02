// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package tournament

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Team254/cheesy-arena/game"
	"github.com/Team254/cheesy-arena/model"
	"github.com/stretchr/testify/assert"
)

func TestCalculateRankings(t *testing.T) {
	rand.Seed(1)
	database := setupTestDb(t)

	setupMatchResultsForRankings(database)
	updatedRankings, err := CalculateRankings(database, false)
	assert.Nil(t, err)
	rankings, err := database.GetAllRankings()
	assert.Nil(t, err)
	assert.Equal(t, updatedRankings, rankings)
	if assert.Equal(t, 6, len(rankings)) {
		// Team 4 has the highest average ranking points (10 points / 2 matches = 5.0 per match)
		assert.Equal(t, 4, rankings[0].TeamId)
		assert.Equal(t, 0, rankings[0].PreviousRank)
		// Team 5 has more total ranking points (11) but lower average (11/3 ≈ 3.67 per match)
		assert.Equal(t, 5, rankings[1].TeamId)
		assert.Equal(t, 0, rankings[1].PreviousRank)
		// Team 6 has 7 ranking points (7/2 = 3.5 per match)
		assert.Equal(t, 6, rankings[2].TeamId)
		assert.Equal(t, 0, rankings[2].PreviousRank)
		// Team 1 has 8 ranking points but lower average (8/3 ≈ 2.67 per match)
		assert.Equal(t, 1, rankings[3].TeamId)
		assert.Equal(t, 0, rankings[3].PreviousRank)
		// Team 3 has 4 ranking points (4/2 = 2.0 per match)
		assert.Equal(t, 3, rankings[4].TeamId)
		assert.Equal(t, 0, rankings[4].PreviousRank)
		// Team 2 has 5 ranking points but 1 DQ (disqualifications affect ranking position)
		assert.Equal(t, 2, rankings[5].TeamId)
		assert.Equal(t, 0, rankings[5].PreviousRank)
	}

	previousRankings := make(map[int]int)
	for _, ranking := range rankings {
		fmt.Printf("%+v\n", ranking)
		previousRankings[ranking.TeamId] = ranking.Rank
	}
	fmt.Println()

	// Test after changing a match result.
	matchResult3 := model.BuildTestMatchResult(3, 3)
	matchResult3.RedScore, matchResult3.BlueScore = matchResult3.BlueScore, matchResult3.RedScore
	err = database.CreateMatchResult(matchResult3)
	assert.Nil(t, err)
	updatedRankings, err = CalculateRankings(database, false)
	assert.Nil(t, err)
	rankings, err = database.GetAllRankings()
	assert.Nil(t, err)
	assert.Equal(t, updatedRankings, rankings)
	if assert.Equal(t, 6, len(rankings)) {
		// Team 6 moves up to first with highest average (12 points / 2 matches = 6.0 per match)
		assert.Equal(t, 6, rankings[0].TeamId)
		assert.Equal(t, previousRankings[rankings[0].TeamId], rankings[0].PreviousRank)
		// Team 5 has more total points (16) but lower average (16/3 ≈ 5.33 per match)
		assert.Equal(t, 5, rankings[1].TeamId)
		assert.Equal(t, previousRankings[rankings[1].TeamId], rankings[1].PreviousRank)
		// Team 4 has 10 ranking points (10/2 = 5.0 per match)
		assert.Equal(t, 4, rankings[2].TeamId)
		assert.Equal(t, previousRankings[rankings[2].TeamId], rankings[2].PreviousRank)
		// Team 1 has 10 ranking points but lower average (10/3 ≈ 3.33 per match)
		assert.Equal(t, 1, rankings[3].TeamId)
		assert.Equal(t, previousRankings[rankings[3].TeamId], rankings[3].PreviousRank)
		// Team 3 has 6 ranking points (6/2 = 3.0 per match)
		assert.Equal(t, 3, rankings[4].TeamId)
		assert.Equal(t, previousRankings[rankings[4].TeamId], rankings[4].PreviousRank)
		// Team 2 has 7 ranking points but 1 DQ (disqualifications affect ranking position)
		assert.Equal(t, 2, rankings[5].TeamId)
		assert.Equal(t, previousRankings[rankings[5].TeamId], rankings[5].PreviousRank)
	}

	for _, ranking := range rankings {
		fmt.Printf("%+v\n", ranking)
	}
	fmt.Println()

	matchResult3 = model.BuildTestMatchResult(3, 4)
	err = database.CreateMatchResult(matchResult3)
	assert.Nil(t, err)
	updatedRankings, err = CalculateRankings(database, true)
	assert.Nil(t, err)
	rankings, err = database.GetAllRankings()
	assert.Nil(t, err)
	assert.Equal(t, updatedRankings, rankings)
	if assert.Equal(t, 6, len(rankings)) {
		// Team 4 has the highest average (10 points / 2 matches = 5.0 per match)
		assert.Equal(t, 4, rankings[0].TeamId)
		assert.Equal(t, previousRankings[rankings[0].TeamId], rankings[0].PreviousRank)
		// Team 3 has 9 ranking points (9/2 = 4.5 per match)
		assert.Equal(t, 3, rankings[1].TeamId)
		assert.Equal(t, previousRankings[rankings[1].TeamId], rankings[1].PreviousRank)
		// Team 6 also has 9 ranking points but lower tiebreaker (9/2 = 4.5 per match)
		assert.Equal(t, 6, rankings[2].TeamId)
		assert.Equal(t, previousRankings[rankings[2].TeamId], rankings[2].PreviousRank)
		// Team 5 has more total points (13) but lower average (13/3 ≈ 4.33 per match)
		assert.Equal(t, 5, rankings[3].TeamId)
		assert.Equal(t, previousRankings[rankings[3].TeamId], rankings[3].PreviousRank)
		// Team 1 also has 13 points but lower average and tiebreakers (13/3 ≈ 4.33 per match)
		assert.Equal(t, 1, rankings[4].TeamId)
		assert.Equal(t, previousRankings[rankings[4].TeamId], rankings[4].PreviousRank)
		// Team 2 has 10 ranking points but 1 DQ (disqualifications affect ranking position)
		assert.Equal(t, 2, rankings[5].TeamId)
		assert.Equal(t, previousRankings[rankings[5].TeamId], rankings[5].PreviousRank)
	}

	for _, ranking := range rankings {
		fmt.Printf("%+v\n", ranking)
	}
	fmt.Println()

}

func TestAddMatchResultToRankingsHandleCards(t *testing.T) {
	rankings := map[int]*game.Ranking{}
	matchResult := model.BuildTestMatchResult(1, 1)
	matchResult.RedCards = map[string]string{"1": "yellow", "2": "red", "3": "dq"}
	matchResult.BlueCards = map[string]string{"4": "red", "5": "dq", "6": "yellow"}
	addMatchResultToRankings(rankings, 1, matchResult, true)
	addMatchResultToRankings(rankings, 2, matchResult, true)
	addMatchResultToRankings(rankings, 3, matchResult, true)
	addMatchResultToRankings(rankings, 4, matchResult, false)
	addMatchResultToRankings(rankings, 5, matchResult, false)
	addMatchResultToRankings(rankings, 6, matchResult, false)
	assert.Equal(t, 0, rankings[1].Disqualifications)
	assert.Equal(t, 1, rankings[2].Disqualifications)
	assert.Equal(t, 1, rankings[3].Disqualifications)
	assert.Equal(t, 1, rankings[4].Disqualifications)
	assert.Equal(t, 1, rankings[5].Disqualifications)
	assert.Equal(t, 0, rankings[6].Disqualifications)
}

// Sets up a schedule and results that touches on all possible variables.
func setupMatchResultsForRankings(database *model.Database) {
	match1 := model.Match{
		Type:      model.Qualification,
		TypeOrder: 1,
		Red1:      1,
		Red2:      2,
		Red3:      3,
		Blue1:     4,
		Blue2:     5,
		Blue3:     6,
		Status:    game.RedWonMatch,
	}
	database.CreateMatch(&match1)
	matchResult1 := model.BuildTestMatchResult(match1.Id, 1)
	matchResult1.RedCards = map[string]string{"2": "red"}
	database.CreateMatchResult(matchResult1)

	match2 := model.Match{
		Type:             model.Qualification,
		TypeOrder:        2,
		Red1:             1,
		Red2:             3,
		Red3:             5,
		Blue1:            2,
		Blue2:            4,
		Blue3:            6,
		Status:           game.BlueWonMatch,
		Red2IsSurrogate:  true,
		Blue3IsSurrogate: true,
	}
	database.CreateMatch(&match2)
	matchResult2 := model.BuildTestMatchResult(match2.Id, 1)
	matchResult2.BlueScore = matchResult2.RedScore
	database.CreateMatchResult(matchResult2)

	match3 := model.Match{
		Type:            model.Qualification,
		TypeOrder:       3,
		Red1:            6,
		Red2:            5,
		Red3:            4,
		Blue1:           3,
		Blue2:           2,
		Blue3:           1,
		Status:          game.TieMatch,
		Red3IsSurrogate: true,
	}
	database.CreateMatch(&match3)
	matchResult3 := model.BuildTestMatchResult(match3.Id, 1)
	database.CreateMatchResult(matchResult3)
	matchResult3 = model.NewMatchResult()
	matchResult3.MatchId = match3.Id
	matchResult3.PlayNumber = 2
	database.CreateMatchResult(matchResult3)

	match4 := model.Match{
		Type:      model.Practice,
		TypeOrder: 1,
		Red1:      1,
		Red2:      2,
		Red3:      3,
		Blue1:     4,
		Blue2:     5,
		Blue3:     6,
		Status:    game.RedWonMatch,
	}
	database.CreateMatch(&match4)
	matchResult4 := model.BuildTestMatchResult(match4.Id, 1)
	database.CreateMatchResult(matchResult4)

	match5 := model.Match{
		Type:      model.Playoff,
		TypeOrder: 8,
		Red1:      1,
		Red2:      2,
		Red3:      3,
		Blue1:     4,
		Blue2:     5,
		Blue3:     6,
		Status:    game.BlueWonMatch,
	}
	database.CreateMatch(&match5)
	matchResult5 := model.BuildTestMatchResult(match5.Id, 1)
	database.CreateMatchResult(matchResult5)

	match6 := model.Match{
		Type:      model.Qualification,
		TypeOrder: 4,
		Red1:      7,
		Red2:      8,
		Red3:      9,
		Blue1:     10,
		Blue2:     11,
		Blue3:     12,
		Status:    game.MatchScheduled,
	}
	database.CreateMatch(&match6)
	matchResult6 := model.BuildTestMatchResult(match6.Id, 1)
	database.CreateMatchResult(matchResult6)
}
