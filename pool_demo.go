package main

import (
	"sync"
	"strconv"
	"fmt"
	"time"
)

type PoolItem struct {
	Value int
}

func (i *PoolItem) String() string {
	return strconv.Itoa(i.Value)
}

func main() {

	pool := &sync.Pool{}

	go getterFromPool(pool)
	go adderToPool(pool)


	time.Sleep(5 * time.Second)

	fmt.Println("wait")
}


func adderToPool(p *sync.Pool) {
	for i := 0; i < 30; i++ {
		p.Put(&PoolItem{i})
		fmt.Println("put ", i)
	}
}

func getterFromPool(p *sync.Pool) {
	for i := 0; i < 10000; i++ {
		item := p.Get()
		if item != nil {
			fmt.Println("get", item, i)
		}
	}
}