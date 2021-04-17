package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//做一个临界区：环形队列 [0,1,2,3,4]
var five [5]int      //共享数据
var cond *sync.Cond  //条件变更
var mutex sync.Mutex //锁
//控制生产者数量
var cond_five *sync.Cond  //条件变更
var mutex_five sync.Mutex //锁
var prod_count int = 0
var cotmer_index = 0

//生产者
func productor() {
	index := 0 //生产下标
	produm := 1000
	for {
		// 抢锁
		mutex.Lock()
		// 生产产品
		time.Sleep(time.Millisecond * 100)
		five[index] = produm
		fmt.Println("i am productor, num ====", produm)
		produm++
		prod_count++
		index = (index + 1) % 5 //环形队列
		cond.Signal()           //唤醒一个goroutine
		// 释放锁
		mutex.Unlock()
		if prod_count == 5 {
			mutex_five.Lock()
			cond_five.Wait()
			mutex_five.Unlock()
		}

	}
	wg.Done()
}

//消费者
func customer(num int) {

	for {
		//1 抢锁
		mutex.Lock()
		//2 wait
		cond.Wait()
		if prod_count > 0 { //防止出现极端情况
			//3 吃饼
			time.Sleep(time.Millisecond * 10)
			fmt.Printf("i am %d cutomer, num ==== %d\n", num, five[cotmer_index])
			cotmer_index = (cotmer_index + 1) % 5
			prod_count--
		}

		//4.释放锁
		mutex.Unlock()
		cond_five.Signal()

	}
	wg.Done()
}

func main() {
	cond = sync.NewCond(&mutex)           //构造条件变量
	cond_five = sync.NewCond(&mutex_five) //构造条件变量
	wg.Add(3)
	go productor()
	go customer(1)
	go customer(2)
	wg.Wait() //阻塞等待goroutine
}
