package main

import (
	"flag"
	"github.com/iamhujiebin/gotools/bitmap"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var test = flag.String("test", "bitmap", "write your test func")

type testFunc func()

type Test struct {
	BitMap testFunc
}

func main() {
	flag.Parse()
	t := Test{
		BitMap: BitMapTestFunc,
	}
	switch *test {
	case "bitmap":
		t.BitMap()
	}
}

func BitMapTestFunc() {
	Wrap(func() {
		b := bitmap.New()
		b.Add(2)
		b.Add(1)
		var i uint64
		for i = 1235000000; i < 1236000000; i++ {
			b.Add(i)
		}
		println(b.Nums())
	})
}

func Wrap(f func()) {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	f()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
