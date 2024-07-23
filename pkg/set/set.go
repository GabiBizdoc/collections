package set

import (
	"maps"
)

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(value string) {
	s[value] = struct{}{}
}

func (s Set) AddMany(value ...string) {
	for i := range value {
		s.Add(value[i])
	}
}

func (s Set) Remove(value string) {
	delete(s, value)
}

func (s Set) Has(value string) bool {
	_, ok := s[value]
	return ok
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Empty() bool {
	return len(s) == 0
}

func (s Set) Copy() Set {
	return maps.Clone(s)
}

func (s Set) ToSlice() []string {
	s2 := make([]string, 0, s.Len())
	for k := range s {
		s2 = append(s2, k)
	}
	return s2
}
