package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * ((successRate) / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var CarsPerHour float64
	CarsPerHour = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(CarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var Single int64
	Single = int64(carsCount) % 10
	var Multiple int64
	Multiple = (int64(carsCount) - Single) / 10
	const tenPack int64 = 95000
	const singlePack int64 = 100 * 100
	costSingle := Single * singlePack
	costMultiple := tenPack * Multiple
	return uint(costMultiple + costSingle)
}
