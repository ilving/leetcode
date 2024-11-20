package main

import (
	"math"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"
)

func twoSum(nums []int, target int) []int {
	var v, i, num int
	var ok bool

	hm := map[int]int{}
	for i, num = range nums {
		if v, ok = hm[num]; ok {
			return []int{v, i}
		}
		hm[target-num] = i
	}

	return []int{0, 0}
}

const amount = 100

func main() {
	// sl := []int{2, 7, 11, 15}
	// target := 9

	sl := []int{}
	rand.Seed(time.Now().UnixMicro())
	for len(sl) < 10000 {
		sl = append(sl, math.MaxInt16-rand.Intn(math.MaxUint16))
	}
	target := sl[len(sl)-1] + sl[len(sl)-2]

	debug.SetGCPercent(-1)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	t := time.Now().UTC()
	allocs := m.TotalAlloc
	for i := 0; i < amount; i++ {
		twoSum(sl, target)
	}
	runtime.ReadMemStats(&m)
	allocs = m.TotalAlloc - allocs

	println(time.Now().UTC().Sub(t).Microseconds()/amount, " mks | ", (allocs/amount)/1024.0, " kb/call")

}
