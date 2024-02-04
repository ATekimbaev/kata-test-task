package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение в виде строки: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	input = strings.TrimSpace(input)

	a, operator, b, err := parseInput(input)
	if err != nil {
		panic(err.Error())
	}

	result, err := calculate(a, operator, b)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}

func parseInput(input string) (int, string, int, error) {
	var a, b int
	var operator string

	n, err := fmt.Sscanf(input, "%d %s %d", &a, &operator, &b)
	if err != nil || n != 3 {
		return 0, "", 0, fmt.Errorf("Некорректный формат ввода: %s", input)
	}

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return 0, "", 0, fmt.Errorf("Неподдерживаемая операция: %s", operator)
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		return 0, "", 0, fmt.Errorf("Числа должны быть в диапазоне от 1 до 10 включительно")
	}

	return a, operator, b, nil
}

func calculate(a int, operator string, b int) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Неподдерживаемая операция: %s", operator)
	}
}
