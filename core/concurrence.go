package core

import "sync"

func DecreaseSemaphore(thread *int, cond *sync.Cond) {
	cond.L.Lock()
	for *thread == 0 {
		cond.Wait()
	}
	*thread--
	cond.L.Unlock()
}

func IncreaseSemaphore(thread *int, cond *sync.Cond) {
	cond.L.Lock()
	*thread++
	cond.Broadcast()
	cond.L.Unlock()
}
