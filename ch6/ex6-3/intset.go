package main

import (
	"fmt"
	"bytes"
)

type IntSet struct {
	words []uint64
}

func main() {
	var x, y, a, b, c, d IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(100)
	y.Add(9)
	y.Add(42)
	y.Add(100)
	x.IntersectWith(&y)
	fmt.Println(x.String())

	a.Add(1)
	a.Add(100)
	a.Add(200)
	b.Add(100)
	a.DifferenceWith(&b)
	fmt.Println(a.String())

	c.Add(1)
	c.Add(100)
	d.Add(1)
	c.SymmetricDifference(&d)
	fmt.Println(c.String())
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint64(x%64)
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(t.words) > len(s.words) {
		t.words = t.words[:len(s.words)]
	}
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, _ := range s.words {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, _ := range s.words {
		if i < len(t.words) {
			s.words[i] &= ^(t.words[i])
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	i := s.Copy()
	s.UnionWith(t)
	i.IntersectWith(t)
	s.DifferenceWith(i)
}

func (s *IntSet) Len() int {
	length := 0
	for _, word := range s.words {
		for bit := 0; bit < 64; bit++ {
			if word&(1<<bit) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = make([]uint64, len(s.words))
	copy(t.words, s.words)
	return &t
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
