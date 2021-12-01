package templates

func MainTemplate() string {
    return `
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main () {
    fileContents, err := readFile("input")
    if err != nil {
        panic(err)
    }

    fmt.Printf("File contents: %v\n", fileContents)
}

func readFile(path string) ( []string, error) {
    buf, err := os.Open(path)
    if err != nil {
        return  []string{}, err
    }

    var lines []string 
    scanLines := bufio.NewScanner(buf)
    for scanLines.Scan() {
        lines = append(lines, scanLines.Text())
    }

    err = buf.Close()
    if err != nil {
        return  []string{}, err
    }

    return lines, nil
}

`
}
