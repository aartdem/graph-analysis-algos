package graphs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadMTX(filename string) (*UndirectedGraph, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var n, m, nnz int
	lineNo := 0

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 || line[0] == '%' {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 3 {
			return nil, fmt.Errorf("unexpected format")
		}
		n, _ = strconv.Atoi(parts[0])
		m, _ = strconv.Atoi(parts[1])
		nnz, _ = strconv.Atoi(parts[2])
		lineNo++
		break
	}

	if n != m {
		return nil, fmt.Errorf("matrix is not square (%d x %d)", n, m)
	}

	g := NewUndirectedGraph(n)

	readEdges := 0
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 || line[0] == '%' {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])

		u--
		v--
		if u < 0 || v < 0 || u >= n || v >= n {
			return nil, fmt.Errorf("ids out of range %d %d", u, v)
		}

		if u != v {
			g.AddEdge(u, v)
		}
		readEdges++
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	if readEdges < nnz {
		fmt.Printf("warn: expected %d edges, read %d\n", nnz, readEdges)
	}

	return g, nil
}
