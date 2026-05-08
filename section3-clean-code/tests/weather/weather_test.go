package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

type WeatherServiceStub struct {
	deg int
}

func (s *WeatherServiceStub) Forecast() int {
	return s.deg
}

func TestForecast(t *testing.T) {
	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			service := &WeatherServiceStub{
				deg: test.deg,
			}
			weather := Weather{service}
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
