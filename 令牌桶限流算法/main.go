package main

import (
	"fmt"
	"time"
)

var capacity = 100      // 容量
var tokenSpeed = 1      // 放入令牌的速度 多少秒钟生产 一个
var capacityCh chan int // 用一个通道来实现令牌桶的数量管理

func main() {
	// 用一个管道来实现 令牌桶的容量
	capacityCh = make(chan int, capacity)
	go genToken(capacityCh)
	var i = 0
	for {
		tokenBucketRequest(i, capacityCh)
		time.Sleep(time.Second * 3)
		i++
	}
}

func genToken(capacityCh chan int) {
	for {
		select {
		case capacityCh <- 1:
		default:
			fmt.Printf("通道已满\n")
		}
		time.Sleep(time.Second * time.Duration(tokenSpeed))

	}
}

//**生成令牌的速度是恒定的，而请求去拿令牌是没有速度限制的。这意味，面对瞬时大流量，该算法可以在短时间内请求拿到大量令牌，
//而且拿令牌的过程并不是消耗很大的事情。**
func tokenBucketRequest(request int, capacityCh chan int) bool {
	if len(capacityCh) == 0 {
		fmt.Printf("当前请求%d没有足够的令牌\n", request)
		return false
	}
	_, ok := <-capacityCh
	if !ok {
		//  通道已经被关闭
		fmt.Println("通道已经被关闭")
		return false
	}
	fmt.Printf("当前请求%d,当前令牌数量%d\n", request, len(capacityCh))
	return true
}
