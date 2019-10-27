package main

/*
	100条协程并发请求1-1000的平方根
	最大的并发数为5
	管道实现
*/
//func main() {
//	chanControl := make(chan string, 5)
//	for i := 0; i < 100; i++ {
//		go GetSqrt("我是第"+strconv.Itoa(i)+"个协程",rand.Intn(1000), chanControl)
//	}
//
//	time.Sleep(30 * time.Second)
//}

//func GetSqrt(name string, n int, c chan string)  {
//	// 这里的channel是一个很巧妙的用法，借助channel具有阻塞的特点来控制并发数
//	c <-name
//	sqrt := math.Sqrt(float64(n))
//	fmt.Println("平方根：",sqrt)
//	time.Sleep(time.Second)
//	<-c
//}
