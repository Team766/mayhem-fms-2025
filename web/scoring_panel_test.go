// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package web

import (
	"testing"
	"time"

	"github.com/Team254/cheesy-arena/field"
	"github.com/Team254/cheesy-arena/websocket"
	gorillawebsocket "github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestScoringPanel(t *testing.T) {
	web := setupTestWeb(t)

	recorder := web.getHttpResponse("/panels/scoring/invalidposition")
	assert.Equal(t, 500, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Invalid position")
	recorder = web.getHttpResponse("/panels/scoring/red_near")
	assert.Equal(t, 200, recorder.Code)
	recorder = web.getHttpResponse("/panels/scoring/red_far")
	assert.Equal(t, 200, recorder.Code)
	recorder = web.getHttpResponse("/panels/scoring/blue_near")
	assert.Equal(t, 200, recorder.Code)
	recorder = web.getHttpResponse("/panels/scoring/blue_far")
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Scoring Panel - Untitled Event - Cheesy Arena")
}

func TestScoringPanelTwoVsTwoAttribute(t *testing.T) {
	web := setupTestWeb(t)

	// TwoVsTwoMode off
	web.arena.EventSettings.TwoVsTwoMode = false
	recorder := web.getHttpResponse("/panels/scoring/red_near")
	assert.Equal(t, 200, recorder.Code)
	body := recorder.Body.String()
	assert.Contains(t, body, "id=\"scoringPanel\"")
	assert.Contains(t, body, "data-two-v-two=\"false\"")

	// TwoVsTwoMode on
	web.arena.EventSettings.TwoVsTwoMode = true
	recorder = web.getHttpResponse("/panels/scoring/red_near")
	assert.Equal(t, 200, recorder.Code)
	body = recorder.Body.String()
	assert.Contains(t, body, "id=\"scoringPanel\"")
	assert.Contains(t, body, "data-two-v-two=\"true\"")
}

func TestScoringPanelWebsocket(t *testing.T) {
	web := setupTestWeb(t)

	server, wsUrl := web.startTestServer()
	defer server.Close()
	_, _, err := gorillawebsocket.DefaultDialer.Dial(wsUrl+"/panels/scoring/blorpy/websocket", nil)
	assert.NotNil(t, err)
	redConn, _, err := gorillawebsocket.DefaultDialer.Dial(wsUrl+"/panels/scoring/red_near/websocket", nil)
	assert.Nil(t, err)
	defer redConn.Close()
	redWs := websocket.NewTestWebsocket(redConn)
	assert.Equal(t, 1, web.arena.ScoringPanelRegistry.GetNumPanels("red_near"))
	assert.Equal(t, 0, web.arena.ScoringPanelRegistry.GetNumPanels("blue_near"))
	blueConn, _, err := gorillawebsocket.DefaultDialer.Dial(wsUrl+"/panels/scoring/blue_near/websocket", nil)
	assert.Nil(t, err)
	defer blueConn.Close()
	blueWs := websocket.NewTestWebsocket(blueConn)
	assert.Equal(t, 1, web.arena.ScoringPanelRegistry.GetNumPanels("red_near"))
	assert.Equal(t, 1, web.arena.ScoringPanelRegistry.GetNumPanels("blue_near"))

	// Should get a few status updates right after connection.
	readWebsocketType(t, redWs, "resetLocalState")
	readWebsocketType(t, redWs, "matchLoad")
	readWebsocketType(t, redWs, "matchTime")
	readWebsocketType(t, redWs, "realtimeScore")
	readWebsocketType(t, blueWs, "resetLocalState")
	readWebsocketType(t, blueWs, "matchLoad")
	readWebsocketType(t, blueWs, "matchTime")
	readWebsocketType(t, blueWs, "realtimeScore")

	// Send some autonomous period scoring commands.
	assert.Equal(t, [3]bool{false, false, false}, web.arena.RedRealtimeScore.CurrentScore.Mayhem.LeaveStatuses)
	leaveData := struct {
		TeamPosition int
	}{}
	web.arena.MatchState = field.AutoPeriod
	leaveData.TeamPosition = 1
	redWs.Write("leave", leaveData)
	leaveData.TeamPosition = 3
	redWs.Write("leave", leaveData)
	for i := 0; i < 2; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}
	assert.Equal(t, [3]bool{true, false, true}, web.arena.RedRealtimeScore.CurrentScore.Mayhem.LeaveStatuses)
	redWs.Write("leave", leaveData)
	readWebsocketType(t, redWs, "realtimeScore")
	readWebsocketType(t, blueWs, "realtimeScore")
	assert.Equal(t, [3]bool{true, false, false}, web.arena.RedRealtimeScore.CurrentScore.Mayhem.LeaveStatuses)

	leaveData.TeamPosition = 1
	redWs.Write("muster", leaveData)
	leaveData.TeamPosition = 3
	redWs.Write("muster", leaveData)
	for i := 0; i < 2; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}
	assert.Equal(t, [3]bool{true, false, true}, web.arena.RedRealtimeScore.CurrentScore.Mayhem.MusterStatuses)
	redWs.Write("muster", leaveData)
	readWebsocketType(t, redWs, "realtimeScore")
	readWebsocketType(t, blueWs, "realtimeScore")
	assert.Equal(t, [3]bool{true, false, false}, web.arena.RedRealtimeScore.CurrentScore.Mayhem.MusterStatuses)

	// Test hull and deck scoring commands
	hullData := struct {
		Autonomous bool
		Adjustment int
	}{}
	deckData := struct {
		Autonomous bool
		Adjustment int
	}{}

	// Initialize counts to 0
	assert.Equal(t, 0, web.arena.RedRealtimeScore.CurrentScore.Mayhem.TeleopHullCount)
	assert.Equal(t, 0, web.arena.BlueRealtimeScore.CurrentScore.Mayhem.TeleopHullCount)
	assert.Equal(t, 0, web.arena.RedRealtimeScore.CurrentScore.Mayhem.TeleopDeckCount)
	assert.Equal(t, 0, web.arena.BlueRealtimeScore.CurrentScore.Mayhem.TeleopDeckCount)

	// Test hull for Blue alliance (teleop)
	hullData.Autonomous = false
	hullData.Adjustment = 1
	blueWs.Write("hull", hullData)
	blueWs.Write("hull", hullData)
	blueWs.Write("hull", hullData)
	hullData.Adjustment = -1
	blueWs.Write("hull", hullData)
	blueWs.Write("hull", hullData)
	hullData.Adjustment = 1
	blueWs.Write("hull", hullData)
	for i := 0; i < 6; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}

	// Test deck for Red alliance (teleop)
	deckData.Autonomous = false
	deckData.Adjustment = 1
	redWs.Write("deck", deckData)
	redWs.Write("deck", deckData)
	deckData.Adjustment = -1
	redWs.Write("deck", deckData)
	redWs.Write("deck", deckData)
	redWs.Write("deck", deckData)
	for i := 0; i < 5; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}

	// Verify counts after initial commands
	assert.Equal(t, 0, web.arena.RedRealtimeScore.CurrentScore.Mayhem.TeleopHullCount)
	assert.Equal(t, 2, web.arena.BlueRealtimeScore.CurrentScore.Mayhem.TeleopHullCount)
	assert.Equal(t, 0, web.arena.RedRealtimeScore.CurrentScore.Mayhem.TeleopDeckCount)
	assert.Equal(t, 0, web.arena.BlueRealtimeScore.CurrentScore.Mayhem.TeleopDeckCount)

	// Test hull and deck for Red alliance (auto and teleop)
	// Auto hull
	hullData.Autonomous = true
	hullData.Adjustment = 1
	redWs.Write("hull", hullData)
	redWs.Write("hull", hullData)
	redWs.Write("hull", hullData)
	hullData.Adjustment = -1
	redWs.Write("hull", hullData)
	for i := 0; i < 4; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}

	// Teleop hull
	hullData.Autonomous = false
	hullData.Adjustment = 1
	redWs.Write("hull", hullData)
	redWs.Write("hull", hullData)

	// Auto deck
	deckData.Autonomous = true
	deckData.Adjustment = 1
	redWs.Write("deck", deckData)
	redWs.Write("deck", deckData)

	// Teleop deck (decrement)
	deckData.Autonomous = false
	deckData.Adjustment = -1
	redWs.Write("deck", deckData)
	for i := 0; i < 5; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}

	// Verify final counts
	assert.Equal(t, 2, web.arena.RedRealtimeScore.CurrentScore.Mayhem.TeleopHullCount)
	assert.Equal(t, 0, web.arena.RedRealtimeScore.CurrentScore.Mayhem.TeleopDeckCount)
	assert.Equal(t, 2, web.arena.RedRealtimeScore.CurrentScore.Mayhem.AutoHullCount)
	assert.Equal(t, 2, web.arena.RedRealtimeScore.CurrentScore.Mayhem.AutoDeckCount)

	// Test kraken lair scoring
	krakenData := struct {
		Adjustment int
	}{}

	// Test kraken lair for Red alliance
	krakenData.Adjustment = 1
	redWs.Write("kraken_lair", krakenData)
	redWs.Write("kraken_lair", krakenData)
	krakenData.Adjustment = -1
	redWs.Write("kraken_lair", krakenData)

	// Process kraken lair messages
	for i := 0; i < 3; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}

	// Verify kraken lair count
	assert.Equal(t, 1, web.arena.RedRealtimeScore.CurrentScore.Mayhem.EndgameKrakenLairCount)
	assert.Equal(t, 0, web.arena.BlueRealtimeScore.CurrentScore.Mayhem.EndgameKrakenLairCount)

	// Send some park status commands
	parkData := struct {
		TeamPosition int
	}{} // Note: ParkStatus field is not used in the implementation, it toggles based on position only
	assert.Equal(t, [3]bool{false, false, false}, web.arena.RedRealtimeScore.CurrentScore.Mayhem.ParkStatuses)
	assert.Equal(t, [3]bool{false, false, false}, web.arena.BlueRealtimeScore.CurrentScore.Mayhem.ParkStatuses)
	parkData.TeamPosition = 1
	redWs.Write("park", parkData)  // true
	blueWs.Write("park", parkData) // true
	parkData.TeamPosition = 2
	blueWs.Write("park", parkData) // true
	parkData.TeamPosition = 3
	blueWs.Write("park", parkData) // true
	redWs.Write("park", parkData)  // true
	redWs.Write("park", parkData)  // false
	parkData.TeamPosition = 2
	redWs.Write("park", parkData) // true
	for i := 0; i < 7; i++ {
		readWebsocketType(t, redWs, "realtimeScore")
		readWebsocketType(t, blueWs, "realtimeScore")
	}
	assert.Equal(
		t,
		[3]bool{true, true, false},
		web.arena.RedRealtimeScore.CurrentScore.Mayhem.ParkStatuses,
	)
	assert.Equal(
		t,
		[3]bool{true, true, true},
		web.arena.BlueRealtimeScore.CurrentScore.Mayhem.ParkStatuses,
	)

	// Test that some invalid commands do nothing and don't result in score change notifications.
	redWs.Write("invalid", nil)
	leaveData.TeamPosition = 0
	redWs.Write("leave", leaveData)

	// Test committing logic.
	redWs.Write("commitMatch", nil)
	readWebsocketType(t, redWs, "error")
	blueWs.Write("commitMatch", nil)
	readWebsocketType(t, blueWs, "error")
	assert.Equal(t, 0, web.arena.ScoringPanelRegistry.GetNumScoreCommitted("red_near"))
	assert.Equal(t, 0, web.arena.ScoringPanelRegistry.GetNumScoreCommitted("blue_near"))
	web.arena.MatchState = field.PostMatch
	redWs.Write("commitMatch", nil)
	blueWs.Write("commitMatch", nil)
	time.Sleep(time.Millisecond * 10) // Allow some time for the commands to be processed.
	assert.Equal(t, 1, web.arena.ScoringPanelRegistry.GetNumScoreCommitted("red_near"))
	assert.Equal(t, 1, web.arena.ScoringPanelRegistry.GetNumScoreCommitted("blue_near"))

	// Load another match to reset the results.
	web.arena.ResetMatch()
	web.arena.LoadTestMatch()
	readWebsocketType(t, redWs, "matchLoad")
	readWebsocketType(t, redWs, "realtimeScore")
	readWebsocketType(t, blueWs, "matchLoad")
	readWebsocketType(t, blueWs, "realtimeScore")
	assert.Equal(t, field.NewRealtimeScore(), web.arena.RedRealtimeScore)
	assert.Equal(t, field.NewRealtimeScore(), web.arena.BlueRealtimeScore)
	assert.Equal(t, 0, web.arena.ScoringPanelRegistry.GetNumScoreCommitted("red_near"))
	assert.Equal(t, 0, web.arena.ScoringPanelRegistry.GetNumScoreCommitted("blue_near"))
}
