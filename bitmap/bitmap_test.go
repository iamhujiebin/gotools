package bitmap

import (
	"testing"
)

func TestBitOp(t *testing.T) {
	bit := 10
	t.Logf("left:%b\n ", uint64(1)<<bit)
	t.Logf("reverse:%b\n ", ^(uint64(1) << bit))
}

func TestNew(t *testing.T) {
	b := New()
	b.Add(1235)
	b.Add(123)
	t.Logf("nums:%s", b.Nums())
	t.Logf("bits:%s", b.Bits())
}
