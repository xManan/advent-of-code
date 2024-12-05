package main

import (
    "testing"
    "strings"
)

func TestPart1(t *testing.T) {
    inp := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

    out := part1(strings.Split(inp, "\n"))
    exp := "18"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}

func TestPart2(t *testing.T) {
    inp := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

    out := part2(strings.Split(inp, "\n"))
    exp := "9"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}
