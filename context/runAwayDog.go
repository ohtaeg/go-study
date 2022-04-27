package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
강아지가 사라졌다. 엄마, 아빠, 누나가 동네에 흩어져서 강아지를 찾기 시작한다.
엄마가 강아지를 찾았다면, 엄마는 아빠에게 연락을 해서 강아지 찾기 미션을 cancel 하고 집으로 오라고 연락해야 한다.
*/
type Person struct {
	Name  string
	Delay time.Duration
}

func (p *Person) findDog(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(p.Name + " 찾기 시작!")
	select {
	case <-time.After(p.Delay):
		fmt.Println(p.Name + " 가 강아지를 찾았다!")
	case <-ctx.Done():
		fmt.Println(p.Name + " 바로 돌아갈께")
	}
}

func init() {
	family = make([]Person, 3)
	family[0] = Person{
		Name:  "mom",
		Delay: 60 * time.Second,
	}
	family[1] = Person{
		Name:  "father",
		Delay: 60 * time.Second,
	}
	family[2] = Person{
		Name:  "me",
		Delay: 3 * time.Second,
	}
}

var family []Person

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	for _, who := range family {
		wg.Add(1)
		go func(person Person) {
			person.findDog(ctx, &wg)
			defer cancel()
		}(who)
	}

	wg.Wait()
}
