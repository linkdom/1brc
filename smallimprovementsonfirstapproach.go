package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type measurement float32

func extractMeasurements(reader *bufio.Reader, measurements map[string][]measurement) (b map[string][]measurement, a error) {

    for {
        city, err := reader.ReadString('\n')

        if err == io.EOF {
            return measurements, nil
        }
        
        if err != nil {
            fmt.Println(err)
            return nil, err
        }

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

}

func main() {
    file, err := os.Open("/home/dom/development/go/1brc/measurements.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    measurements := make(map[string][]measurement)
    reader := bufio.NewReader(file)
    measurements, err = extractMeasurements(reader, measurements)
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
