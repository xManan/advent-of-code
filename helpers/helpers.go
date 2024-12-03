package helpers

import (
    "os"
    "strings"
    _ "fmt"
)

func ReadFile(fileName string) ([]string, error) {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        return nil, err
    }
    lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
    return lines, nil
}


