package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	lCnt := 0
	start := time.Now()
	c := make(chan int)
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		var spawns []int
		var s1024 string
		for i := 0; i < 4; i++ {
			m := []int{1, 0, 2, 4}
			go func() {
				time.Sleep(time.Duration(r.Intn(91)+10) * time.Millisecond)
				c <- m[r.Intn(4)]
			}()
		}
		for i := 0; i < 4; i++ {
			spawns = append(spawns, <-c)
		}
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
