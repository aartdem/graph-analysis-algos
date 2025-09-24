package graphs

import (
	"testing"
)

func net(n, s, t int) *PR {
	return NewPR(n, s, t)
}

func TestClassical23(t *testing.T) {
	pr := net(6, 0, 5)
	pr.AddEdge(0, 1, 16)
	pr.AddEdge(0, 2, 13)
	pr.AddEdge(1, 2, 10)
	pr.AddEdge(2, 1, 4)
	pr.AddEdge(1, 3, 12)
	pr.AddEdge(3, 2, 9)
	pr.AddEdge(2, 4, 14)
	pr.AddEdge(4, 3, 7)
	pr.AddEdge(3, 5, 20)
	pr.AddEdge(4, 5, 4)

	got := pr.MaxFlow()
	if got != 23 {
		t.Fatalf("MaxFlow() = %d, want 23", got)
	}
}

func TestNoPath(t *testing.T) {
	// No path from s to t — flow is 0
	pr := net(4, 0, 3)
	pr.AddEdge(0, 1, 5)
	pr.AddEdge(2, 3, 5)
	// no connection between {0,1} and {2,3}
	got := pr.MaxFlow()
	if got != 0 {
		t.Fatalf("MaxFlow() = %d, want 0", got)
	}
}

func TestParallelEdges(t *testing.T) {
	pr := net(3, 0, 2)
	pr.AddEdge(0, 1, 3)
	pr.AddEdge(0, 1, 4) // parallel edge
	pr.AddEdge(1, 2, 10)
	got := pr.MaxFlow()
	if got != 7 {
		t.Fatalf("MaxFlow() = %d, want 7", got)
	}
}

func TestBottleneck(t *testing.T) {
	// Bottleneck on a single edge
	pr := net(4, 0, 3)
	pr.AddEdge(0, 1, 100)
	pr.AddEdge(1, 2, 1) // bottleneck
	pr.AddEdge(2, 3, 100)
	got := pr.MaxFlow()
	if got != 1 {
		t.Fatalf("MaxFlow() = %d, want 1", got)
	}
}

func TestMaxFlow_SmallGrid(t *testing.T) {
	// Small "grid":
	// s=0, t=5
	// 0 -> 1,2
	// 1 -> 3
	// 2 -> 3,4
	// 4 -> 3
	// 3,4 -> 5
	// All edges have capacity 1
	// Intuitively, two disjoint paths to the sink -> answer 2
	pr := net(6, 0, 5)
	add := func(u, v int) { pr.AddEdge(u, v, 1) }

	add(0, 1)
	add(0, 2)
	add(1, 3)
	add(2, 3)
	add(2, 4)
	add(4, 3)
	add(3, 5)
	add(4, 5)

	got := pr.MaxFlow()
	if got != 2 {
		t.Fatalf("MaxFlow() = %d, want 2", got)
	}
}

func TestZeroCapacityEdges(t *testing.T) {
	pr := net(4, 0, 3)
	pr.AddEdge(0, 1, 0) // zero
	pr.AddEdge(0, 2, 5)
	pr.AddEdge(2, 1, 5)
	pr.AddEdge(1, 3, 5)
	got := pr.MaxFlow()
	if got != 5 {
		t.Fatalf("MaxFlow() = %d, want 5", got)
	}
}

func TestSourceSinkSelfLoopsIgnored(t *testing.T) {
	pr := net(3, 0, 2)
	pr.AddEdge(0, 0, 100)
	pr.AddEdge(1, 1, 100)
	pr.AddEdge(2, 2, 100)
	pr.AddEdge(0, 1, 4)
	pr.AddEdge(1, 2, 3)
	got := pr.MaxFlow()
	if got != 3 {
		t.Fatalf("MaxFlow() = %d, want 3", got)
	}
}
