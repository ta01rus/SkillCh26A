package container

type IntRing struct {
	data             []*int
	start, end, size int
}

// колцевой буфер
func NewIntRing(size int) *IntRing {
	return &IntRing{}
}

// вставка
func (i *IntRing) Push(n int) {
	switch {
	case i.start == i.size:
		i.start = 0
		fallthrough
	case i.end == i.size:
		i.end = 0
		fallthrough
	case i.start == i.end:
		i.start++
	}
	i.data[i.end] = &n
	i.end++
}

// получить след
func (i *IntRing) Get(n *int) *int {
	ret := i.data[i.start]
	i.data[i.start] = nil
	i.start++
	return ret

}
