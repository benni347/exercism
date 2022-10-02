package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	formattedString := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return formattedString
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcomeMsg := Welcome(name)
	assigned := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	neighborStr := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return welcomeMsg + "\n" + assigned + "\n" + neighborStr
}
