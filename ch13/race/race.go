package race

import "sync"

func GetCounter() int {
	var counter int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			tmp := 0
			for i := 0; i < 1000; i++ {
				tmp++
			}
			mu.Lock()
			counter += tmp
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}
