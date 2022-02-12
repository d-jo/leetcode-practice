package fan_in_fan_out

import (
	"sort"
	"sync"
)

func FanOut(done <-chan interface{}, input_stream <-chan interface{}, fanoutfunc func(done <-chan interface{}, inputStream <-chan interface{}) <-chan interface{}, n_workers int) <-chan interface{} {
	var wg sync.WaitGroup

	output := make(chan interface{})

	wg.Add(n_workers)
	// for n workers
	for i := 0; i < n_workers; i++ {
		// start a goroutine
		go func() {
			// defer waitgroup
			defer wg.Done()

			// for each output of fanoutfunc
			for midout := range fanoutfunc(done, input_stream) {
				// check done
				select {
				case <-done:
					return
				default:
					output <- midout // send to output
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// takes in a channel of ints and returns a channel of ints. all output are prime, if it isnt prime it is trashed
func FilterPrime(done <-chan interface{}, input_stream <-chan interface{}) <-chan interface{} {

	output := make(chan interface{})

	go func() {
		defer close(output)
	stream:
		for numRaw := range input_stream {
			select {
			case <-done:
				return
			default:
			}
			num := numRaw.(int64)

			for i := int64(2); i < num; i++ {
				if num%i == 0 {
					// if its not prime, dump it and continue outter loop
					continue stream
				}
			}

			// if we get here, its prime
			output <- num
		}
	}()

	return output
}

func Start(max int, n_workers int) []int64 {

	done := make(chan interface{})
	defer close(done)

	input_stream := make(chan interface{})
	//defer close(input_stream)

	// start the fanout
	output := FanOut(done, input_stream, FilterPrime, n_workers)

	// start the input
	go func() {
		for i := 2; i < max; i++ {
			input_stream <- int64(i)
		}
		//close(done)
		close(input_stream)
	}()

	// start the output
	var result []int64
	for num := range output {
		result = append(result, num.(int64))
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	return result
}

func StartNoSort(max int, n_workers int) []int64 {

	done := make(chan interface{})
	defer close(done)

	input_stream := make(chan interface{})
	//defer close(input_stream)

	// start the fanout
	output := FanOut(done, input_stream, FilterPrime, n_workers)

	// start the input
	go func() {
		for i := 2; i < max; i++ {
			input_stream <- int64(i)
		}
		//close(done)
		close(input_stream)
	}()

	// start the output
	var result []int64
	for num := range output {
		result = append(result, num.(int64))
	}

	//sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	return result
}
