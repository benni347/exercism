package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	whichOption := strings.Compare(option1, option2)
	if whichOption == -1 {
		return option1 + " is clearly the better choice."
	} else if whichOption == 1 {
		return option2 + " is clearly the better choice."
	} else {
		return "Is the same car"
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// if age is < 3 years = 80% of original price
	// if age is >= 10 years = 50% of original price
	// else 70%
	if age < 3.0 {
		return originalPrice * 0.8
	} else if age >= 10.0 {
		return originalPrice * 0.5
	} else {
		return originalPrice * 0.7
	}
}
