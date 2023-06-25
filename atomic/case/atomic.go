package _case

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCase() {
	var count int64 = 5
	atomic.StoreInt64(&count, 10)
	fmt.Println("获取变量的值：", atomic.LoadInt64(&count))
	//count+=10
	fmt.Println("在原有的基础上累加：", atomic.AddInt64(&count, 10))
	//交换变量的值，并返回原有的值
	fmt.Println("交换变量的值，并返回原有的值：", atomic.SwapInt64(&count, 30))
	fmt.Println("获取变量的值：", atomic.LoadInt64(&count))
	//比较并替换原有的值，并返回是否替换成功,会提前对比旧值确认是否发生变化，如果相等就不替换，不相等则替换
	fmt.Println("比较并替换原有的值，并返回是否替换成功:", atomic.CompareAndSwapInt64(&count, 30, 40))
	fmt.Println("获取变量的值：", atomic.LoadInt64(&count))
}

type atomicCounter struct {
	count int64
}

func (c *atomicCounter) Inc() {
	atomic.AddInt64(&c.count, 1)
}

func (c *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&c.count)
}

func AtomicCase1() {
	var count = atomicCounter{}
	//locker := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
}

func AtomicCase2() {
	list := []string{"A", "B", "C", "D"}
	//定义一个原子值
	var atomicMap atomic.Value

	//定义一个集合
	mp := map[string]int{}
	//将集合存储到原子值
	atomicMap.Store(&mp)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

		atomicLabel:

			m := atomicMap.Load().(*map[string]int)

			m1 := map[string]int{}
			for k, v := range *m {
				m1[k] = v
			}

			for _, item := range list {
				_, ok := m1[item]
				if !ok {
					m1[item] = 0
				}
				m1[item] += 1
			}
			swap := atomicMap.CompareAndSwap(m, &m1)
			if !swap {
				//重新执行替换逻辑
				goto atomicLabel
			}
		}()
	}
	wg.Wait()
	fmt.Println(atomicMap.Load())

}
