package graphs

import (
	"math"
)

/* ---------------- Stoer–Wagner ---------------- */

// EdgeConnectivitySW returns the edge connectivity for an undirected graph.
func EdgeConnectivitySW(g *UndirectedGraph) int {
	n := g.N
	if n <= 1 {
		return 0
	}

	// Weight matrix (int64), summing parallel edges
	w := make([][]int64, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int64, n)
	}
	for _, e := range g.Edges {
		u, v := e[0], e[1]
		if u == v {
			continue
		}
		if u < 0 || v < 0 || u >= n || v >= n {
			panic("EdgeConnectivitySW: индекс вершины вне диапазона")
		}
		w[u][v]++
		w[v][u]++
	}

	deg := make([]int64, n)
	for i := 0; i < n; i++ {
		var s int64
		for j := 0; j < n; j++ {
			s += w[i][j]
		}
		deg[i] = s
	}
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			return 0
		}
	}

	best := int64(math.MaxInt64)
	nn := n

	for nn > 1 {
		added := make([]bool, nn)
		adj := make([]int64, nn)
		prev := -1

		for i := 0; i < nn; i++ {
			sel := -1
			var bestW int64 = -1
			for v := 0; v < nn; v++ {
				if !added[v] && adj[v] > bestW {
					bestW = adj[v]
					sel = v
				}
			}
			if sel == -1 {
				return 0
			}

			added[sel] = true

			if i == nn-1 {
				if adj[sel] < best {
					best = adj[sel]
				}

				for v := 0; v < nn; v++ {
					if v == prev {
						continue
					}
					w[prev][v] += w[sel][v]
					w[v][prev] = w[prev][v]
				}

				if sel != nn-1 {
					for v := 0; v < nn; v++ {
						w[sel][v] = w[nn-1][v]
						w[v][sel] = w[v][nn-1]
					}
				}
				nn--
				break
			}

			for v := 0; v < nn; v++ {
				if !added[v] {
					adj[v] += w[sel][v]
				}
			}
			prev = sel
		}
	}

	if best == int64(math.MaxInt64) {
		return 0
	}
	return int(best)
}
