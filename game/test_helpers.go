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
			AutoHullCount:          1,
			TeleopHullCount:        2,
			AutoDeckCount:          2,
			TeleopDeckCount:        4,
			EndgameKrakenLairCount: 2,
			LeaveStatuses:          [3]bool{true, true, false},
			MusterStatuses:         [3]bool{false, true, false},
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
			AutoHullCount:          2,
			TeleopHullCount:        3,
			AutoDeckCount:          1,
			TeleopDeckCount:        5,
			EndgameKrakenLairCount: 1,
			LeaveStatuses:          [3]bool{false, true, false},
			MusterStatuses:         [3]bool{false, true, false},
			ParkStatuses:           [3]bool{false, true, false},
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
