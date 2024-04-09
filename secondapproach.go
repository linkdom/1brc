package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type measurement struct {
    count int
    min float64
    max float64
    sum float64
}

func extractMeasurements(reader *bufio.Reader, measurements map[string]measurement) (b map[string]measurement, a error) {

    for {
        city, err := reader.ReadString('\n')

        if err == io.EOF {
            return measurements, nil
        }
        
        if err != nil {
            fmt.Println(err)
            return nil, err
        }

        split := strings.Split(strings.TrimSpace(city), ";")
        cityname := split[0]
        temp := strings.Split(split[1], "\n")[0]

        n, err := strconv.ParseFloat(temp, 32)
        if err != nil {
            fmt.Println(err)
            return
        }

        if measurement, ok := measurements[cityname]; ok {
            measurement.count++
            measurement.sum += n

            if n < measurement.min {
                measurement.min = n
            }

            if n > measurement.max {
                measurement.max = n
            }

            measurements[cityname] = measurement
            continue
        }

        measurements[cityname] = measurement{
            count: 1,
            min: n,
            max: n,
            sum: n,
        }

    }

}

func main() {
    file, err := os.Open("/home/dom/programming/1brc/measurements.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    measurements := make(map[string]measurement)
    reader := bufio.NewReader(file)
    measurements, err = extractMeasurements(reader, measurements)
    keys := make([]string, 0, len(measurements))

    for k := range measurements {
        keys = append(keys, k)
    }

    sort.Strings(keys)
    fmt.Printf("{")

    for _, city := range keys {
        mean := measurements[city].sum / float64(measurements[city].count)

        if city == "İzmir" {
            fmt.Printf("%s=%.1f/%.1f/%.1f", city, measurements[city].min, mean, measurements[city].max)
            continue
        }

        fmt.Printf("%s=%.1f/%.1f/%.1f, ", city, measurements[city].min, mean, measurements[city].max)

    }

    fmt.Printf("}")

}
