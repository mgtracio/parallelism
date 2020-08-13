package main

import (
	"fmt"
	"latamautos/challenge/benchmark/config"
	"math"
	"runtime"
	"sync"
	"time"
)

/**
 * Parallel Batch Processing in Golang (mainImages1)::: INTENSIVE TASKS SUCCESSFULLY COMPLETED -> Elapsed time
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
func main() {
	runtime.GOMAXPROCS(4)
	var totalSum int64 = 0
	worker := func(numbers []int, wg *sync.WaitGroup) {
		defer wg.Done()
		var sum int64 = 0
		for _, number := range numbers {
			sum += int64(number)
		}
		fmt.Printf("+++ WORKER FINISHED (sum: %d of %d elements)\n", sum, len(numbers))
		totalSum = totalSum + sum
	}

	start := time.Now()
	var wg sync.WaitGroup // Goroutines Group
	segmentsPerWorker := int(math.Ceil(float64(len(config.Config.Numbers)) / float64(config.Config.WorkersForNumbers)))
	for i := 0; i < len(config.Config.Numbers); i += segmentsPerWorker {
		wg.Add(1)
		numbers := config.Config.Numbers[i:config.MinValue(i+segmentsPerWorker, len(config.Config.Numbers))]
		go worker(numbers, &wg)
	}
	wg.Wait() // Wait for Goroutines Group to finish
	duration := time.Since(start)
	fmt.Printf("::: INTENSIVE TASKS SUCCESSFULLY COMPLETED -> Workers: %d -> Elapsed time: %2d ms -> Total Sum: %d) \n", config.Config.WorkersForNumbers, duration.Milliseconds(), totalSum)
}