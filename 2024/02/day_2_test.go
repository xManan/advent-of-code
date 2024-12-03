package main

import (
    "testing"
    "strings"
)

func TestPart1(t *testing.T) {
    inp := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

    out := part1(strings.Split(inp, "\n"))
    exp := "2"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}

func TestPart2(t *testing.T) {
    inp := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

    out := part2(strings.Split(inp, "\n"))
    exp := "4"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}
