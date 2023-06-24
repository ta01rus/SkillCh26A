package container

import (
	"log"
	"testing"
)

// func TestMain(m *testing.M) {
// 	rand.Seed(time.Now().UnixNano())
// 	m.Run()
// }

func TestRing(t *testing.T) {

	r := NewIntRing(100)

	for i := 0; i < 1000; i++ {
		// i := rand.Intn(100) + 1
		r.Put(i)
	}

	log.Println(r)

}
