package cars

import "math"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(math.Round(float64(productionRate)*successRate/100) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	return uint(carsCount%10*10000 + (carsCount / 10 * 95000))
}
