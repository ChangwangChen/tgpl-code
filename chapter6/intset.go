package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func process(x int) (int, uint) {
	return x / 64, uint(x % 64)
}

func (s *IntSet) Has(x int) bool {
	word, bit := process(x)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := process(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// s 合并 t 结果保存在 s 中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
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

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	word, bit := process(x)
	s.words[word] |= 1 << bit
	s.words[word] ^= 1 << bit
}

func main() {
	var x, y IntSet

	x.Add(1)
	x.Add(23)
	x.Add(145)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	fmt.Println((&x).Len())

	y.Add(23)
	y.Add(144)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(23))
	fmt.Println(y.Has(34))

	fmt.Println(x)
	fmt.Println(y)

	x.Remove(1)
	fmt.Println(x.String())
}
