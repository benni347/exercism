// Package weather is a package which tells you the current weather based on your position and the current weather condition.
package weather

// CurrentCondition is a string what the current weather scenario is.
var CurrentCondition string

// CurrentLocation is a string based on your current location.
var CurrentLocation string

// Forecast will return the Current location and the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
