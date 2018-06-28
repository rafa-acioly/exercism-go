package leap

const testVersion = 3

// Verify if it is a leap year
func IsLeapYear(year int) bool {
	// Is divisible by 400? if so, it its a leap year
	if year%400 == 0 {
		return true
	}

	// the year must be divisible by 4 and not by 100
	return (year%4 == 0) && !(year%100 == 0)
}
