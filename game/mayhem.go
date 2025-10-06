// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Game specific constants, set up as placeholders that can be easily customized.
package game

type Mayhem struct {
	AutoHullCount          int
	TeleopHullCount        int
	AutoDeckCount          int
	TeleopDeckCount        int
	EndgameKrakenLairCount int
	LeaveStatuses          [3]bool
	MusterStatuses         [3]bool
	ParkStatuses           [3]bool
}

const (
	LeavePoints             = 4
	MusterPoints            = 6
	ParkPoints              = 3
	AutoHullPoints          = 2 * TeleopHullPoints
	TeleopHullPoints        = 2
	AutoDeckPoints          = 2 * TeleopDeckPoints
	TeleopDeckPoints        = 5
	EndgameKrakenLairPoints = 10

	Gamepiece1RPThreshold = 8
	MinorFoulPoints       = 5
	MajorFoulPoints       = 10
)
