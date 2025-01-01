package unionfind

type WeightedQuickUnionUF struct {
	id    []int
	sz    []int
	count int
}

func NewUnionFind(n int) UnionFind {
	uf := &WeightedQuickUnionUF{
		id:    make([]int, n),
		sz:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return uf
}
func (uf *WeightedQuickUnionUF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *WeightedQuickUnionUF) Union(p, q int) {
	i := uf.Find(p)
	j := uf.Find(q)
	if i == j {
		return
	}
	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] += uf.sz[i]
	} else {
		uf.id[j] = i
		uf.sz[i] += uf.sz[j]
	}
	uf.count--
}

func (uf *WeightedQuickUnionUF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *WeightedQuickUnionUF) Count() int {
	return uf.count
}
