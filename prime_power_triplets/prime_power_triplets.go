package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

func pow(base, exp int32) int32 {
	amt := int32(1)
	for i := int32(0); i < exp; i++ {
		amt *= base
	}
	return amt
}

/*
	checks all index combos between start and end for each list

	start - the index to start at for the lists
	end - the index to end at for the lists
*/
func search(starta, enda, startb, endb, startc, endc int, A, B, C *[]int32, done <-chan interface{}) chan int {

	results := make(chan int)

	go func() {
		defer close(results)

		// for a,b,c
		for a := starta; a <= enda; a++ {
			for b := startb; b <= endb; b++ {
				for c := startc; c <= endc; c++ {
					// check done, if it's closed, return
					select {
					case <-done:
						return
					default:
						// do the work!
						//fmt.Printf("%d, %d, %d\n", (*A)[a], (*B)[b], (*C)[c])
						//fmt.Printf("%d, %d, %d\n", pow((*A)[a], 2), pow((*B)[b], 3), pow((*C)[c], 4))
						prod := pow((*A)[a], 2) + pow((*B)[b], 3) + pow((*C)[c], 4)
						// work done, if its less than 50000000 send the result on the chan
						// or if done is closed, die
						if prod < 50000000 {
							select {
							case results <- int(prod):
							case <-done:
								return
							}
						}
					}
				}
			}
		}

	}()

	return results
}

func primesUpTo(max int32) []int32 {
	results := make([]int32, 0)
	for i := int32(2); i <= max; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			results = append(results, int32(i))
		}
	}
	return results
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//okay, here is the prob:
	/*
		Problem 87

		The smallest number expressible as the sum of a prime square, prime cube, and prime fourth power is 28. In fact, there are exactly four numbers below fifty that can be expressed in such a way:

		28 = 22 + 23 + 24
		33 = 32 + 23 + 24
		49 = 52 + 23 + 24
		47 = 22 + 33 + 24

		How many numbers below fifty million can be expressed as the sum of a prime square, prime cube, and prime fourth power?
	*/

	// notes:
	// we only need permutations of primes that are less than the fourth root of 50mil
	// primes less than 7072 for position a
	// primes less than 369 for position b
	// primes less than 84 for position c

	// step 1: get primes up to the fourth root of 50mil
	// step 2: check all the combos of primes a,b,c s.t. a^2+b^3+c^4 < 50mil

	aPrimes := primesUpTo(7074) // 908
	bPrimes := primesUpTo(371)  // 73
	cPrimes := primesUpTo(89)   // 23
	fmt.Printf("aPrimes: %d, bPrimes: %d, cPrimes: %d\n", len(aPrimes), len(bPrimes), len(cPrimes))
	fmt.Printf("c: %+v\n", cPrimes)

	// now we check each combo of primes from a in A,b in B,c in C
	// doing it concurrently using go channels.

	//resultChan := make(chan int, 4)
	done := make(chan interface{})
	closed := false
	closeDone := func() {
		if !closed {
			close(done)
		}
	}
	defer time.AfterFunc(time.Second*10, closeDone)

	var wg sync.WaitGroup
	var mux sync.Mutex
	var numWorkers int = 4

	startPartA, startPartB, startPartC := 0, 0, 0
	endPartA, endPartB, endPartC := (len(aPrimes)/numWorkers)-1, len(bPrimes)-1, len(cPrimes)-1 //(len(bPrimes)/numWorkers)-1, (len(cPrimes)/numWorkers)-1

	resultsArr := make([]chan int, numWorkers)
	//count := 0
	allCounts := make([]int, numWorkers)

	uniqueness := make(map[int]bool)

	// add to wait group
	wg.Add(numWorkers)
	// for each, start worker
	for i := 0; i < numWorkers; i++ {
		fmt.Printf("starting worker %d for ranges %d-%d, %d-%d, %d-%d\n", i, startPartA, endPartA, startPartB, endPartB, startPartC, endPartC)
		// search creates a goroutine and returns a channel for results, add it to array
		resultsArr[i] = search(startPartA, endPartA, startPartB, endPartB, startPartC, endPartC, &aPrimes, &bPrimes, &cPrimes, done)
		// start a func for receiving results
		go func(ind int) {
			// defer the waitgroup complete
			defer wg.Done()
			// for the results
			for v := range resultsArr[ind] {

				// check done
				select {
				case <-done:
					return
				default:
				}

				allCounts[ind] += 1

				mux.Lock()
				uniqueness[v] = true
				mux.Unlock()
			}
		}(i)

		partAinc := len(aPrimes) / numWorkers
		//partBinc := len(bPrimes) / numWorkers
		//partCinc := len(cPrimes) / numWorkers
		startPartA += partAinc
		//startPartB += partBinc
		//startPartC += partCinc
		endPartA = minInt(startPartA+partAinc, len(aPrimes)-1)
		//endPartB = minInt(startPartB+partBinc, len(bPrimes)-1)
		//endPartC = minInt(startPartC+partCinc, len(cPrimes)-1)
	}

	// now we wait for the results to come back
	wg.Wait()

	//fmt.Printf("uniqueness: %+v\n", uniqueness)
	fmt.Printf("count is %d\n", len(uniqueness))
	fmt.Printf("allCounts: %+v\n", allCounts)
	closeDone()
}
