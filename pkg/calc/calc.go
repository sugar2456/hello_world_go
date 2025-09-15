package calc

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}

func Execute() {
	opMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"2", "+", "three"},
		{"5"},
	}
	for _, exexpression := range expressions {
		if len(exexpression) != 3 {
			fmt.Print(exexpression, "不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(exexpression[0])
		if err != nil {
			fmt.Print(exexpression[0], "不正なオペランドです\n")
			continue
		}
		op := exexpression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(op, "不正な演算子です\n")
			continue
		}
		p2, err := strconv.Atoi(exexpression[2])
		if err != nil {
			fmt.Print(exexpression[2], "不正なオペランドです\n")
			continue
		}
		result := opFunc(p1, p2)
		fmt.Printf("%d %s %d = %d\n", p1, op, p2, result)
	}
}
