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
    rules := make(map[string][]string)
    var pages [][]string
    readPages := false
    for _,line := range inp {
        if line == "" {
            readPages = true
            continue
        }
        if readPages {
            pages = append(pages, strings.Split(line, ","))    
        } else {
            pageOrder := strings.Split(line, "|")   
            rules[pageOrder[0]] = append(rules[pageOrder[0]], pageOrder[1])
        }
    }
    sum := 0
    for _,page := range pages {
        rightOrder := true
        outer:
        for i := 0; i < len(page) - 1; i++ {
            for j := i+1; j < len(page); j++ {
                rule, ok := rules[page[j]]
                if !ok {
                    continue
                }
                for _,p := range rule {
                    if page[i] == p {
                        rightOrder = false
                        break outer
                    }
                }
            }
        }
        if rightOrder {
            middlePage, _ := strconv.Atoi(page[len(page)/2])
            sum += middlePage
        }
    }
	return strconv.Itoa(sum)
}

func part2(inp []string) string {
    rules := make(map[string][]string)
    var pages [][]string
    readPages := false
    for _,line := range inp {
        if line == "" {
            readPages = true
            continue
        }
        if readPages {
            pages = append(pages, strings.Split(line, ","))    
        } else {
            pageOrder := strings.Split(line, "|")   
            rules[pageOrder[0]] = append(rules[pageOrder[0]], pageOrder[1])
        }
    }
    sum := 0
    for _,page := range pages {
        rightOrder := true
        for i := 0; i < len(page) - 1; i++ {
            for j := i+1; j < len(page); j++ {
                rule, ok := rules[page[j]]
                if !ok {
                    continue
                }
                for _,p := range rule {
                    if page[i] == p {
                        rightOrder = false
                        tmp := page[i]
                        page[i] = page[j]
                        page[j] = tmp
                    }
                }
            }
        }
        if !rightOrder {
            middlePage, _ := strconv.Atoi(page[len(page)/2])
            sum += middlePage
        }
    }
	return strconv.Itoa(sum)
}
