package graphs

import (
	"testing"
)

/* Graph builders */

func buildPath(n int) *UndirectedGraph {
	g := NewUndirectedGraph(n)
	for i := 0; i+1 < n; i++ {
		g.AddEdge(i, i+1)
	}
	return g
}

func buildCycle(n int) *UndirectedGraph {
	g := NewUndirectedGraph(n)
	if n == 0 {
		return g
	}
	for i := 0; i < n; i++ {
		g.AddEdge(i, (i+1)%n)
	}
	return g
}

func buildComplete(n int) *UndirectedGraph {
	g := NewUndirectedGraph(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g.AddEdge(i, j)
		}
	}
	return g
}

func buildCompleteBipartite(a, b int) *UndirectedGraph {
	n := a + b
	g := NewUndirectedGraph(n)
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			g.AddEdge(i, a+j)
		}
	}
	return g
}

func buildTwoCyclesWithBridge(a, b int) *UndirectedGraph {
	n := a + b
	g := NewUndirectedGraph(n)
	for i := 0; i < a; i++ { // cycle A
		g.AddEdge(i, (i+1)%a)
	}
	for i := 0; i < b; i++ { // cycle B
		u := a + i
		v := a + ((i + 1) % b)
		g.AddEdge(u, v)
	}
	if a > 0 && b > 0 {
		g.AddEdge(0, a) // bridge
	}
	return g
}

func buildParallelEdges(k int) *UndirectedGraph {
	g := NewUndirectedGraph(2)
	for i := 0; i < k; i++ {
		g.AddEdge(0, 1)
	}
	return g
}

/* Tests */

type edgeConnAlgo func(*UndirectedGraph) int

func runEdgeConnectivityTests(t *testing.T, name string, algo edgeConnAlgo) {
	t.Helper()
	t.Parallel()

	tests := []struct {
		name  string
		build func() *UndirectedGraph
		want  int
	}{
		{"Path_n2", func() *UndirectedGraph { return buildPath(2) }, 1},
		{"Path_n5", func() *UndirectedGraph { return buildPath(5) }, 1},

		{"Cycle_n3", func() *UndirectedGraph { return buildCycle(3) }, 2},
		{"Cycle_n8", func() *UndirectedGraph { return buildCycle(8) }, 2},

		{"Complete_K3", func() *UndirectedGraph { return buildComplete(3) }, 2},
		{"Complete_K5", func() *UndirectedGraph { return buildComplete(5) }, 4},

		{"K33_Bipartite", func() *UndirectedGraph { return buildCompleteBipartite(3, 3) }, 3},
		{"K_2_5_Bipartite", func() *UndirectedGraph { return buildCompleteBipartite(2, 5) }, 2},

		{"TwoCyclesWithBridge_5_6", func() *UndirectedGraph { return buildTwoCyclesWithBridge(5, 6) }, 1},

		{"ParallelEdges_k1", func() *UndirectedGraph { return buildParallelEdges(1) }, 1},
		{"ParallelEdges_k4", func() *UndirectedGraph { return buildParallelEdges(4) }, 4},

		{"BigGraph1", func() *UndirectedGraph { g, _ := LoadMTX("../data/san1_100.mtx"); return g }, 40},
		{"BigGraph2", func() *UndirectedGraph { g, _ := LoadMTX("../data/san201_300.mtx"); return g }, 33},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(name+"/"+tc.name, func(t *testing.T) {
			t.Parallel()
			g := tc.build()
			got := algo(g)
			if got != tc.want {
				t.Fatalf("%s: got %d, want %d", name, got, tc.want)
			}
		})
	}
}

func TestEdgeConnectivityMaxFlow(t *testing.T) {
	runEdgeConnectivityTests(t, "EdgeConnectivity-MaxFlow", func(g *UndirectedGraph) int { return EdgeConnectivity(g) })
}

func TestEdgeConnectivityStoerWagner(t *testing.T) {
	runEdgeConnectivityTests(t, "EdgeConnectivity-StoerWagner", func(g *UndirectedGraph) int { return EdgeConnectivitySW(g) })
}
