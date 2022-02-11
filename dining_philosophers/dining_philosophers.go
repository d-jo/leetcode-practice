package dining_philosophers

import (
	"fmt"
	"sync"
	"time"
)

/*
	"Imagine you have five philosophers sitting around a roundtable. The philosophers do only two kinds of activities. One: they contemplate, and two: they eat.

However, they have only five forks between themselves to eat their food with. Each philosopher requires both the fork to his left and the fork to his right to eat his food.

Design a solution where each philosopher gets a chance to eat his food without causing a deadlock."
*/

/*
	should use semaphores to restrict forks

	forks: [0,1,2,3,4]
	phils: forks
		0: 0, 4
		1: 1, 0
		2: 2, 1
		3: 3, 2
		4: 4, 3

*/

func GetForkIndiciesForPhil(phil_ind int) (int, int) {
	// for these we add 5 because we may get an input of -1
	// and -1 % 5 == -1, but we want it to equal 4 as they are sitting
	// on a circular table. so by adding 5 (the number of forks) we
	// change its cycle
	return (phil_ind + 5) % 5, (phil_ind - 1 + 5) % 5
}

func Eat(phil_ind int, forks []chan int, done chan interface{}, wg *sync.WaitGroup, meals int) {
	left, right := GetForkIndiciesForPhil(phil_ind)
	currMeals := 0

	if (phil_ind+2)%2 == 0 {
		// check left first
		for {
			select {
			case <-done:
				return
			case forks[left] <- 1:
				forks[right] <- 1
				// eat
				fmt.Printf("%d is eating with forks %d, %d\n", phil_ind, left, right)
				time.Sleep(time.Second)
				fmt.Printf("%d is done eating with forks %d, %d\n", phil_ind, left, right)
				// return forks
				<-forks[right]
				<-forks[left]
			}
			currMeals += 1

			if currMeals >= meals {
				break
			}
		}
	} else {
		// check right first
		for {
			select {
			case <-done:
				return
			case forks[right] <- 1:
				forks[left] <- 1
				// eat
				fmt.Printf("%d is eating with forks %d, %d\n", phil_ind, left, right)
				time.Sleep(time.Second)
				fmt.Printf("%d is done eating with forks %d, %d\n", phil_ind, left, right)
				// return forks
				<-forks[left]
				<-forks[right]
			}

			currMeals += 1

			if currMeals >= meals {
				break
			}

		}
	}

	wg.Done()
}

func Start() {
	// general idea:
	// we have a semaphore for each fork
	// a go routine for each philosopher
	// when a philosopher wants to eat, it will first
	// acquire the semaphore for the forks to its left and right
	// if it cant, wait
	// eating: aquire forks, sleep, then return forks

	// I dont know if this solution can guarantee that it never deadlocks
	// or incurs starvation. I think it might be safe, but cant think of a way
	// to prove it. Maybe change the sleep time and run many benchmarks

	forks := make([]chan int, 5)
	for i := 0; i < 5; i++ {
		forks[i] = make(chan int, 1)
	}

	var wg sync.WaitGroup
	var done chan interface{}
	meals := 1

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go Eat(i, forks, done, &wg, meals)
	}

	fmt.Println("waiting")
	wg.Wait()
}
