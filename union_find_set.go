package algoutil

type UnionFindSet struct {
	Rank   []int
	Parent []int
	Size   []int
}

func NewUnionFindSet(size int) *UnionFindSet {
	set := &UnionFindSet{}
	set.Rank = make([]int, size)
	set.Size = make([]int, size)
	set.Parent = make([]int, size)
	for i := 0; i < size; i++ {
		set.Rank[i] = 1
		set.Size[i] = 1
		set.Parent[i] = i
	}
	return set
}

func (s *UnionFindSet) Find(element int) int {
	if s.Parent[element] == element {
		return element
	}
	s.Parent[element] = s.Find(s.Parent[element])
	return s.Parent[element]
}

func (s *UnionFindSet) Union(e1, e2 int) {
	root1, root2 := s.Find(e1), s.Find(e2)
	if root1 == root2 {
		return
	}
	var unionRoot, beUnioned int
	if s.Rank[root1] > s.Rank[root2] {
		unionRoot, beUnioned = root1, root2
	} else if s.Rank[root1] < s.Rank[root2] {
		unionRoot, beUnioned = root2, root1
	} else {
		unionRoot, beUnioned = root2, root1
		s.Rank[root2]++
	}
	s.Parent[beUnioned] = unionRoot
	s.Size[unionRoot] += s.Size[beUnioned]
}

func (s *UnionFindSet) isConnected(e1, e2 int) bool {
	return s.Find(e1) == s.Find(e2)
}
