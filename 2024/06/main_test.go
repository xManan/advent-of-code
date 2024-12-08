package main

import (
    "testing"
    "strings"
)

func TestPart1(t *testing.T) {
    inp := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

    out := part1(strings.Split(inp, "\n"))
    exp := "41"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}

func TestPart2(t *testing.T) {
    inp := ``

    out := part2(strings.Split(inp, "\n"))
    exp := "0"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}
