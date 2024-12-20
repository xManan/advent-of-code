package main

import (
    "testing"
    "strings"
)

func TestPart1(t *testing.T) {
    inp := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

    out := part1(strings.Split(inp, "\n"))
    exp := "3749"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}

func TestPart2(t *testing.T) {
    inp := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

    out := part2(strings.Split(inp, "\n"))
    exp := "11387"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}
