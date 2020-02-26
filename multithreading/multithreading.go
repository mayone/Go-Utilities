package multithreading

import "github.com/mayone/OGGo/Go-Utilities/multithreading/threadpool"

type task struct {
	a []int
	b []int
	c []int
}

func (t *task) Run(tInfo threadpool.ThreadInfo) {
	offset := tInfo.id
	step := tInfo.numThreads
	for i := offset; i < len(t.c); i += step {
		t.c[i] = t.a[i] + t.b[i]
	}
}

func addArraysParallel(a []int, b []int) []int {

	if len(a) != len(b) {
		return nil
	}
	c := make([]int, len(a))
	tp := threadpool.NewThreadPool()
	task := &task{a: a, b: b, c: c}
	tp.RunWorkers(task)

	return c
}

func addArraysSequential(a []int, b []int) []int {
	if len(a) != len(b) {
		return nil
	}
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}
