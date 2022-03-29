// 此脚本用于从正整数 1~14 中随机输出 3~6 个不重复的数

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	var num int
	num = rand.Intn(4) + 3

	var a = [6]int{}

	m := 0
	for {
		x := rand.Intn(14) + 1

		var b bool
		for _, y := range a {
			if y == x {
				b = true
				break
			}
		}
		if b == true {
			continue
		}

		a[m] = x
		m++
		if m == num {
			break
		}
	}
	//fmt.Println(a)
	for _, y := range a {
		if y != 0 {
			fmt.Println(y)
		}
	}
	time.Sleep(time.Second * 20)
}
