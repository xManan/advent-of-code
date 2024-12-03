package main

import (
    "fmt"
    "os"
    "advent-of-code/helpers"
    "strconv"
    "unicode"
    _ "strings"
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
    sum := 0
    for _,line := range inp {
        runes := []rune(line)
        for i := 0; i < len(runes); i++ {
            if runes[i] == 'm' && 
            (i+1 < len(runes) && runes[i+1] == 'u') && 
            (i+2 < len(runes) && runes[i+2] == 'l') {
                j := i+3
                if j < len(runes) && runes[j] == '(' {
                    j++
                    x := ""
                    for j < len(runes) && unicode.IsDigit(runes[j]) {
                        x += string(runes[j])
                        j++
                    }
                    if j < len(runes) && runes[j] == ',' {
                        j++
                        y := ""
                        for j < len(runes) && unicode.IsDigit(runes[j]) {
                            y += string(runes[j])
                            j++
                        }
                        if j < len(runes) && runes[j] == ')' {
                            xInt, _ := strconv.Atoi(x)
                            yInt, _ := strconv.Atoi(y)
                            sum += xInt * yInt
                            i = j
                        }
                    }
                }
            }
        }
    }
    return strconv.Itoa(sum)
}

func part2(inp []string) string {
    sum := 0
    disabled := false
    for _,line := range inp {
        runes := []rune(line)
        for i := 0; i < len(runes); i++ {
            if runes[i] == 'm' && 
            (i+1 < len(runes) && runes[i+1] == 'u') && 
            (i+2 < len(runes) && runes[i+2] == 'l') {
                if disabled {
                    continue
                }
                j := i+3
                if j < len(runes) && runes[j] == '(' {
                    j++
                    x := ""
                    for j < len(runes) && unicode.IsDigit(runes[j]) {
                        x += string(runes[j])
                        j++
                    }
                    if j < len(runes) && runes[j] == ',' {
                        j++
                        y := ""
                        for j < len(runes) && unicode.IsDigit(runes[j]) {
                            y += string(runes[j])
                            j++
                        }
                        if j < len(runes) && runes[j] == ')' {
                            xInt, _ := strconv.Atoi(x)
                            yInt, _ := strconv.Atoi(y)
                            sum += xInt * yInt
                            i = j
                        }
                    }
                }
            } else if runes[i] == 'd' {
                if i+3 < len(runes) && line[i:i+4] == "do()" {
                    disabled = false
                } else if i+6 < len(runes) && line[i:i+7] == "don't()" {
                    disabled = true
                }
            }
        }
    }
    return strconv.Itoa(sum)
}
