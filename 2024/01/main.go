package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "sort"
    "advent-of-code/helpers"
)

func main() {
    inp, err := helpers.ReadFile("input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err)
    }
    out := part1(inp)
    fmt.Print(out)
}

func part1(inp []string) string {
    var lefts []int
    var rights []int
    for _, line := range inp {
        split := strings.Split(line, "   ")
        left, _ := strconv.Atoi(split[0])
        right, _ := strconv.Atoi(split[1])

        lefts = append(lefts, left)
        rights = append(rights, right)
    }
    sort.Ints(lefts)
    sort.Ints(rights)
    distanceSum := 0
    for i, left := range lefts {
        diff := left - rights[i]
        if diff < 0 {
            distanceSum += -1 * diff
        } else {
            distanceSum += diff
        }
    }
    return strconv.Itoa(distanceSum)
}

func part2(inp []string) string {
    var lefts []int
    var rights []int
    for _, line := range inp {
        split := strings.Split(line, "   ")
        if len(split) > 1 {
            left, _ := strconv.Atoi(split[0])
            right, _ := strconv.Atoi(split[1])

            lefts = append(lefts, left)
            rights = append(rights, right)
        }
    }
    rightFreqMap := make(map[int]int)
    for _,n := range rights {
        count, exists := rightFreqMap[n]
        if exists {
            rightFreqMap[n] = count+1
        } else {
            rightFreqMap[n] = 1
        }
    }

    sum := 0
    for _,n := range lefts {
        count, exists := rightFreqMap[n]
        if exists {
            sum += n*count
        } 
    }
    return strconv.Itoa(sum)
}




