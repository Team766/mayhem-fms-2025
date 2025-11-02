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
			AutoHullCount:          2,                          // 2 * 4 = 8
			TeleopHullCount:        10,                         // 10 * 2 = 20
			AutoDeckCount:          3,                          // 3 * 10 = 30
			TeleopDeckCount:        4,                          // 4 * 5 = 20
			EndgameKrakenLairCount: 3,                          // 3 * 10 = 30
			LeaveStatuses:          [3]bool{true, true, true},  // 3 * 4 = 12
			MusterStatuses:         [3]bool{true, true, true},  // 3 * 6 = 18
			ParkStatuses:           [3]bool{true, true, false}, // 2 * 3 = 6
		}, // Auton Points: 68 (RP), Scoring Count: 14 (RP), Endgame Count: 3 (RP)
		Fouls:     fouls, // (2 minor * 5) + 5 major * 10 = 60
		PlayoffDq: false,
	}
}

func TestScore2() *Score {
	return &Score{
		RobotsBypassed: [3]bool{false, false, false},
		Mayhem: Mayhem{
			AutoHullCount:          3,                         // 3 * 4 = 12
			TeleopHullCount:        5,                         // 5 * 2 = 10
			AutoDeckCount:          2,                         // 2 * 10 = 20
			TeleopDeckCount:        10,                        // 10 * 5 = 50
			EndgameKrakenLairCount: 4,                         // 4 * 10 = 40
			LeaveStatuses:          [3]bool{true, true, true}, // 3 * 4 = 12
			MusterStatuses:         [3]bool{true, true, true}, // 3 * 6 = 18
			ParkStatuses:           [3]bool{true, true, true}, // 3 * 3 = 9
		}, // Auton Points: 62 (RP), Scoring Count: 15 (RP), Endgame Count: 4 (RP)
		Fouls:     []Foul{}, // 0
		PlayoffDq: false,
	}
}

func TestRanking1() *Ranking {
	return &Ranking{TeamId: 254, Rank: 1, PreviousRank: 0, RankingFields: RankingFields{RankingPoints: 20, MatchPoints: 625, AutoPoints: 90, EndgameKrakenLairPoints: 40, Wins: 3, Losses: 2, Ties: 1, Disqualifications: 0, Played: 10}}
}

func TestRanking2() *Ranking {
	return &Ranking{TeamId: 1114, Rank: 2, PreviousRank: 1, RankingFields: RankingFields{RankingPoints: 18, MatchPoints: 700, AutoPoints: 100, EndgameKrakenLairPoints: 50, Wins: 1, Losses: 3, Ties: 2, Disqualifications: 0, Played: 10}}
}
