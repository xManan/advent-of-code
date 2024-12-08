package main

import (
	"advent-of-code/helpers"
	"fmt"
	"os"
	"strconv"
	_ "strings"
	_ "unicode"
)

func main() {
	inp, err := helpers.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
	out := part1(inp)
	fmt.Print(out)
}

type coord struct {
    i int
    j int
}

func part1(inp []string) string {
    var guardPos coord
    var guardDir coord
    outer:
    for i, line := range inp {
        for j, c := range line {
            if c != '.' && c != '#' {
                guardPos = coord{i: i, j: j}
                switch c {
                case '>':
                    guardDir = coord{i: 0, j: 1}
                case 'V':
                    guardDir = coord{i: 1, j: 0}
                case '<':
                    guardDir = coord{i: 0, j: -1}
                case '^':
                    guardDir = coord{i: -1, j: 0}
                }
                break outer
            }
        }
    }
    distinctPoints := make(map[coord]bool)
    for {
        distinctPoints[coord{i: guardPos.i, j: guardPos.j}] = true
        nextPos := coord{i: guardPos.i + guardDir.i, j: guardPos.j + guardDir.j}
        if nextPos.i >= len(inp) || nextPos.j >= len(inp[0]) || nextPos.i < 0 || nextPos.j < 0 {
            break
        }
        if inp[nextPos.i][nextPos.j] == '#' {
            tmp := guardDir.i
            guardDir.i = guardDir.j
            guardDir.j = -tmp
        } else {
            guardPos = nextPos
        }
    }
	return strconv.Itoa(len(distinctPoints))
}

func part2(inp []string) string {
	return strconv.Itoa(0)
}
