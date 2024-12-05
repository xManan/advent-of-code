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
	var grid [][]string
	for _, line := range inp {
		grid = append(grid, strings.Split(line, ""))
	}
	wordCount := 0
	word := "XMAS"
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] != string(word[0]) {
				continue
			}
			dirs := [][]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
				{1, 1},
				{1, -1},
				{-1, 1},
				{-1, -1},
			}
			for _, dir := range dirs {
				i, j, w := x+(dir[0]*1), y+(dir[1]*1), 1
				for i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && w < len(word) {
					if grid[i][j] == string(word[w]) {
						if w == len(word)-1 {
							wordCount++
							// fmt.Printf("[%d, %d] [%d, %d]\n", x, y, i, j)
							break
						}
						i += dir[0] * 1
						j += dir[1] * 1
						w++
					} else {
						break
					}
				}
			}
		}
	}
	return strconv.Itoa(wordCount)
}

func part2(inp []string) string {
	var grid [][]string
	for _, line := range inp {
		grid = append(grid, strings.Split(line, ""))
	}
    wordCount := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == "A" {
				if ((grid[i-1][j-1] == "M" && grid[i+1][j+1] == "S") || (grid[i-1][j-1] == "S" && grid[i+1][j+1] == "M")) &&
				   ((grid[i+1][j-1] == "M" && grid[i-1][j+1] == "S") || (grid[i+1][j-1] == "S" && grid[i-1][j+1] == "M")) {
                    wordCount++
				}
			}
		}
	}
	return strconv.Itoa(wordCount)
}
