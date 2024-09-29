package path

import (
	"fmt"
	"sync"
)

type Runtime interface {
	LifecycleManager[string, ExistingPath]
}

type DefaultRuntime struct {
	SynchronousLifecycleManager[string, ExistingPath]
}

type GoRoutinesRuntime struct {
	GoRoutinesLifecycleManager[string, ExistingPath]
	quit chan int
}

type LifecycleManager[T, R any] interface {
	AllowedToRun() bool
	RunAll(func(int, T) (R, error), []T) ([]R, error)
}

type SynchronousLifecycleManager[T, R any] struct {
	LifecycleManager[T, R]
}

func (m *SynchronousLifecycleManager[T, R]) AllowedToRun() bool {
	return true
}

func (m *SynchronousLifecycleManager[T, R]) RunAll(f func(int, T) (R, error), elems []T) ([]R, error) {
	results := make([]R, len(elems))
	for i, elem := range elems {
		if result, err := f(i, elem); err == nil {
			results[i] = result
		} else {
			return nil, err
		}
	}
	return results, nil
}

type GoRoutinesLifecycleManager[T, R any] struct {
	quit chan int
}

func (m *GoRoutinesLifecycleManager[T, R]) AllowedToRun() bool {
	fmt.Println("GoRoutinesLifecycleManager.AllowedToRun()")
	select {
	case c := <-m.quit:
		fmt.Println("GoRoutinesLifecycleManager.AllowedToRun(): quit", c)
		return false
	default:
		return true
	}
}

func (m *GoRoutinesLifecycleManager[T, R]) RunAll(f func(int, T) (R, error), elems []T) ([]R, error) {
	ch := make(chan R, len(elems))
	defer close(ch)
	err_ch := make(chan error, len(elems))
	defer close(err_ch)
	var wg sync.WaitGroup
	wg.Add(len(elems))
	for i, elem := range elems {
		go func(i int, elem T) {
			defer wg.Done()
			if result, err := f(i, elem); err == nil {
				ch <- result
				err_ch <- nil
			} else {
				ch <- result
				err_ch <- err
			}
		}(i, elem)
	}
	wg.Wait()

	results := make([]R, len(elems))
	for i := 0; i < len(elems); i++ {
		err := <-err_ch
		if err != nil {
			return nil, err
		}
		results[i] = <-ch
	}
	return results, nil
}
