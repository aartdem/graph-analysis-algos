package graphs

import (
	"math"
)

type Edge struct {
	to  int
	rev int
	cap int64
}

type PR struct { // push-relabel
	n    int
	s, t int
	adj  [][]Edge

	h      []int
	excess []int64
	cur    []int

	q       []int
	inQueue []bool
	qHead   int
}

// AddEdge adds undirected edge with capacity cap
func (pr *PR) AddEdge(u, v int, cap int64) {
	if u == v || cap == 0 {
		return
	}

	fwd := Edge{to: v, rev: len(pr.adj[v]), cap: cap}
	rev := Edge{to: u, rev: len(pr.adj[u]), cap: 0}
	pr.adj[u] = append(pr.adj[u], fwd)
	pr.adj[v] = append(pr.adj[v], rev)
}

func (pr *PR) enqueue(v int) {
	if v != pr.s && v != pr.t && !pr.inQueue[v] && pr.excess[v] > 0 {
		pr.inQueue[v] = true
		pr.q = append(pr.q, v)
	}
}

func (pr *PR) pop() (int, bool) {
	for pr.qHead < len(pr.q) {
		v := pr.q[pr.qHead]
		pr.qHead++
		pr.inQueue[v] = false
		if pr.excess[v] > 0 && v != pr.s && v != pr.t {
			return v, true
		}
	}
	return -1, false
}

func (pr *PR) push(u int, ei int) {
	e := pr.adj[u][ei]
	if e.cap == 0 || pr.h[u] != pr.h[e.to]+1 {
		return
	}

	delta := pr.excess[u]
	if e.cap < delta {
		delta = e.cap
	}
	// decrease residual on the forward edge
	pr.adj[u][ei].cap -= delta
	// increase residual on the reverse edge
	rev := pr.adj[u][ei].rev
	pr.adj[e.to][rev].cap += delta

	pr.excess[u] -= delta
	pr.excess[e.to] += delta

	pr.enqueue(e.to)
}

func (pr *PR) relabel(u int) {
	// new height = 1 + min{ h[v] | (u->v) has positive residual }
	minH := math.MaxInt32
	for _, e := range pr.adj[u] {
		if e.cap > 0 && pr.h[e.to] < minH {
			minH = pr.h[e.to]
		}
	}
	if minH < math.MaxInt32 {
		pr.h[u] = minH + 1
	} else {
		// isolated in the residual graph — set a large "infinite" height
		pr.h[u] = math.MaxInt32 / 4
	}
	// reset current pointer to scan edges again
	pr.cur[u] = 0
}

func (pr *PR) discharge(u int) {
	for pr.excess[u] > 0 {
		if pr.cur[u] == len(pr.adj[u]) {
			pr.relabel(u)
			continue
		}
		pr.push(u, pr.cur[u])
		if pr.adj[u][pr.cur[u]].cap == 0 {
			// edge is exhausted — move to the next
			pr.cur[u]++
		} else {
			// if after push the edge still has positive residual,
			// but the height condition no longer holds — advance anyway
			if pr.h[u] != pr.h[pr.adj[u][pr.cur[u]].to]+1 {
				pr.cur[u]++
			}
		}
	}
}

func (pr *PR) MaxFlow() int64 {
	n := pr.n
	// Preflow initialization
	pr.h[pr.s] = n
	for i := range pr.adj[pr.s] {
		e := pr.adj[pr.s][i]
		if e.cap > 0 {
			flow := e.cap
			pr.adj[pr.s][i].cap -= flow
			rev := pr.adj[pr.s][i].rev
			pr.adj[e.to][rev].cap += flow

			pr.excess[e.to] += flow
			pr.excess[pr.s] -= flow
			pr.enqueue(e.to)
		}
	}

	// Main loop: process active vertices in FIFO order
	for {
		u, ok := pr.pop()
		if !ok {
			break
		}
		pr.discharge(u)
		// If after discharge the vertex is still active (excess>0), re-enqueue it
		pr.enqueue(u)
	}

	if pr.excess[pr.t] < 0 {
		return 0
	}
	return pr.excess[pr.t]
}

// NewPR creates an empty network with n vertices.
func NewPR(n int, s, t int) *PR {
	return &PR{
		n:       n,
		s:       s,
		t:       t,
		adj:     make([][]Edge, n),
		h:       make([]int, n),
		excess:  make([]int64, n),
		cur:     make([]int, n),
		inQueue: make([]bool, n),
	}
}
