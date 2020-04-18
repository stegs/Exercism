// Package leap contains methods to determine if the year is a leap year
package leap

// IsLeapYear returns true or false depending on if the year provided is a leap year
func IsLeapYear(year int) bool {

	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
