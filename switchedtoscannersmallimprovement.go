package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type measurement float32

func main() {
    file, err := os.Open("/home/dom/development/go/1brc/measurements.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    measurements := make(map[string][]measurement)

    for scanner.Scan() {

        city := scanner.Text()
        split := strings.Split(city, ";")
        cityname := split[0]
        temp := split[1][:len(split[1])-1]

        n, err := strconv.ParseFloat(temp, 32)
        if err != nil {
            fmt.Println(err)
            return
        }

        temperature := measurement(n)
        measurements[cityname] = append(measurements[cityname], temperature)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
        return
    }

    keys := make([]string, 0, len(measurements))

    for k := range measurements {
        keys = append(keys, k)
    }

    sort.Strings(keys)
    fmt.Printf("{")

    for _, city := range keys {
        var sum measurement

        for _, v := range measurements[city] {
            sum += v
        }
        mean := sum / measurement(len(measurements[city]))

        if city == "İzmir" {
            fmt.Printf("%s=%.1f/%.1f/%.1f}", city, slices.Min(measurements[city]), mean, slices.Max(measurements[city]) )
            continue
        }

        fmt.Printf("%s=%.1f/%.1f/%.1f, ", city, slices.Min(measurements[city]), mean, slices.Max(measurements[city]) )

    }

}
