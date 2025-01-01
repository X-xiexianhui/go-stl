package unionfind

type UnionFind interface {
	Find(p int) int
	Union(p, q int)
	Connected(p, q int) bool
	Count() int
}
