package unionfind

type QuickFindUF struct {
	id    []int
	count int
}

func NewQuickFindUF(n int) UnionFind {
	uf := &QuickFindUF{
		id:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return uf
}

func (uf *QuickFindUF) Find(p int) int {
	return uf.id[p]
}

func (uf *QuickFindUF) Union(p int, q int) {
	pID := uf.Find(p)
	qID := uf.Find(q)
	if pID == qID {
		return
	}
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	uf.count--
}

func (uf *QuickFindUF) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *QuickFindUF) Count() int {
	return uf.count
}
