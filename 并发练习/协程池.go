package main

import (
	"fmt"
	"sync"
	"time"
)

/*
设计思路：首先需要先定义好一个任务，创建任务对象之后可以将这个任务推送到协程池对外的入口channel中，
		协程池有三个属性，一个是入口channel，一个是对内的channel，还有一个就是协程数，这个是用于控制有多少个协程可以进行并发操作
		定义协程池的一个方法worker，这个方法的作用是操作操作协程池中对内channel的中的数据
		创建一个Run方法，这个函数用于启动协程池，主要的操作是创建worker协程，然后不断的从入口channel中获取数据推送到对内channel
*/

// 一个任务就是操作一些事情
type Stask struct {
	todo func() error
}

// 执行Stask对象中的属性
func (s *Stask) Execute()  {
	s.todo()
}

// 用于创建结构体对象,对象中的方法需要根据实际需求动态传进去
func NewStask(f func() error)  *Stask{

	return &Stask{todo: f}
}

type Pool struct {
	// 对外的入口 chan
	EntryChan chan *Stask
	// 连接内外管道
	JobsChan chan *Stask
	// 协程数
	worker_count int

}

// 实例化协程池
func InitPool(num int) *Pool {
	return &Pool{
		EntryChan:    make(chan *Stask),
		JobsChan:     make(chan *Stask),
		worker_count: num,
	}
}

// 定义协程池的worker
func (p *Pool)Worker(index int)  {
	for value := range p.JobsChan {
		fmt.Printf("我是第%d个协程，", index)
		fmt.Println("操作的数据是:", value)
	}

}

// 启动协程池，
func (p *Pool) Run()  {
	for i := 0; i < p.worker_count; i++ {
		go p.Worker(i)
	}

	for value := range p.EntryChan {
		fmt.Println("执行2", value)
		p.JobsChan <- value
	}


}


func main() {
	var wg sync.WaitGroup
	pool := InitPool(3)

	wg.Add(1)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for  {
			<- ticker.C
			stask := NewStask(func() error {
				fmt.Println("创建线程：", time.Now())

				return nil
			})
			fmt.Println("执行1")
			pool.EntryChan <- stask

		}
		wg.Done()
	}()

	pool.Run()

	wg.Wait()
	//pool.Run()
}
