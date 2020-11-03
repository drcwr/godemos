package main

// 闭包不加接收的参数，就会默认使用外部变量的地址拷贝
import (
	"fmt"
	"sync"
)

func main() {
	input()
	noinput()
}

func input() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// time.Sleep(time.Duration(1) * time.Microsecond)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func noinput() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// time.Sleep(time.Duration(1) * time.Microsecond)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func mutextest() {
	var m *sync.Mutex
	m.Lock()
	m.Unlock()
}
