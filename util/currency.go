package util

// Constants for all supported currencies
const (
	ZAR = "ZAR"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case ZAR:
		return true
	}

	return false
}