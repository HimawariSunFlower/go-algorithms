package set

var (
	void struct{}
)

type IntSet map[int]struct{}

func (s IntSet) Has(key int) bool {
	_, ok := s[key]
	return ok
}

func (s IntSet) HasAll(ll []int) bool {
	for _, v := range ll {
		_, r := s[v]
		if !r {
			return r
		}
	}
	return true
}

func (s IntSet) Insert(key int) bool {
	_, r := s[key]
	if !r {
		s[key] = void
	}
	return !r
}

func (s IntSet) Length() int {
	i := 0
	for range s {
		i++
	}
	return i
}

func (s IntSet) Remove(v int) {
	delete(s, v)
}

func (s IntSet) Clear() {
	for k, _ := range s {
		delete(s, k)
	}
}

func (s IntSet) FromSlice(ll []int) {
	for _, v := range ll {
		s.Insert(v)
	}
}

func (s IntSet) ToSlice() []int {
	ll := make([]int, 0, s.Length())
	for k := range s {
		ll = append(ll, k)
	}
	return ll
}

func (s IntSet) Equal(tg IntSet) bool {
	if s.Length() != tg.Length() {
		return false
	}
	for k := range s {
		if !tg.Has(k) {
			return false
		}
	}
	return true
}

//并集
func (s IntSet) Union(tg IntSet) []int {
	r := make([]int, 0)
	for k := range s {
		r = append(r, k)
	}
	for k := range tg {
		if !s.Has(k) {
			r = append(r, k)
		}
	}
	return r
}

//差集
func (s IntSet) Difference(tg IntSet) []int {
	r := make([]int, 0)
	for k := range s {
		if !tg.Has(k) {
			r = append(r, k)
		}
	}
	return r
}

type CommonSet map[interface{}]struct{}

func NewCommonSet() CommonSet {
	return make(CommonSet)
}

func (s CommonSet) Has(key interface{}) bool {
	_, ok := s[key]
	return ok
}

func (s CommonSet) Insert(key interface{}) {
	s[key] = struct{}{}
}

func (s CommonSet) Length() int {
	i := 0
	for range s {
		i++
	}
	return i
}
