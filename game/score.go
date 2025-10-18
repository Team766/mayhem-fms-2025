// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	RobotsBypassed [3]bool
	Mayhem         Mayhem
	Fouls          []Foul
	PlayoffDq      bool
}

// Summarize calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize(opponentScore *Score) *ScoreSummary {
	summary := new(ScoreSummary)

	// Leave the score at zero if the alliance was disqualified.
	if score.PlayoffDq {
		return summary
	}

	// Calculate autonomous period points.
	for _, status := range score.Mayhem.LeaveStatuses {
		if status {
			summary.LeavePoints += LeavePoints
		}
	}

	for _, status := range score.Mayhem.MusterStatuses {
		if status {
			summary.MusterPoints += MusterPoints
		}
	}

	summary.AutoPoints = summary.LeavePoints + summary.MusterPoints +
		score.Mayhem.AutoHullCount*AutoHullPoints +
		score.Mayhem.AutoDeckCount*AutoDeckPoints

	summary.DeckPoints = score.Mayhem.AutoDeckCount*AutoDeckPoints +
		score.Mayhem.TeleopDeckCount*TeleopDeckPoints

	summary.TreasureShipPoints = score.Mayhem.AutoHullCount*AutoHullPoints +
		score.Mayhem.TeleopHullCount*TeleopHullPoints +
		summary.DeckPoints

	summary.KrakenLairPoints = score.Mayhem.EndgameKrakenLairCount * EndgameKrakenLairPoints

	// Calculate park points.
	for _, status := range score.Mayhem.ParkStatuses {
		if status {
			summary.ParkPoints += ParkPoints
		}
	}

	summary.MatchPoints = summary.LeavePoints + summary.MusterPoints + summary.TreasureShipPoints + summary.KrakenLairPoints + summary.ParkPoints

	// Calculate penalty points.
	for _, foul := range opponentScore.Fouls {
		summary.FoulPoints += foul.PointValue()
	}

	summary.Score = summary.MatchPoints + summary.FoulPoints

	// Calculate bonus ranking points.
	// Auton Ranking Point - score 20+ points during auton
	summary.AutonRankingPoint = summary.AutoPoints >= AutonRankingPointThreshold
	if summary.AutonRankingPoint {
		summary.BonusRankingPoints++
	}

	// Scoring Ranking Point - score 14+ cannonballs during teleop+endgame
	summary.ScoringRankingPoint = score.Mayhem.TeleopHullCount+score.Mayhem.TeleopDeckCount >= ScoringRankingPointThreshold
	if summary.ScoringRankingPoint {
		summary.BonusRankingPoints++
	}

	// Endgame Ranking Point - score 3+ Kraken Lair
	summary.EndgameRankingPoint = score.Mayhem.EndgameKrakenLairCount >= EndgameRankingPointThreshold
	if summary.EndgameRankingPoint {
		summary.BonusRankingPoints++
	}

	return summary
}

// Equals returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool {
	if score.Mayhem.LeaveStatuses != other.Mayhem.LeaveStatuses ||
		score.Mayhem.MusterStatuses != other.Mayhem.MusterStatuses ||
		score.Mayhem.AutoHullCount != other.Mayhem.AutoHullCount ||
		score.Mayhem.TeleopHullCount != other.Mayhem.TeleopHullCount ||
		score.Mayhem.AutoDeckCount != other.Mayhem.AutoDeckCount ||
		score.Mayhem.TeleopDeckCount != other.Mayhem.TeleopDeckCount ||
		score.Mayhem.EndgameKrakenLairCount != other.Mayhem.EndgameKrakenLairCount ||
		score.Mayhem.ParkStatuses != other.Mayhem.ParkStatuses ||
		score.RobotsBypassed != other.RobotsBypassed ||
		score.PlayoffDq != other.PlayoffDq ||
		len(score.Fouls) != len(other.Fouls) {
		return false
	}

	for i, foul := range score.Fouls {
		if foul != other.Fouls[i] {
			return false
		}
	}

	return true
}
