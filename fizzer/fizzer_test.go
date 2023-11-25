package fizzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fizzbuzz(t *testing.T) {
	resultData := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"}

	t.Run("When variabe are correctly set Should return the correct response", func(t *testing.T) {
		fizzer, err := NewClient(20, 3, 5, "fizz", "buzz")

		result := fizzer.Fizzbuzz()

		assert.Equal(t, resultData, result)
		assert.NoError(t, err)
	})

	t.Run("When the variable are incorrect Should return an error", func(t *testing.T) {
		fizzer, err := NewClient(20, 3, 3, "fizz", "buzz")

		result := fizzer.Fizzbuzz()

		assert.Equal(t, []string(nil), result)
		assert.Error(t, err)
	})
}

func Test_validator(t *testing.T) {
	t.Run("When the modulo are equals Should return an error", func(t *testing.T) {
		err := validator(10, 3, 3, "fizz", "buzz")

		assert.Error(t, err)
	})

	t.Run("When the words are equals Should return an error", func(t *testing.T) {
		err := validator(10, 3, 5, "fizz", "fizz")

		assert.Error(t, err)
	})

	t.Run("When the limit number is inferior to 1 Should return an error", func(t *testing.T) {
		err := validator(0, 3, 5, "fizz", "buzz")

		assert.Error(t, err)
	})

	t.Run("When variabe are correctly set Should return NoError", func(t *testing.T) {
		err := validator(10, 3, 5, "fizz", "buzz")

		assert.NoError(t, err)
	})
}

func Test_Transformer(t *testing.T) {
	t.Run("When limitNumber variabe are not a string integer Should return an error", func(t *testing.T) {
		limitNumber, fizzModulo, buzzModulo, err := Transformer("testValue", "1", "1")

		assert.Equal(t, 0, limitNumber)
		assert.Equal(t, 0, fizzModulo)
		assert.Equal(t, 0, buzzModulo)
		assert.Error(t, err)
	})

	t.Run("When fizzModulo variabe are not a string integer Should return an error", func(t *testing.T) {
		limitNumber, fizzModulo, buzzModulo, err := Transformer("1", "testValue", "1")

		assert.Equal(t, 0, limitNumber)
		assert.Equal(t, 0, fizzModulo)
		assert.Equal(t, 0, buzzModulo)
		assert.Error(t, err)
	})

	t.Run("When buzzModulo variabe are not a string integer Should return an error", func(t *testing.T) {
		limitNumber, fizzModulo, buzzModulo, err := Transformer("1", "1", "testValue")

		assert.Equal(t, 0, limitNumber)
		assert.Equal(t, 0, fizzModulo)
		assert.Equal(t, 0, buzzModulo)
		assert.Error(t, err)
	})

	t.Run("When variabe are empty Should return an error", func(t *testing.T) {
		limitNumber, fizzModulo, buzzModulo, err := Transformer("", "", "")

		assert.Equal(t, 0, limitNumber)
		assert.Equal(t, 0, fizzModulo)
		assert.Equal(t, 0, buzzModulo)
		assert.Error(t, err)
	})

	t.Run("When variabe are string integer Should return the correct response", func(t *testing.T) {
		limitNumber, fizzModulo, buzzModulo, err := Transformer("1", "1", "1")

		assert.Equal(t, 1, limitNumber)
		assert.Equal(t, 1, fizzModulo)
		assert.Equal(t, 1, buzzModulo)
		assert.NoError(t, err)
	})
}
