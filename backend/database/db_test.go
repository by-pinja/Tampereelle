package database

import (
	"testing"
	"math"
)

func question(latitude float64, longitude float64) Question {
	place := Place{Latitude: latitude, Longitude: longitude}
	return Question{Place: place}
}


func TestScoreCalculation1(t *testing.T) {
	q := question(90.0, 90.0)
	a := Answer{PlayerLatitude: 0.0, PlayerLongitude: 0.0, Angle: 0}
	score := getPlayerScore(q, a)
	if score != math.Pi / 4 {
		t.Error(score)
	}
}


func TestScoreCalculation2(t *testing.T) {
	q := question(0, -40.0)
	a := Answer{PlayerLatitude: 0.0, PlayerLongitude: 0.0, Angle: 90}
	score := getPlayerScore(q, a)
	if score != math.Pi / 2 {
		t.Error(score)
	}
}

