package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var a string

	fmt.Println(" Выберите калькулятор в какой системе хотите считать ")
	fmt.Println("")
	fmt.Println(" Введите: ")
	fmt.Println("")
	fmt.Println(" |1| если будет использован калькулятор для римских цифр ")
	fmt.Println("")
	fmt.Println(" |2| если будет использован калькулятор для арабских цифр ")
	fmt.Println("")

	fmt.Scanln(&a)

	if a == "1" {
		var operator string
		var rim1, rim2 string
		var arab1, arab2 int
		fmt.Print("Введите первое число: ")
		fmt.Scan(&rim1)
		if rim1 == "I" || rim1 == "II" || rim1 == "III" || rim1 == "IV" || rim1 == "V" || rim1 == "VI" || rim1 == "VII" || rim1 == "VIII" || rim1 == "IX" || rim1 == "X" {
			arab1 = romanToInt(rim1)

			if arab1 > 10 {
				fmt.Print("Превышено число максимума по заданию")
				return
			}

			fmt.Print("Введите второе число: ")
			fmt.Scan(&rim2)
			if rim2 == "I" || rim2 == "II" || rim2 == "III" || rim2 == "IV" || rim2 == "V" || rim2 == "VI" || rim2 == "VII" || rim2 == "VIII" || rim2 == "IX" || rim2 == "X" {
				arab2 = romanToInt(rim2)
				if arab2 > 10 {
					fmt.Print("Превышено число максимума по заданию")
					return
				}
				{

					fmt.Print("\n Введите оператор (+ - * /): ")
					fmt.Scan(&operator)
					switch operator {
					case "+":
						fmt.Printf("%v %s %v = %v\n", intToRoman(arab1), operator, intToRoman(arab2), intToRoman(arab1+arab2))
					case "/":
						if arab1 > arab2 {
							fmt.Printf("%v %s %v = %v\n", intToRoman(arab1), operator, intToRoman(arab2), intToRoman(arab1/arab2))
						} else if arab1 < arab2 {
							fmt.Print("Так как в римской системе нет нуля ")
						}
					case "*":
						fmt.Printf("%v %s %v = %v\n", intToRoman(arab1), operator, intToRoman(arab2), intToRoman(arab1*arab2))
					case "-":
						if (arab1 - arab2) > 1 {
							fmt.Printf("%v %s %v = %v\n", intToRoman(arab1), operator, intToRoman(arab2), intToRoman(arab1-arab2))
						} else if (arab1 - arab2) < 1 {
							fmt.Print("Ошибка, так как в римской системе нет отрицательных чисел")
						}
					default:
						fmt.Println("Некорректный оператор")
					}
				}
			} else {
				fmt.Print("Ошибка, так как используются одновременно разные системы счисления")
			}
		} else {
			fmt.Print("Ошибка, так как используются одновременно разные системы счисления")
		}
	} else if a == "2" {
		var operator string
		fmt.Print("Введите первое число: ")
		var number1, number2 int

		fmt.Scan(&number1)

		if number1 < 1 {
			fmt.Print(" Превышено минимальное число по заданию")
			return
		} else if number1 > 10 {
			fmt.Print("Превышено число максимума по заданию")
			return
		}

		fmt.Print("Введите второе число: ")

		fmt.Scan(&number2)

		if number2 < 1 {
			fmt.Print("Превышено минимальное число по заданию")
			return
		} else if number2 > 10 {
			fmt.Print("Превышено число максимума по заданию")
			return
		}

		fmt.Print("\n Введите оператор (+ - * /): ")
		fmt.Scan(&operator)
		switch operator {
		case "+":
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1+number2)
		case "/":
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1/number2)
		case "*":
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1*number2)
		case "-":
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1-number2)
		default:
			fmt.Println("Недопустимый оператор")
		}
	} else {
		fmt.Println("Не правильный выбор системы решение калькулятора")
		return
	}

}

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func intToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
