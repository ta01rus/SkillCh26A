package container

type IntRing struct {
	data    []*int
	pozPut  int
	pozRead int
}

// колцевой буфер
func NewIntRing(size int) *IntRing {
	return &IntRing{
		pozPut: -1,
		data:   make([]*int, size),
	}
}

// вставка
func (i *IntRing) Put(n int) {
	i.pozPut++
	if i.pozPut == len(i.data) {
		i.pozPut = 0
	}
	if i.data[i.pozPut] != nil {
		i.pozRead = i.pozPut + 1
	}
	i.data[i.pozPut] = &n
}

// получить след
func (i *IntRing) Get() *int {
	var now *int
	if i.pozRead == len(i.data) {
		i.pozRead = 0
	}
	if i.data[i.pozRead] != nil {
		now = i.data[i.pozRead]
		i.data[i.pozRead] = nil
		i.pozRead++
	}

	return now
}
