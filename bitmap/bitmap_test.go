package bitmap

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestBitOp(t *testing.T) {
	bit := 10
	t.Logf("left:%b\n ", uint64(1)<<bit)
	t.Logf("reverse:%b\n ", ^(uint64(1) << bit))
}

func TestNew(t *testing.T) {
	b := New()
	b.Add(uint64(1235))
	b.Add(123)
	t.Logf("nums:%s", b.Nums())
	t.Logf("bits:%s", b.Bits())
	t.Logf("mem:%v", unsafe.Sizeof(b))
}

func TestRemove(t *testing.T) {
	b := New()
	b.Add(1111)
	b.Add(9999)
	t.Logf(b.Nums())
	b.Remove(9999)
	t.Logf(b.Nums())
}

func TestHas(t *testing.T) {
	b := New()
	b.Add(1111)
	b.Add(9999)
	t.Logf("has 9999:%v", b.Has(9999))
	t.Logf(b.Nums())
	b.Remove(9999)
	t.Logf(b.Nums())
	t.Logf("has 9999:%v", b.Has(9999))
}

func TestBig(t *testing.T) {
	b := New()
	b.Add(uint64(4000000000))
	t.Logf("nums:%s", b.Nums())
	bytes := 8 * len(b.words)
	t.Logf("size:%v bytes, %v kb , %v mb", bytes, float64(bytes)/1024, float64(bytes)/1024/1024)
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	bytes = int(m.TotalAlloc)
	t.Logf("mem:%v bytes, %v kb , %v mb", bytes, float64(bytes)/1024, float64(bytes)/1024/1024)
}
