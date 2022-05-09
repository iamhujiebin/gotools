package bitmap

import (
	"bytes"
	"fmt"
)

type BitMap struct {
	words  []uint64
	length int
}

func New() *BitMap {
	return &BitMap{}
}

func (b *BitMap) Has(num int) bool {
	word, bit := num/64, uint(num%64) // 64是uint64的64位
	return word < len(b.words) && (b.words[word]&(1<<bit)) != 0
}

func (b *BitMap) Add(num int) {
	word, bit := num/64, uint(num%64) // 64是uint64的64位
	for word >= len(b.words) {
		b.words = append(b.words, 0)
	}
	if b.words[word]&(1<<bit) == 0 {
		b.words[word] |= 1 << bit
		b.length++
	}
}

func (b *BitMap) Remove(num int) {
	word, bit := num/64, uint(num%64)
	if word >= len(b.words) {
		return
	}
	if b.words[word]&(1<<bit) == 1 {
		b.words[word] &= ^(uint64(1) << bit) // uint64(1) 注意要这样,直接用1是有符号整数,去反不行
		b.length--
	}
}

func (b *BitMap) Len() int {
	return b.length
}

func (b *BitMap) Nums() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range b.words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if v&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*uint(i)+j)
			}
		}
	}
	buf.WriteByte('}')
	fmt.Fprintf(&buf, "\nLength: %d,bitWords:%d", b.length, len(b.words))
	return buf.String()
}

func (b *BitMap) Bits() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i := len(b.words) - 1; i >= 0; i-- {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%b", b.words[i])
	}
	buf.WriteByte('}')
	fmt.Fprintf(&buf, "\nLength: %d,bitWords:%d", b.length, len(b.words))
	return buf.String()
}
