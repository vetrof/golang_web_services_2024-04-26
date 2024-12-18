package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum__ = 7
	goroutinesNum__ = 5
)

func startWorker__(in int, waiter *sync.WaitGroup) {
	defer waiter.Done() // wait_2.go уменьшаем счетчик на 1
	for j := 0; j < iterationsNum__; j++ {
		fmt.Printf(formatWork__(in, j))
		time.Sleep(time.Millisecond) // попробуйте убрать этот sleep
	}
}

func main() {
	wg := &sync.WaitGroup{} // wait_2.go инициализируем группу
	for i := 0; i < goroutinesNum__; i++ {
		// wg.Add надо вызывать в той горутине, которая порождает воркеров
		// иначе другая горутина может не успеть запуститься и выполнится Wait
		wg.Add(1) // wait_2.go добавляем
		go startWorker__(i, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait() // wait_2.go ожидаем, пока waiter.Done() не приведёт счетчик к 0

	fmt.Println(11111)

}

func formatWork__(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum__-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}
