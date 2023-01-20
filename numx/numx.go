package text

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Number interface {
	constraints.Integer | constraints.Float | ~string
}

// CurrencyString formats as following: 10.00 -> $10, 10.50 -> $10.50, 10.55 -> $10.55, 10.556 -> $10.55
func CurrencyString[T Number](value T) string {
	result := DecimalString(value)
	return fmt.Sprintf("$%s", result)
}

// DecimalString formats as following: 10.00 -> 10, 10.50 -> 10.50, 10.55 -> 10.55, 10.556 -> 10.55
func DecimalString[T Number](value T) string {
	num, err := ToDecimal(value)
	if err != nil {
		return ""
	}

	p := message.NewPrinter(language.English)
	result := p.Sprintf("%.*f", 2, num.Truncate(2).InexactFloat64())

	return strings.TrimSuffix(result, ".00") // Remove trailing zero
}

// CustomDecimalString formats value given 2 precision as following: 10.00 -> 10, 10.50 -> 10.50, 10.55 -> 10.55, 10.556 -> 10.55
// If showThousandsSeparator is true, it will add thousands separator as following: 1000000.00 -> 1,000,000
func CustomDecimalString[T Number](value T, precision int32, showThousandsSeparator bool) string {
	num, err := ToDecimal(value)
	if err != nil {
		return ""
	}

	var result string

	if showThousandsSeparator {
		p := message.NewPrinter(language.English)
		result = p.Sprintf("%.*f", precision, num.Truncate(precision).InexactFloat64())
	} else {
		result = num.Truncate(precision).StringFixed(precision)
	}

	return strings.TrimSuffix(result, ".00") // Remove trailing zero
}

func ToDecimal[T Number](value T) (decimal.Decimal, error) {
	s := fmt.Sprintf("%v", value)
	return decimal.NewFromString(s)
}
