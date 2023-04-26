package utils

// List of Constant Currency
const (
	RUB = "RUB"
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

// Registry For Allowed Currency
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case RUB, USD, CAD, EUR:
		return true
	}
	return false
}
