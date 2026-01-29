package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//numbers1_6 := numberInput(6)
	numbers2_7 := numberInput(4)

	//task1_6(numbers1_6[0], numbers1_6[1], numbers1_6[2], numbers1_6[3], numbers1_6[4], numbers1_6[5])
	task2_7(numbers2_7[0], numbers2_7[1], numbers2_7[2], numbers2_7[3])
}
func numberInput(n int) []int {
	reader := bufio.NewReader(os.Stdin)
	var numbers = make([]int, n)

	for i := 0; i < n; i++ {
		for {
			fmt.Printf("Введите число %d: ", i+1)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			var number, err = strconv.Atoi(input)
			if err == nil {
				numbers[i] = number
				break
			} else {
				fmt.Println("Ошибка: введено некорректное число. Попробуйте ещё раз.")
			}
		}
	}
	return numbers
}

// found x,y in system of equations with coefficient
// ax + by = c <=> dx + ey = f

func task1_6(a, b, c, d, e, f int) {

	var delta = a*e - b*d
	var x = c*e - b*f/delta
	var y = a*f - c*d/delta
	fmt.Println("X =", x)
	fmt.Print("Y = ", y)
}

// Existence of square
func task2_7(a, b, c, d int) bool {

	isSquare := false
	if a == b && b == c && c == d {
		isSquare = true
		fmt.Print("Квадрат")
	} else {
		isSquare = false
		fmt.Print("Не может быть квадратом")
	}
	return isSquare
}
