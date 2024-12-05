package main

import (
    "testing"
    "strings"
)

func TestPart1(t *testing.T) {
    inp := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

    out := part1(strings.Split(inp, "\n"))
    exp := "143"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}

func TestPart2(t *testing.T) {
    inp := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

    out := part2(strings.Split(inp, "\n"))
    exp := "123"
    if out != exp {
        t.Fatalf("inp:\n%s\n\nexpected: %s, got %s", inp, exp, out)
    }
}
