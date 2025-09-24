package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	u, v int
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run main.go input.mtx output.mtx startVertex endVertex")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	start, err1 := strconv.Atoi(os.Args[3])
	end, err2 := strconv.Atoi(os.Args[4])

	if err1 != nil || err2 != nil || start <= 0 || end < start {
		log.Fatalf("invalid range: start=%d end=%d", start, end)
	}

	in, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("cannot open file: %v", err)
	}
	defer in.Close()

	var edges []Edge
	var header string
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "%") {
			if strings.HasPrefix(line, "%%MatrixMarket") {
				header = line
			}
			continue
		}

		var u, v, w int
		n, _ := fmt.Sscanf(line, "%d %d %d", &u, &v, &w)
		if n >= 2 {
			if u >= start && u <= end && v >= start && v <= end {
				edges = append(edges, Edge{u - start + 1, v - start + 1})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("cannot create file: %v", err)
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	defer writer.Flush()

	if header != "" {
		fmt.Fprintln(writer, header)
	} else {
		fmt.Fprintln(writer, "%%MatrixMarket matrix coordinate pattern symmetric")
	}

	fmt.Fprintf(writer, "%% Reduced to vertices from %d to %d\n", start, end)

	numVertices := end - start + 1
	fmt.Fprintf(writer, "%d %d %d\n", numVertices, numVertices, len(edges))

	for _, e := range edges {
		fmt.Fprintf(writer, "%d %d\n", e.u, e.v)
	}
}
