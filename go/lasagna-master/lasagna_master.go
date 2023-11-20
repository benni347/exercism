package lasagna

// PreparationTime calculates the time needed to prepare the lasagna
func PreparationTime(layers []string, layerTime int) int {
	if layerTime == 0 {
		layerTime = 2
	}
	return len(layers) * layerTime
}

// Quantities calculates the quantities of ingredients
func Quantities(layers []string) (int, float64) {
	wordCount := make(map[string]int)
	for _, word := range layers {
		wordCount[word]++
	}
	var sauceAmount float64
	sauceAmount = float64(wordCount["sauce"]) * 0.2
	noodleAmount := wordCount["noodles"] * 50
	return noodleAmount, sauceAmount
}

// AddSecretIngredient adds a secret ingredient
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales the recipe
// One recipe is two portions
func ScaleRecipe(amounts []float64, numPortions int) []float64 {
	scaledAmounts := make([]float64, len(amounts))

	scalingFactor := float64(numPortions) / 2.0

	for i, amount := range amounts {
		scaledAmounts[i] = amount * scalingFactor
	}

	return scaledAmounts
}
