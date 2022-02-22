package adt

// UF unite & find
type UF struct {
	id       []int
	nodeSize []int
}

func NewUF(size int) *UF {
	id := make([]int, size)
	nodeSize := make([]int, size)
	for i := 0; i < size; i++ {
		id[i] = i
		nodeSize[i] = 1
	}
	return &UF{id: id, nodeSize: nodeSize}
}

func (uf *UF) find(v int) int {
	// find root
	for v != uf.id[v] {
		uf.id[v] = uf.id[uf.id[v]] // halving the path
		v = uf.id[v]
	}
	return v
}

func (uf *UF) Find(v1 int, v2 int) bool {
	//comparison of roots
	return uf.find(v1) == uf.find(v2)
}

func (uf *UF) Unite(v1 int, v2 int) {
	i := uf.find(v1)
	j := uf.find(v2)
	if i == j {
		return
	}
	if uf.nodeSize[j] < uf.nodeSize[i] {
		uf.id[j] = i
		uf.nodeSize[i] += uf.nodeSize[j]
	} else {
		uf.id[i] = j
		uf.nodeSize[j] += uf.nodeSize[i]
	}
}
