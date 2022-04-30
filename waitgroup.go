package waitgroup

import "sync"

// WaitGroup implements a simple goroutine pool.
type WaitGroup struct {
	size      int
	pool      chan byte
	waitGroup sync.WaitGroup
}

// New creates a waitgroup with a specific size (the maxium number of
// goroutines to run at the same time). If you use -1 as the size, all items
// will run concurrentlu (just like a normal sync.WaitGroup).
func New(size int) *WaitGroup {
	wg := &WaitGroup{
		size: size,
	}
	if size > 0 {
		wg.pool = make(chan byte, size)
	}

	return wg
}

// Add adds the function to the waitgroup.
func (wg *WaitGroup) Add(fns ...func()) {
	for _, f := range fns {
		fn := f
		wg.blockAdd()
		go func() {
			defer wg.Done()

			fn()
		}()
	}
}

func (wg *WaitGroup) blockAdd() {
	if wg.size > 0 {
		wg.pool <- 1
	}

	wg.waitGroup.Add(1)
}

// Done pops one out of the pool.
func (wg *WaitGroup) Done() {
	if wg.size > 0 {
		<-wg.pool
	}

	wg.waitGroup.Done()
}

// Wait waits the pool empty.
func (wg *WaitGroup) Wait() {
	wg.waitGroup.Wait()
}

// PendingCount returns the number of pending tasks.
func (wg *WaitGroup) PendingCount() int64 {
	return int64(len(wg.pool))
}
