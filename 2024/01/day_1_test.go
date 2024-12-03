package main

import (
    "testing"
    "strings"
)

func TestPart1(t *testing.T) {
    inp := `3   4
4   3
2   5
1   3
3   9
3   3`

    out := part1(strings.Split(inp, "\n"))
    exp := "11"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}

func TestPart2(t *testing.T) {
    inp := `3   4
4   3
2   5
1   3
3   9
3   3`

    out := part2(strings.Split(inp, "\n"))
    exp := "31"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}




