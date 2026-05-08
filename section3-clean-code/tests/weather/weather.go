// 🤔 Testing the weather
//
// Suppose we have a WeatherService that predicts the weather. It works by a
// mysterious and unpredictable algorithm full of artificial intelligence,
// returning the expected daytime temperature for tomorrow in Celsius:
//
//   func (ws *WeatherService) Forecast() int {
//       // magic
//       return value
//   }
//
// We made a Weather struct that queries WeatherService and returns the forecast,
// but as text, not as a number.
//
// Now we'd like to test Weather. The problem is we have no control over the
// behavior of WeatherService, so the tests fail.
//
// Fix Weather so it can be tested properly. Make a stub for WeatherService and
// use it in TestForecast.
//
// Do NOT change: WeatherService, Weather.Forecast() method, the tests slice.
//
// Hint: as a result, Weather should work both for normal use (with WeatherService)
// and for tests (with a stub). If you see "cannot use stub (variable of type
// *WeatherServiceStub_) as type..." then you haven't achieved that yet.
//
// Test data:
//
//   # Test     Input data    Output data
//     1        (none)        PASS

package main

type Forecaster interface {
	Forecast() int
}

// WeatherService predicts the weather.
type WeatherService struct{}

// Forecast reports the expected daytime temperature for tomorrow.
func (ws *WeatherService) Forecast() int {
	// magic
	return 0
}

// Weather provides a textual weather forecast.
type Weather struct {
	service Forecaster
}

// Forecast reports the textual weather forecast for tomorrow.
func (w Weather) Forecast() string {
	deg := w.service.Forecast()
	switch {
	case deg < 10:
		return "холодно"
	case deg >= 10 && deg < 15:
		return "прохладно"
	case deg >= 15 && deg < 20:
		return "идеально"
	case deg >= 20:
		return "жарко"
	}
	return "инопланетно"
}
