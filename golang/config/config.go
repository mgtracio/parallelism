package config

import (
	"io/ioutil"
	"math/rand"
	"time"
)

/**
 * Parallel Batch Processing in Golang (Config)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
type config struct {
	X						int
	Y						int
	Quality	 				int
	ImagePath				string
	ImageResultPath			string
	WorkersForImages		int
	Slash					string
	Dot						string

	WorkersForNumbers		int
	Numbers					[]int
}

var Files, _ = ioutil.ReadDir(Config.Dot + Config.Slash + Config.ImagePath)
func MinValue(x int , y int) int {
	if x <= y {return x}
	return y
}

func GetNumbers(limit int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, limit)
	for i := range numbers {numbers[i] = rand.Intn(10 - 1 + 1) + 1}
	return numbers
}
var Config = config{800, 600, 80, "images", "imagesResult", 20, "/", ".", 10, GetNumbers(50000000)}