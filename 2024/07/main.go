package main

import (
	"advent-of-code/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
	_ "unicode"
)

func main() {
	inp, err := helpers.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
	out := part2(inp)
	fmt.Print(out)
}

func part1(inp []string) string {
    sum := 0
    for _, line := range inp {
        lineSplit := strings.Split(line, ":")
        ans, _ := strconv.Atoi(lineSplit[0])
        operands := strings.Split(lineSplit[1], " ")[1:]
        possible := false

        a, _ := strconv.Atoi(operands[0])
        operands = operands[1:]
        recCalcPart1(a, "+", operands, ans, &possible)
        if possible {
            sum += ans
            continue
        }
        recCalcPart1(a, "*", operands, ans, &possible)
        if possible {
            sum += ans
            continue
        }
    }
	return strconv.Itoa(sum)
}

func recCalcPart1(a int, operator string, operands []string, expected int, possible *bool) {
    if *possible {
        return
    }
    if a == expected {
        *possible = true
        return
    }
    if len(operands) == 0 {
        return
    }
    b, _ := strconv.Atoi(operands[0])
    operands = operands[1:]
    if operator == "+" {
        recCalcPart1(a+b, "+", operands, expected, possible)
        recCalcPart1(a+b, "*", operands, expected, possible)
    } else if operator == "*" {
        recCalcPart1(a*b, "+", operands, expected, possible)
        recCalcPart1(a*b, "*", operands, expected, possible)
    }
}

func part2(inp []string) string {
    sum := 0
    for _, line := range inp {
        lineSplit := strings.Split(line, ":")
        ans, _ := strconv.Atoi(lineSplit[0])
        operands := strings.Split(lineSplit[1], " ")[1:]
        possible := false

        a, _ := strconv.Atoi(operands[0])
        operands = operands[1:]
        recCalcPart2(a, "+", operands, ans, &possible)
        if possible {
            sum += ans
            continue
        }
        recCalcPart2(a, "*", operands, ans, &possible)
        if possible {
            sum += ans
            continue
        }
        recCalcPart2(a, "||", operands, ans, &possible)
        if possible {
            sum += ans
            continue
        }
    }
	return strconv.Itoa(sum)
}

func recCalcPart2(a int, operator string, operands []string, expected int, possible *bool) {
    if *possible {
        return
    }
    if a == expected {
        *possible = true
        return
    }
    if len(operands) == 0 {
        return
    }
    b, _ := strconv.Atoi(operands[0])
    operands = operands[1:]
    if operator == "+" {
        recCalcPart2(a+b, "+", operands, expected, possible)
        recCalcPart2(a+b, "*", operands, expected, possible)
        recCalcPart2(a+b, "||", operands, expected, possible)
    } else if operator == "*" {
        recCalcPart2(a*b, "+", operands, expected, possible)
        recCalcPart2(a*b, "*", operands, expected, possible)
        recCalcPart2(a*b, "||", operands, expected, possible)
    } else if operator == "||" {
        num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
        recCalcPart2(num, "+", operands, expected, possible)
        recCalcPart2(num, "*", operands, expected, possible)
        recCalcPart2(num, "||", operands, expected, possible)
    }
}
