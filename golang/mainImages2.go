package main

import (
	"fmt"
	"github.com/h2non/bimg"
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

	worker := func(images []os.FileInfo, wg *sync.WaitGroup)  {
		defer wg.Done()
		for _, img := range images {
			options := bimg.Options{Width:config.Config.X, Height: config.Config.Y, Quality: config.Config.Quality}
			buffer, _ := bimg.Read(config.Config.Dot+config.Config.Slash+config.Config.ImagePath+config.Config.Slash+img.Name())
			newImage, _ := bimg.NewImage(buffer).Process(options)
			bimg.Write(config.Config.Dot+config.Config.Slash+config.Config.ImageResultPath+config.Config.Slash+img.Name(), newImage)
		}
		fmt.Printf("+++ WORKER FINISHED (compressed files: %d)\n", len(images))
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