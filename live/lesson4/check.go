package lesson4

import "math"

func CheckBudget(oldBudget, newBudget float64, countryCode string) bool {
	var (
		limit     float64 = 20 //分界线
		spanLimit float64 = 1  //跨度值
		diff      float64
		rateLimit = 0.05
	)

	diff = math.Abs(newBudget - oldBudget)
	if diff == 0 {
		return false
	}

	if countryCode == "JP" {
		limit *= 100
		spanLimit *= 100
	}

	if oldBudget > limit {
		if diff >= spanLimit {
			return true
		} else {
			return false
		}
	} else {
		rate := diff / oldBudget
		if rate >= rateLimit {
			return true
		} else {
			return false
		}
	}
}
