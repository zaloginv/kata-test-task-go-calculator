package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	operands    = "+-*/"
	leftBorder  = 1
	rightBorder = 10
	digits      = "1234567890"
)

var (
	romanMap    = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	romansSlice = []string{"I", "V", "X", "L", "C", "D", "M"}
	digitSlice  = strings.Split(digits, "")
)

//func mapToSlice(mapObject map[string]int) []string {
//	sliceObj := make([]string, 0, len(mapObject))
//	for obj, _ := range mapObject {
//		sliceObj = append(sliceObj, obj)
//	}
//
//	return sliceObj
//}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func arabicCalc(num1 int, operand string, num2 int) (result int) {

	switch operand {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	return result
}

func romanToArabic(roman string) (result int) {
	romanSlice := strings.Split(roman, "")
	for index, _ := range romanSlice {
		if index+1 < len(roman) &&
			romanMap[romanSlice[index]] < romanMap[romanSlice[index+1]] {
			result -= romanMap[romanSlice[index]]
		} else {
			result += romanMap[romanSlice[index]]
		}
	}
	return result
}

func arabicToRoman(arabic int) (result string) {
	arabicStr := strconv.Itoa(arabic)
	arabicStrReverse := Reverse(arabicStr)
	result = ""
	romansPointer := 0

	for index, symbol := range arabicStrReverse {
		strArabic := string(arabicStrReverse[index])
		intArabic, _ := strconv.Atoi(strArabic)
		if stringInSlice(string(symbol), []string{"0", "1", "2", "3"}) {
			result = strings.Repeat(romansSlice[romansPointer], intArabic) + result
		} else if stringInSlice(string(symbol), []string{"4"}) {
			result = romansSlice[romansPointer] + romansSlice[romansPointer+1] + result
		} else if stringInSlice(string(symbol), []string{"5", "6", "7", "8"}) {
			result = romansSlice[romansPointer+1] + strings.Repeat(romansSlice[romansPointer], intArabic-5) + result
		} else if stringInSlice(string(symbol), []string{"9"}) {
			result = romansSlice[romansPointer] + romansSlice[romansPointer+2] + result
		}
		romansPointer += 2
	}

	return result
}

func numbersCheck(num1 int, num2 int) (result bool) {
	result = (leftBorder <= num1 && num1 <= rightBorder) && (leftBorder <= num2 && num2 <= rightBorder)
	return result
}

func anyNumberCheck(num string, anyNumbers []string) (result bool) {
	var checks []bool

	for _, element := range num {
		elementString := string(element)
		if stringInSlice(elementString, anyNumbers) {
			checks = append(checks, true)
		} else {
			checks = append(checks, false)
		}
	}

	for _, element := range checks {
		if element == false {
			return false
		}
	}

	return true
}

func calculate(userString string) string {
	symbolsSlice := strings.Split(userString, " ")

	if len(symbolsSlice) < 3 {
		return "The string you entered contains not enough values."
	} else if len(symbolsSlice) > 3 {
		return "The string you entered contains too many values."
	}

	num1 := symbolsSlice[0]
	operand := symbolsSlice[1]
	num2 := symbolsSlice[2]

	if !(stringInSlice(operand, strings.Split(operands, ""))) {
		return "There is no operand like this"

	} else if anyNumberCheck(num1, digitSlice) && anyNumberCheck(num2, digitSlice) {
		num1Int, _ := strconv.Atoi(num1)
		num2Int, _ := strconv.Atoi(num2)

		if numbersCheck(num1Int, num2Int) {
			return strconv.Itoa(arabicCalc(num1Int, operand, num2Int))
		} else {
			return "One or both numbers are not in the specified numeric range."
		}

	} else if anyNumberCheck(num1, romansSlice) && anyNumberCheck(num2, romansSlice) {
		num1Int := romanToArabic(num1)
		num2Int := romanToArabic(num2)

		if numbersCheck(num1Int, num2Int) {
			if num1Int <= num2Int && operand == "-" {
				return "There are no negative numbers (and zero) in the Roman numeral system."
			}
			return arabicToRoman(arabicCalc(num1Int, operand, num2Int))
		} else {
			return "One or both numbers are not in the specified numeric range."
		}

	} else if anyNumberCheck(num1, romansSlice) && anyNumberCheck(num2, digitSlice) ||
		anyNumberCheck(num2, romansSlice) && anyNumberCheck(num1, digitSlice) {
		return "Numbers in different number systems."
	}

	return "Something wrong with entered string."
}

func main() {
	for {
		var userString string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userString = scanner.Text()
		result := calculate(userString)
		fmt.Println(result)
	}
}
