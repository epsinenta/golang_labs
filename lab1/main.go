package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calc(args []float64) []float64 {
	var result []float64
	a, b, c := args[0], args[1], args[2]

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return result
	}

	sqrtDisc := math.Sqrt(discriminant)
	y1 := (-b + sqrtDisc) / (2 * a)
	y2 := (-b - sqrtDisc) / (2 * a)

	if y1 >= 0 {
		sqrtY1 := math.Sqrt(y1)
		result = append(result, sqrtY1, -sqrtY1)
	}

	if y2 >= 0 {
		sqrtY2 := math.Sqrt(y2)
		result = append(result, sqrtY2, -sqrtY2)
	}

	return result
}

func inputCoefficient(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	var coeff float64
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) 
		val, err := strconv.ParseFloat(input, 64)
		if err == nil {
			coeff = val
			break 
		} else {
			fmt.Println("Некорректный ввод, попробуйте снова.")
		}
	}
	return coeff
}

func main() {
	args := os.Args
	var resultArgs []float64

	if len(args) == 4 {
		for i := 1; i <= 3; i++ {
			val, err := strconv.ParseFloat(args[i], 64)
			if err == nil {
				resultArgs = append(resultArgs, val)
			} else {
				fmt.Printf("Аргумент %d некорректен, вводите коэффициенты заново.\n", i)
				break
			}
		}
	}

	if len(resultArgs) != 3 {
		a := inputCoefficient("Введите коэффициент a: ")
		b := inputCoefficient("Введите коэффициент b: ")
		c := inputCoefficient("Введите коэффициент c: ")
		resultArgs = []float64{a, b, c}
	}

	result := calc(resultArgs)
	fmt.Println("Количество корней:", len(result))
	fmt.Print("Корни: ", result)
}
