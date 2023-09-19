package util

const (
	USD = "USD"
	EUR = "EUR"
	TL  = "TL"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TL:
		return true
	}
	return false
}
