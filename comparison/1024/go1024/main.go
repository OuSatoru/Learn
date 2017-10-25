package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	lCnt := 0
	start := time.Now()
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		var spawns []int
		var s1024 string
		for i := 0; i < 4; i++ {
			wg.Add(1)
			m := []int{1, 0, 2, 4}
			go func() {
				time.Sleep(time.Duration(r.Intn(91)+10) * time.Millisecond)
				mutex.Lock()
				spawns = append(spawns, m[r.Intn(4)])
				mutex.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
		for _, n := range spawns {
			s1024 += strconv.Itoa(n)
		}
		fmt.Println(s1024)
		lCnt++
		if s1024 == "1024" {
			fmt.Printf("经历了 %d 次随机生成，花费 %v 秒。10.24 程序员节快乐！(*^▽^*)", lCnt, time.Since(start))
			break
		}
	}
}
