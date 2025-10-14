// Copyright 2020 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package game

type Rule struct {
	Id             int
	RuleNumber     string
	IsMajor        bool
	IsRankingPoint bool
	Description    string
}

// A curated list of generic rules that carry point penalties.
// @formatter:off
var rules = []*Rule{
	// General Conduct
	{1, "G206", false, false, "A team or ALLIANCE may not collude with another team to each purposefully violate a rule."},
	{2, "G210", true, false, "A strategy aimed at forcing an opponent to violate a rule is not allowed."},
	{3, "G401", false, false, "In AUTO, each DRIVE TEAM member must remain in their staged areas. A DRIVE TEAM member staged behind a HUMAN STARTING LINE may not contact anything in front of that HUMAN STARTING LINE, unless for personal or equipment safety, to press the E-Stop or A-Stop, or granted permission by a Head REFEREE or FTA."},
	{4, "G402", false, false, "In AUTO, a DRIVE TEAM member may not directly or indirectly interact with a ROBOT or an OPERATOR CONSOLE unless for personal safety, OPERATOR CONSOLE safety, or pressing an E-Stop or A-Stop."},
	{5, "G408", true, false, "Neither a ROBOT nor a HUMAN PLAYER may damage a CANNONBALL."},
	{6, "G422", false, false, "A ROBOT may not contact another robot inside of its perimeter."},
	{7, "G423", true, false, "A ROBOT may not damage or functionally impair an opponent ROBOT in either of the following ways: A. deliberately. B. regardless of intent, by initiating contact, either directly or transitively via a SCORING ELEMENT CONTROLLED by the ROBOT, inside the vertical projection of an opponent's ROBOT PERIMETER."},
	{8, "G424", true, false, "A ROBOT may not deliberately attach to, tip, or entangle with an opponent ROBOT."},
	{9, "G425", false, false, "A ROBOT may not PIN an opponent's ROBOT for more than 3 seconds."},
	{10, "G429", false, false, "A DRIVE TEAM member must remain in their designated area as follows: A. DRIVERS and COACHES may not contact anything outside their ALLIANCE AREA, B. a DRIVER must use the OPERATOR CONSOLE in the DRIVER STATION to which they are assigned, as indicated on the team sign, C. a HUMAN PLAYER may not contact anything outside their ALLIANCE AREA, and D. a TECHNICIAN may not contact anything outside their designated area."},
	{11, "G430", true, false, "A ROBOT shall be operated only by the DRIVERS."},
	{12, "G434", false, false, "COACHES may not touch CANNONBALLS, unless for safety purposes."},
	{13, "G435", true, false, "A ROBOT or HUMAN PLAYER may not intentionally permanently damage a gamepiece."},
	{14, "M2500", true, false, "A ROBOT may not intentionally eject a CANNONBALL from the FIELD (either directly or by bouncing off a FIELD element or other ROBOT)."},
	{15, "M2501", true, true, "A ROBOT may not go into other teams' endgame areas within the endgame for over 3 seconds."},
	{16, "M2502", true, true, "A ROBOT may not contact an opposing robot during ENDGAME, while the opposing robot is in their own KRAKEN LAIR."},
	{17, "M2503", true, true, "A ROBOT may not enter the opposing alliance's SAFE HARBOR perimter during auton."},
	{18, "M2504", false, false, "A ROBOT may not shoot a CANNONBALL from inside the HUMAN PLAYER ZONE."},
	{19, "M2505", false, false, "A ROBOT may not cross the midfield line during auton."},
	{20, "M2506", false, false, "A ROBOT may not enter the opposing team's HUMAN PLAYER ZONE."},
	{21, "M2507", false, false, "A HUMAN PLAYER may not place or throw a CANNONBALL onto the field except through the holes in the Human Player Station."},
	{22, "M2508", false, false, "A HUMAN PLAYER must place or throw a CANNONBALL on their own alliance’s robots, or the Human Player Station Floor."},
	{23, "M2510", false, false, "A ROBOT may not repeatedly contact scoring structures in an unsafe manner."},
	{24, "M2511", false, false, "A ROBOT may not shoot a CANNONBALL into the opposing team's HUMAN PLAYER ZONE."},
	{25, "M2512", false, false, "A ROBOT may not shoot a CANNONBALL into the KRAKEN LAIR outside of ENDGAME."},
	{26, "M2513", false, false, "A ROBOT may not have more than momentary control of more than one CANNONBALL at any time."},
	{27, "M2514", false, false, "A ROBOT may not start with more than one pre-loaded CANNONBALLs."},
}

// @formatter:on
var ruleMap map[int]*Rule

// Returns the rule having the given ID, or nil if no such rule exists.
func GetRuleById(id int) *Rule {
	return GetAllRules()[id]
}

// Returns a slice of all defined rules that carry point penalties.
func GetAllRules() map[int]*Rule {
	if ruleMap == nil {
		ruleMap = make(map[int]*Rule, len(rules))
		for _, rule := range rules {
			ruleMap[rule.Id] = rule
		}
	}
	return ruleMap
}
