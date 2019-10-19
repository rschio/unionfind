package unionfind

import "fmt"

// Lener is a collection.
type Lener interface {
	Len() int
}

// Set is a Union Finder (disjoint) data structure
// with Quick-Union and Path Compression.
type Set struct {
	Lener // Lener is the base collection of Set.
	// root stores the root of Lener at position i or if i is its own root,
	// root stores the number os items in its group.
	root []int
}

// New returns a new Set with data embedded. Data should
// not change the order or len.
func New(data Lener) *Set {
	s := Set{Lener: data, root: make([]int, data.Len())}
	for i := 0; i < data.Len(); i++ {
		s.root[i] = -1
	}
	return &s
}

// Find returns the root of element at position i
// or -1 if i is a invalid index.
func (s *Set) Find(i int) int {
	if i < 0 || i >= s.Len() {
		return -1
	}
	if s.root[i] < 0 {
		return i
	}
	// The algorithm only update the root when is
	// necessary to know the root.
	s.root[i] = s.Find(s.root[i])
	return s.root[i]
}

// Union unify i and j under the same root.
func (s *Set) Union(i, j int) error {
	ri := s.Find(i)
	rj := s.Find(j)

	if ri == -1 || rj == -1 {
		return fmt.Errorf("invalid index")
	}
	// i and j already are of same union nothing to do.
	if ri == rj && ri != -1 {
		return nil
	}
	// New root is the root with bigger |size|.
	// Note that the sizes are stored as negative
	// numbers.
	if s.root[rj] < s.root[ri] {
		s.root[rj] += s.root[ri]
		s.root[ri] = rj
		return nil
	}
	s.root[ri] += s.root[rj]
	s.root[rj] = ri
	return nil
}
