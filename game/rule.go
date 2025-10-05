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
	{1, "G210", true, false, "A strategy aimed at forcing an opponent to violate a rule is not allowed."},
	{2, "G401", false, false, "In AUTO, each DRIVE TEAM member must remain in their staged areas. A DRIVE TEAM member staged behind a HUMAN STARTING LINE may not contact anything in front of that HUMAN STARTING LINE, unless for personal or equipment safety, to press the E-Stop or A-Stop, or granted permission by a Head REFEREE or FTA."},
	{3, "G402", false, false, "In AUTO, a DRIVE TEAM member may not directly or indirectly interact with a ROBOT or an OPERATOR CONSOLE unless for personal safety, OPERATOR CONSOLE safety, or pressing an E-Stop or A-Stop."},
	{4, "G408", true, false, "Neither a ROBOT nor a HUMAN PLAYER may damage a CANNONBALL."},
	{5, "G422", false, false, "A ROBOT may not use a COMPONENT outside its ROBOT PERIMETER (except its BUMPERS) to initiate contact with an opponent ROBOT inside the vertical projection of the opponent's ROBOT PERIMETER."},
	{6, "G423", true, false, "A ROBOT may not damage or functionally impair an opponent ROBOT in either of the following ways: A. deliberately. B. regardless of intent, by initiating contact, either directly or transitively via a SCORING ELEMENT CONTROLLED by the ROBOT, inside the vertical projection of an opponent's ROBOT PERIMETER."},
	{7, "G425", false, false, "A ROBOT may not PIN an opponent's ROBOT for more than 3 seconds."},
	{8, "G429", false, false, "A DRIVE TEAM member must remain in their designated area as follows: A. DRIVERS and COACHES may not contact anything outside their ALLIANCE AREA, B. a DRIVER must use the OPERATOR CONSOLE in the DRIVER STATION to which they are assigned, as indicated on the team sign, C. a HUMAN PLAYER may not contact anything outside their ALLIANCE AREA, and D. a TECHNICIAN may not contact anything outside their designated area."},
	{9, "G430", true, false, "A ROBOT shall be operated only by the DRIVERS and/or HUMAN PLAYERS of that team. A COACH activating their E-Stop or A-Stop is the exception to this rule."},
	{10, "G434", false, false, "COACHES may not touch CANNONBALLS, unless for safety purposes."},
	{11, "G435", true, false, "A ROBOT or HUMAN PLAYER may not intentionally permanently damage a gamepiece."},
	{12, "M101", true, false, "A ROBOT may not intentionally eject a CANNONBALL from the FIELD (either directly or by bouncing off a FIELD element or other ROBOT)."},
	{13, "M102", true, true, "A ROBOT may not go into other teams' endgame areas within the endgame for over 3 seconds."},
	{14, "M103", true, true, "A ROBOT may contact an opposing robot while they are attempting to score while in their own KRAKEN LAIR."},
	{15, "M104", true, true, "A ROBOT may not enter the opposing alliance's SAFE HARBOR perimter during auton."},
	{16, "M105", false, false, "A ROBOT may not shoot a CANNONBALL from inside the HUMAN PLAYER ZONE."},
	{17, "M106", false, false, "A ROBOT may not cross the midfield line during auton."},
	{18, "M107", false, false, "A ROBOT may not enter the opposing team's HUMAN PLAYER ZONE."},
	{19, "M108", false, false, "A HUMAN PLAYER may not place or throw a CANNONBALL onto the field except through the holes in the HUMAN PLAYER STATION holes."},
	{20, "M109", false, false, "A ROBOT may not repeatedly contact scoring structures in an unsafe manner."},
	{21, "M110", false, false, "A ROBOT may not shoot a CANNONBALL into the opposing team's HUMAN PLAYER ZONE."},
	{22, "M111", false, false, "A ROBOT may not shoot a CANNONBALL into the KRAKEN LAIR outside of ENDGAME."},
	{23, "M112", false, false, "A ROBOT may not have more than momentary control of more than one CANNONBALL at any time."},
	{24, "M113", false, false, "A ROBOT may not start with more than one pre-loaded CANNONBALLs."},
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
