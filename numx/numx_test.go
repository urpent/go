package text

import (
	"testing"

	"github.com/urpent/go/ut"
)

func TestCurrencyString(t *testing.T) {
	t.Run("test currency", func(t *testing.T) {
		value := 1

		result := CurrencyString(value)
		ut.AssertEqual(t, "$1", result)
	})

	t.Run("test currency with float", func(t *testing.T) {
		value := 1000.347

		result := CurrencyString(value)
		ut.AssertEqual(t, "$1,000.34", result)
	})

	t.Run("test currency with string value", func(t *testing.T) {
		value := "1000000.3699999"

		result := CurrencyString(value)
		ut.AssertEqual(t, "$1,000,000.36", result)
	})

	t.Run("test currency with invalid string value", func(t *testing.T) {
		value := "abc"

		result := CurrencyString(value)
		ut.AssertEqual(t, "$", result)
	})
}

func TestDecimalString(t *testing.T) {
	t.Run("test with int", func(t *testing.T) {
		value := 1

		result := DecimalString(value)
		ut.AssertEqual(t, "1", result)
	})

	t.Run("test with int", func(t *testing.T) {
		type CustomInt int

		var value CustomInt = 1

		result := DecimalString(value)
		ut.AssertEqual(t, "1", result)
	})

	t.Run("test with float", func(t *testing.T) {
		value := 1000.347

		result := DecimalString(value)
		ut.AssertEqual(t, "1,000.34", result)
	})

	t.Run("test with custom float", func(t *testing.T) {
		type CustomFloat float64

		var value CustomFloat = 1000.347

		result := DecimalString(value)
		ut.AssertEqual(t, "1,000.34", result)
	})

	t.Run("test with string value", func(t *testing.T) {
		value := "1000000.3699999"

		result := DecimalString(value)
		ut.AssertEqual(t, "1,000,000.36", result)
	})

	t.Run("test with type", func(t *testing.T) {
		type CustomString string

		var value CustomString = "1000000.3699999"

		result := DecimalString(value)
		ut.AssertEqual(t, "1,000,000.36", result)
	})

	t.Run("test with invalid string value", func(t *testing.T) {
		value := "abc"

		result := DecimalString(value)
		ut.AssertEqual(t, "", result)
	})
}

func TestCustomDecimalString(t *testing.T) {
	t.Run("test with float", func(t *testing.T) {
		value := 1000.116999

		result := CustomDecimalString(value, 3, true)
		ut.AssertEqual(t, "1,000.116", result)
	})

	t.Run("test without separator", func(t *testing.T) {
		value := 1000.116999

		result := CustomDecimalString(value, 3, false)
		ut.AssertEqual(t, "1000.116", result)
	})

	t.Run("test with invalid string", func(t *testing.T) {
		value := "abc"

		result := CustomDecimalString(value, 3, true)
		ut.AssertEqual(t, "", result)
	})
}

func TestToDecimal(t *testing.T) {
	t.Run("test int", func(t *testing.T) {
		var int0 int8 = 1

		result, err := ToDecimal(int0)
		ut.AssertEqual(t, "1", result.String())
		ut.AssertEqual(t, nil, err)

		var int1 int32 = 1

		result, err = ToDecimal(int1)
		ut.AssertEqual(t, "1", result.String())
		ut.AssertEqual(t, nil, err)

		var int2 int64 = 1
		result, err = ToDecimal(int2)
		ut.AssertEqual(t, "1", result.String())
		ut.AssertEqual(t, nil, err)
	})

	t.Run("test float", func(t *testing.T) {
		var float1 float32 = 1.1111

		result, err := ToDecimal(float1)
		ut.AssertEqual(t, "1.1111", result.String())
		ut.AssertEqual(t, nil, err)

		var float2 float64 = 1.111111
		result, err = ToDecimal(float2)
		ut.AssertEqual(t, "1.111111", result.String())
		ut.AssertEqual(t, nil, err)
	})

	t.Run("test string", func(t *testing.T) {
		var string1 string = "1.1111"

		result, err := ToDecimal(string1)
		ut.AssertEqual(t, "1.1111", result.String())
		ut.AssertEqual(t, nil, err)
	})
}
