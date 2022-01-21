package set

import (
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	ct := NewCommonSet()
	ct.Insert(1)
	ct.Insert("xx")
	ct.Insert(&f{id: 11})
	for k := range ct {
		fmt.Println(k)
	}
}

type f struct {
	name string
	id   int
}
