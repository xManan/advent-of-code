package main

import (
    "fmt"
    "os"
    "advent-of-code/helpers"
    "strconv"
    "strings"
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
    safeReports := 0
    for _, report := range inp {
        levels := strings.Split(report, " ")
        if isReportSafe(levels) {
            safeReports += 1
        }
    }
    return strconv.Itoa(safeReports)
}

func part2(inp []string) string {
    safeReports := 0
    for _, report := range inp {
        levels := strings.Split(report, " ")
        if isReportSafe(levels) {
            safeReports += 1
        } else {
            for i := range levels {
                newReportLevels := make([]string, len(levels))
                copy(newReportLevels, levels)
                newReportLevels = append(newReportLevels[:i], newReportLevels[i+1:]...)
                if isReportSafe(newReportLevels) {
                    safeReports += 1
                    break
                }
            }
        }
    }
    return strconv.Itoa(safeReports)
}

func isReportSafe(levels []string) bool {
    initialChange := 0
    safe := true
    for i := 0; i < len(levels); i++ {
        if i >= len(levels) - 1 {
            break
        }
        levelInt, _ := strconv.Atoi(levels[i])
        levelNextInt, _ := strconv.Atoi(levels[i+1])
        diff := levelInt - levelNextInt
        var change int
        if diff < 0 {
            change = 1
            diff = -1 * diff
        } else {
            change = -1
        }
        if initialChange == 0 {
            initialChange = change
        } else {
            if change != initialChange {
                safe = false
                break
            }
        }
        if diff < 1 || diff > 3 {
            safe = false
            break
        }
    }
    return safe;
}


