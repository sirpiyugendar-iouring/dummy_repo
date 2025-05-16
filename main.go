package main

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

var m map[string]int
var Cache atomic.Value
var keys []string

func main() {
	keys = strings.Split("1 2 3 4 5 6 7 8 9 10", " ")
	m = make(map[string]int)
	Cache.Store(m)
	for i := 0; i < 10; i++ {
		i := i
		go func(val int) {
			for {
				Map := newMap(val)
				Cache.Store(Map)
			}
		}(i)
	}
	go func() {
		for {
			time.Sleep(time.Microsecond)
			c := Cache.Load().(map[string]int)
			fmt.Println(c)
		}
	}()
	time.Sleep(time.Minute / 4)
}

func newMap(id int) map[string]int {
	res := map[string]int{}
	for _, key := range keys {
		res[key] = id
	}
	return res
}
