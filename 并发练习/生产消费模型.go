package main

import (
	"fmt"
	"strconv"
	"time"
)

type product struct {
	Name string
}

func main() {
	/*
		1.创建一个结构体，作为存放在channel中的数据类型
		2.开启两个协程，生产者与消费者
		3.创建两个channel，一个是用于生产者协程与消费者协程进行通信的，一个是用于消费者与主协程进行通信的
	*/
	chanShop := make(chan product, 10)
	//chanCount := make(chan int, 10)

	// 生产者
	go func() {
		for {
			car := product{Name:"product"+strconv.Itoa(time.Now().Second())}

			chanShop <- car
			
			fmt.Println("生产数据:", car)

			time.Sleep(1 * time.Second)
		}
	}()
	
	
	go func() {
		for {
			carmsg := <-chanShop
			fmt.Println("消费数据:", carmsg)
			
			//chanCount <- 1
			time.Sleep(1 * time.Second)
		}
	}()

	for value := range chanShop {
		fmt.Println("这是主协程的：",value)
	}

	time.Sleep(2 * time.Second)

}
