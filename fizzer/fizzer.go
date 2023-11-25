// Package fizzer allow you to get the list of number
package fizzer

import (
	"fmt"
	"strconv"
)

type fizzer struct {
	limitNumber int
	fizzWord    string
	buzzWord    string
	fizzModulo  int
	buzzModulo  int
}

func NewClient(limitNumber, fizzModulo, buzzModulo int, fizzWord, buzzWord string) (fizzer, error) {
	// Checking we are using the correct data
	err := validator(limitNumber, fizzModulo, buzzModulo, fizzWord, buzzWord)
	if err != nil {
		return fizzer{}, err
	}

	return fizzer{
		limitNumber: limitNumber,
		fizzWord:    fizzWord,
		buzzWord:    buzzWord,
		fizzModulo:  fizzModulo,
		buzzModulo:  buzzModulo,
	}, nil
}

func (f *fizzer) Fizzbuzz() (result []string) {
	for i := 1; i <= f.limitNumber; i++ {
		switch {
		case i%f.fizzModulo == 0 && i%f.buzzModulo == 0:
			result = append(result, f.fizzWord+f.buzzWord)
		case i%f.fizzModulo == 0:
			result = append(result, f.fizzWord)
		case i%f.buzzModulo == 0:
			result = append(result, f.buzzWord)
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return
}

func validator(limitNumber, fizzModulo, buzzModulo int, fizzWord, buzzWord string) error {
	if fizzModulo == buzzModulo {
		return fmt.Errorf("invalid: the modulo are equals, the number provided are %v and %v ", fizzModulo, buzzModulo)
	}

	if limitNumber < 1 {
		return fmt.Errorf("invalid: the limit number is inferior to 1, the number provided is %v", limitNumber)
	}

	if fizzWord == buzzWord {
		return fmt.Errorf("invalid: the words are equals, the words provided are %v and %v ", fizzWord, buzzWord)
	}

	return nil
}

// Transformer transform fizzer params string data into int
func Transformer(limitNumber, fizzModulo, buzzModulo string) (limitNumberInt, fizzModuloInt, buzzModuloInt int, err error) {
	limitNumberInt, err = strconv.Atoi(limitNumber)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid: the limitNumber provided is incorrect : %s", limitNumber)
	}

	fizzModuloInt, err = strconv.Atoi(fizzModulo)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid: the modulo provided is incorrect : %s", fizzModulo)
	}

	buzzModuloInt, err = strconv.Atoi(buzzModulo)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid: the Modulo provided is incorrect : %s", buzzModulo)
	}

	return
}
