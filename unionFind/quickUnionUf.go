package unionfind

type QuickUnionUF struct {
	id    []int
	count int
}

func NewQuickUnionUF(n int) UnionFind {
	uf := &QuickUnionUF{
		id:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return uf
}
func (uf *QuickUnionUF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *QuickUnionUF) Union(p, q int) {
	i := uf.Find(p)
	j := uf.Find(q)
	if i == j {
		return
	}
	uf.id[i] = j
	uf.count--
}
func (uf *QuickUnionUF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}
func (uf *QuickUnionUF) Count() int {
	return uf.count
}
