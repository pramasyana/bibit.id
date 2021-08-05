package main

import "sync"

func main() {
	service := InitServer()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.Serve()
	}()

	wg.Wait()
}
