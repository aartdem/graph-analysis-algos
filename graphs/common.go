package graphs

type UndirectedGraph struct {
	N     int
	Edges [][2]int
}

func NewUndirectedGraph(n int) *UndirectedGraph {
	return &UndirectedGraph{N: n}
}

func (g *UndirectedGraph) AddEdge(u, v int) {
	g.Edges = append(g.Edges, [2]int{u, v})
}
