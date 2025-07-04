// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package web

import (
	"github.com/Team254/cheesy-arena/game"
	"github.com/Team254/cheesy-arena/model"
	"github.com/Team254/cheesy-arena/tournament"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRankingsCsvReport(t *testing.T) {
	web := setupTestWeb(t)

	ranking1 := game.TestRanking2()
	ranking2 := game.TestRanking1()
	web.arena.Database.CreateRanking(ranking1)
	web.arena.Database.CreateRanking(ranking2)

	recorder := web.getHttpResponse("/reports/csv/rankings")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "text/plain", recorder.Header()["Content-Type"][0])
	expectedBody1 := "Rank,TeamId,RankingPoints,CoopertitionPoints,MatchPoints,AutoPoints,BargePoints,Wins,Losses," +
		"Ties,Disqualifications,Played\r\n1,254,20,625,90,554,12,3,2,1,0,10\r\n2,1114,18,700,625,90,23,1,3,2,0,10\r\n\r\n"
	expectedBody2 := "Rank,TeamId,RankingPoints,CoopertitionPoints,MatchPoints,AutoPoints,BargePoints,Wins,Losses," +
		"Ties,Disqualifications,Played\n1,254,20,625,90,554,12,3,2,1,0,10\n2,1114,18,700,625,90,23,1,3,2,0,10\n\n"
	//assert.Equal(t, expectedBody, recorder.Body.String())
	if recorder.Body.String() != expectedBody1 && recorder.Body.String() != expectedBody2 {
        t.Errorf("Unexpected body: %s", recorder.Body.String())
    }
}

func TestRankingsPdfReport(t *testing.T) {
	web := setupTestWeb(t)

	ranking1 := game.TestRanking2()
	ranking2 := game.TestRanking1()
	web.arena.Database.CreateRanking(ranking1)
	web.arena.Database.CreateRanking(ranking2)

	// Can't really parse the PDF content and check it, so just check that what's sent back is a PDF.
	recorder := web.getHttpResponse("/reports/pdf/rankings")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "application/pdf", recorder.Header()["Content-Type"][0])
}

func TestScheduleCsvReport(t *testing.T) {
	web := setupTestWeb(t)

	match1Time := time.Unix(0, 0)
	match1 := model.Match{Type: model.Qualification, ShortName: "Q1", Time: match1Time, Red1: 1, Red2: 2, Red3: 3,
		Blue1: 4, Blue2: 5, Blue3: 6, Blue1IsSurrogate: true, Blue2IsSurrogate: true, Blue3IsSurrogate: true}
	match2Time := time.Unix(600, 0)
	match2 := model.Match{Type: model.Qualification, ShortName: "Q2", Time: match2Time, Red1: 7, Red2: 8, Red3: 9,
		Blue1: 10, Blue2: 11, Blue3: 12, Red1IsSurrogate: true, Red2IsSurrogate: true, Red3IsSurrogate: true}
	match3 := model.Match{Type: model.Practice, ShortName: "P1", Time: time.Now(), Red1: 6, Red2: 5, Red3: 4,
		Blue1: 3, Blue2: 2, Blue3: 1}
	web.arena.Database.CreateMatch(&match1)
	web.arena.Database.CreateMatch(&match2)
	web.arena.Database.CreateMatch(&match3)

	recorder := web.getHttpResponse("/reports/csv/schedule/qualification")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "text/plain", recorder.Header()["Content-Type"][0])
	expectedBody1 := "Match,Type,Time,Red1,Red1IsSurrogate,Red2,Red2IsSurrogate,Red3,Red3IsSurrogate,Blue1," +
		"Blue1IsSurrogate,Blue2,Blue2IsSurrogate,Blue3,Blue3IsSurrogate\r\nQ1,Qualification," + match1Time.String() +
		",1,false,2,false,3,false,4,true,5,true,6,true\r\nQ2,Qualification," + match2Time.String() +
		",7,true,8,true,9,true,10,false,11,false,12,false\r\n\r\n"
	expectedBody2 := "Match,Type,Time,Red1,Red1IsSurrogate,Red2,Red2IsSurrogate,Red3,Red3IsSurrogate,Blue1," +
		"Blue1IsSurrogate,Blue2,Blue2IsSurrogate,Blue3,Blue3IsSurrogate\nQ1,Qualification," + match1Time.String() +
		",1,false,2,false,3,false,4,true,5,true,6,true\nQ2,Qualification," + match2Time.String() +
		",7,true,8,true,9,true,10,false,11,false,12,false\n\n"
	//assert.Equal(t, expectedBody, recorder.Body.String())
	if recorder.Body.String() != expectedBody1 && recorder.Body.String() != expectedBody2 {
        t.Errorf("Unexpected body: %s", recorder.Body.String())
    }
}

func TestSchedulePdfReport(t *testing.T) {
	web := setupTestWeb(t)

	match := model.Match{Type: model.Practice, ShortName: "P1", Time: time.Unix(0, 0), Red1: 1, Red2: 2, Red3: 3,
		Blue1: 4, Blue2: 5, Blue3: 6, Blue1IsSurrogate: true, Blue2IsSurrogate: true, Blue3IsSurrogate: true}
	web.arena.Database.CreateMatch(&match)
	team := model.Team{Id: 254, Name: "NASA", Nickname: "The Cheesy Poofs", City: "San Jose", StateProv: "CA",
		Country: "USA", RookieYear: 1999, RobotName: "Barrage"}
	web.arena.Database.CreateTeam(&team)

	// Can't really parse the PDF content and check it, so just check that what's sent back is a PDF.
	recorder := web.getHttpResponse("/reports/pdf/schedule/practice")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "application/pdf", recorder.Header()["Content-Type"][0])
}

func TestTeamsCsvReport(t *testing.T) {
	web := setupTestWeb(t)

	team1 := model.Team{Id: 254, Name: "NASA", Nickname: "The Cheesy Poofs", SchoolName: "Bellarmine College Preparatory", City: "San Jose", StateProv: "CA",
		Country: "USA", RookieYear: 1999, RobotName: "Barrage"}
	team2 := model.Team{Id: 1114, Name: "GM", Nickname: "Simbotics", SchoolName: "Governor Simcoe Secondary School", City: "St. Catharines", StateProv: "ON",
		Country: "Canada", RookieYear: 2003, RobotName: "Simbot Evolution"}
	web.arena.Database.CreateTeam(&team1)
	web.arena.Database.CreateTeam(&team2)

	recorder := web.getHttpResponse("/reports/csv/teams")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "text/plain", recorder.Header()["Content-Type"][0])
	expectedBody1 := "Number,Name,Nickname,SchoolName,City,StateProv,Country,RookieYear,RobotName,HasConnected\r\n254,\"NASA\"," +
		"\"The Cheesy Poofs\",\"Bellarmine College Preparatory\",\"San Jose\",\"CA\",\"USA\",1999,\"Barrage\",false\r\n1114,\"GM\",\"Simbotics\"," +
		"\"Governor Simcoe Secondary School\",\"St. Catharines\",\"ON\",\"Canada\",2003,\"Simbot Evolution\",false\r\n\r\n"
	expectedBody2 := "Number,Name,Nickname,SchoolName,City,StateProv,Country,RookieYear,RobotName,HasConnected\n254,\"NASA\"," +
		"\"The Cheesy Poofs\",\"Bellarmine College Preparatory\",\"San Jose\",\"CA\",\"USA\",1999,\"Barrage\",false\n1114,\"GM\",\"Simbotics\"," +
		"\"Governor Simcoe Secondary School\",\"St. Catharines\",\"ON\",\"Canada\",2003,\"Simbot Evolution\",false\n\n"
	//assert.Equal(t, expectedBody, recorder.Body.String())
	if recorder.Body.String() != expectedBody1 && recorder.Body.String() != expectedBody2 {
        t.Errorf("Unexpected body: %s", recorder.Body.String())
    }
}

func TestTeamsPdfReport(t *testing.T) {
	web := setupTestWeb(t)

	team := model.Team{Id: 254, Name: "NASA", Nickname: "The Cheesy Poofs", City: "San Jose", StateProv: "CA",
		Country: "USA", RookieYear: 1999, RobotName: "Barrage"}
	web.arena.Database.CreateTeam(&team)

	// Can't really parse the PDF content and check it, so just check that what's sent back is a PDF.
	recorder := web.getHttpResponse("/reports/pdf/teams")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "application/pdf", recorder.Header()["Content-Type"][0])
}

func TestWpaKeysCsvReport(t *testing.T) {
	web := setupTestWeb(t)

	team1 := model.Team{Id: 254, WpaKey: "12345678"}
	team2 := model.Team{Id: 1114, WpaKey: "9876543210"}
	web.arena.Database.CreateTeam(&team1)
	web.arena.Database.CreateTeam(&team2)

	recorder := web.getHttpResponse("/reports/csv/wpa_keys")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "text/csv", recorder.Header()["Content-Type"][0])
	assert.Equal(t, "attachment; filename=keys.csv", recorder.Header()["Content-Disposition"][0])
	assert.Equal(t, "254,12345678\r\n1114,9876543210\r\n", recorder.Body.String())
}

func TestAlliancesPdfReport(t *testing.T) {
	web := setupTestWeb(t)
	tournament.CreateTestAlliances(web.arena.Database, 8)
	web.arena.CreatePlayoffTournament()

	// Can't really parse the PDF content and check it, so just check that what's sent back is a PDF.
	recorder := web.getHttpResponse("/reports/pdf/alliances")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "application/pdf", recorder.Header()["Content-Type"][0])
}

func TestBracketPdfReport(t *testing.T) {
	web := setupTestWeb(t)

	recorder := web.getHttpResponse("/reports/pdf/bracket")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "text/html; charset=utf-8", recorder.Header()["Content-Type"][0])
	assert.Contains(t, recorder.Body.String(), "Finals")
}

func TestJudgingSchedulePdfReport(t *testing.T) {
	web := setupTestWeb(t)

	// Can't really parse the PDF content and check it, so just check that what's sent back is a PDF.
	recorder := web.getHttpResponse("/reports/pdf/judging_schedule")
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "application/pdf", recorder.Header()["Content-Type"][0])
}
