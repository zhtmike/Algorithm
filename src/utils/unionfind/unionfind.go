package unionfind

// UnionFind is a data structure stores each meaber's group ID and group size
type UnionFind struct {
	ids     []int
	sizes   map[int]int
	members map[int][]int
}

// Find method finds group that id belongs to
func (u *UnionFind) Find(id int) int {
	return u.ids[id]
}

// Merge merges two groups where ci and cj belongs to
func (u *UnionFind) Merge(ci int, cj int) {
	// find the groups of ci and cj
	gpi, gpj := u.ids[ci], u.ids[cj]
	// get two groups size
	si, sj := u.sizes[gpi], u.sizes[gpj]
	// merge the smaller group to the larger one
	if si > sj {
		for _, idj := range u.members[gpj] {
			u.ids[idj] = gpi
			u.members[gpi] = append(u.members[gpi], idj)
		}
		u.sizes[gpi] += sj
		delete(u.members, gpj)
		delete(u.sizes, gpj)
	} else {
		for _, idi := range u.members[gpi] {
			u.ids[idi] = gpj
			u.members[gpj] = append(u.members[gpj], idi)
		}
		u.sizes[gpj] += si
		delete(u.members, gpi)
		delete(u.sizes, gpi)
	}
}

// CreateUnionFind creats unionfind instance
func CreateUnionFind(n int) UnionFind {
	var u UnionFind
	u.ids = make([]int, n)
	u.sizes = make(map[int]int)
	u.members = make(map[int][]int)
	for i := 0; i < n; i++ {
		u.ids[i] = i
		u.sizes[i] = 1
		u.members[i] = []int{i}
	}
	return u
}
