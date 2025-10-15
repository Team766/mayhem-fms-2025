// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Helper methods for use in tests in this package and others.

package game

func TestScore1() *Score {
	fouls := []Foul{
		{true, 25, 16},
		{false, 1868, 13},
		{false, 1868, 13},
		{true, 25, 15},
		{true, 25, 15},
		{true, 25, 15},
		{true, 25, 15},
	}
	return &Score{
		RobotsBypassed: [3]bool{false, false, true},
		Mayhem: Mayhem{
			AutoHullCount:          2,  // Increased from 1 to 2 for more auto points
			TeleopHullCount:        10, // Increased from 2 to 10 to trigger Scoring RP (10 + 4 = 14)
			AutoDeckCount:          3,  // Increased from 2 to 3 for more auto points
			TeleopDeckCount:        4,
			EndgameKrakenLairCount: 3,                         // Increased from 2 to 3 to trigger Endgame RP
			LeaveStatuses:          [3]bool{true, true, true}, // All robots left to maximize auto points
			MusterStatuses:         [3]bool{true, true, true}, // All robots mustered to maximize auto points
			ParkStatuses:           [3]bool{true, true, false},
		},
		Fouls:     fouls,
		PlayoffDq: false,
	}
}

func TestScore2() *Score {
	return &Score{
		RobotsBypassed: [3]bool{false, false, false},
		Mayhem: Mayhem{
			AutoHullCount:          3,                         // Increased from 2 to 3 for more auto points
			TeleopHullCount:        5,                         // Increased from 3 to 5
			AutoDeckCount:          2,                         // Increased from 1 to 2 for more auto points
			TeleopDeckCount:        10,                        // Increased from 5 to 10 to trigger Scoring RP (5 + 10 = 15)
			EndgameKrakenLairCount: 4,                         // Increased from 1 to 4 to trigger Endgame RP
			LeaveStatuses:          [3]bool{true, true, true}, // All robots left to maximize auto points
			MusterStatuses:         [3]bool{true, true, true}, // All robots mustered to maximize auto points
			ParkStatuses:           [3]bool{true, true, true},
		},
		Fouls:     []Foul{},
		PlayoffDq: false,
	}
}

func TestRanking1() *Ranking {
	return &Ranking{TeamId: 254, Rank: 1, PreviousRank: 0, RankingFields: RankingFields{RankingPoints: 20, MatchPoints: 625, AutoPoints: 90, EndgameKrakenLairPoints: 40, Wins: 3, Losses: 2, Ties: 1, Disqualifications: 0, Played: 10}}
}

func TestRanking2() *Ranking {
	return &Ranking{TeamId: 1114, Rank: 2, PreviousRank: 1, RankingFields: RankingFields{RankingPoints: 18, MatchPoints: 700, AutoPoints: 100, EndgameKrakenLairPoints: 50, Wins: 1, Losses: 3, Ties: 2, Disqualifications: 0, Played: 10}}
}
