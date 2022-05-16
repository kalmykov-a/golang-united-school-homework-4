package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	//Проверка, что строка не пустая
	if input == "" {
		return "", fmt.Errorf("Empty string: %w", errorEmptyInput)
	}
	sum := 0
	// удаляем лишние пробелы
	str := strings.ReplaceAll(input, " ", "")
	// Добавляем оператор + перед каждым минусом не в начале строки
	str = str[:1] + strings.ReplaceAll(input, "-", "+-")
	if strings.Contains(str[1:], "-") {
		str = str[:strings.Index(str[1:], "-")+1] + "+" + str[strings.Index(str[1:], "-")+1:]
	}
	strArr := strings.Split(str, "+")
	if len(strArr) == 2 {
		for _, v := range strArr {
			num, err := strconv.Atoi(string(v))
			sum += num
			if err != nil {
				return "", fmt.Errorf("Not a number: %w", err)
			}
		}
	} else {
		return "", fmt.Errorf("Less or more than two operands: %w", errorNotTwoOperands)
	}
	return strconv.Itoa(sum), nil
}
