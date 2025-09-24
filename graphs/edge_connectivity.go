package graphs

import "math"

// builds a unit-capacity network for an undirected graph
func buildUnitCapacityNetwork(g *UndirectedGraph, s, t int) *PR {
	pr := NewPR(g.N, s, t)
	for _, e := range g.Edges {
		u, v := e[0], e[1]
		if u == v {
			continue
		}
		pr.AddEdge(u, v, 1)
		pr.AddEdge(v, u, 1)
	}
	return pr
}

// EdgeConnectivity returns global edge connectivity
func EdgeConnectivity(g *UndirectedGraph) int {
	if g.N <= 1 {
		return 0
	}

	deg := make([]int, g.N)
	for _, e := range g.Edges {
		u, v := e[0], e[1]
		if u == v {
			continue
		}
		deg[u]++
		deg[v]++
	}
	best := math.MaxInt32
	minDeg := math.MaxInt32
	for _, d := range deg {
		if d < minDeg {
			minDeg = d
		}
	}
	if minDeg < best {
		best = minDeg
	}

	for s := 0; s < g.N; s++ {
		if best == 0 {
			break
		}
		for t := s + 1; t < g.N; t++ {
			pr := buildUnitCapacityNetwork(g, s, t)
			flow := pr.MaxFlow()
			if int(flow) < best {
				best = int(flow)
				if best == 0 {
					return 0
				}
			}
		}
	}
	return best
}
