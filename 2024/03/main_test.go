package main

import (
    "testing"
    "strings"
)

func TestPart1(t *testing.T) {
    inp := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

    out := part1(strings.Split(inp, "\n"))
    exp := "161"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}

func TestPart2(t *testing.T) {
    inp := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

    out := part2(strings.Split(inp, "\n"))
    exp := "48"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}
