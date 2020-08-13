package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"latamautos/challenge/benchmark/config"
	"math"
	"os"
	"sync"
	"time"
)

/**
 * Parallel Batch Processing in Golang (mainImages1)::: INTENSIVE TASKS SUCCESSFULLY COMPLETED -> Elapsed time
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
func main() {
	//runtime.GOMAXPROCS(3)
	worker := func(images []os.FileInfo, wg *sync.WaitGroup)  {
		defer wg.Done()
		for _, img := range images {
			src, _ := imaging.Open(config.Config.Dot+config.Config.Slash+config.Config.ImagePath+config.Config.Slash+img.Name())
			imaging.Save(imaging.Fit(src, config.Config.X, config.Config.Y, imaging.Lanczos), "."+config.Config.Slash+config.Config.ImageResultPath+config.Config.Slash+img.Name())
		}
		fmt.Printf("+++ WORKER FINISHED (processed images: %d)\n", len(images))
	}

	start := time.Now()
	var wg sync.WaitGroup // Goroutines Group
	segmentsPerWorker := int(math.Ceil(float64(len(config.Files)) / float64(config.Config.WorkersForImages)))
	for i := 0; i < len(config.Files); i += segmentsPerWorker {
		wg.Add(1)
		images := config.Files[i:config.MinValue(i+segmentsPerWorker, len(config.Files))]
		go worker(images, &wg)
	}
	wg.Wait() // Wait for Goroutines Group to finish
	duration := time.Since(start)
	fmt.Printf("::: INTENSIVE TASKS SUCCESSFULLY COMPLETED -> Workers: %1d -> Elapsed time: %2d ms", config.Config.WorkersForImages, duration.Milliseconds())
}