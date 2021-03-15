package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	// Get file pat
	filePath := flag.String("file", "", "File path")
	fileThreshold := flag.Int("threshold", 1, "File threshold size in MB")
	flag.Parse()

	if *filePath == "" {
		log.Fatalln("please specify a file path with --file")
	}

	_, err := os.Stat(*filePath)

	if os.IsNotExist(err) {
		log.Fatalln("file does not exists", err)
	}

	buf, err := ioutil.ReadFile(*filePath)
	strArr := filterStrings(strings.Split(string(buf), "\n"))

	os.RemoveAll("./output")

	if isSorted(strArr) {
		log.Println("Text is already sorted")
		os.Exit(0)
	}

	sort.Strings(strArr)

	buf = []byte(strings.Join(strArr, "\n"))

	// Free space
	strArr = nil

	maxFileSize := (*fileThreshold) * 1_000_000 // Convert to mb
	numberOfDirectories := math.Ceil(float64(len(buf)) / float64(maxFileSize))

	bufReminder := len(buf) % maxFileSize

	smallerFilesSize := 500_000
	start := 0

	for i := 1; i <= int(numberOfDirectories); i++ {

		wg.Add(1)
		// Push to a go routine
		go func(i, numberOfDirectories, start, smallerFilesSize, maxFileSize, bufReminder int, group *sync.WaitGroup) {
			defer wg.Done()
			err := os.Mkdir(fmt.Sprintf("./output/%d", i), os.ModePerm)
			if err != nil {
				log.Fatalln("error occurred creating output directory", err)
			}

			if i == numberOfDirectories && bufReminder > 0 {
				smallerFilesSize = bufReminder
			}

			filesInDirectory := math.Floor(float64(maxFileSize) / float64(smallerFilesSize))

			for j := 1; j <= int(filesInDirectory); j++ {
				file, err := os.OpenFile(fmt.Sprintf("./output/%d/%d.%d.txt", i, i, j), os.O_CREATE|os.O_RDWR, os.ModePerm)
				if err != nil {
					log.Fatalln("error occurred")
				}

				file.Write(buf[start : smallerFilesSize+start])
			}
		}(i, int(numberOfDirectories), start, smallerFilesSize, maxFileSize, bufReminder, wg)

		start += smallerFilesSize

	}

	wg.Wait()

}

func filterStrings(strArr []string) (ret []string) {
	for _, s := range strArr {
		if s != "" {
			ret = append(ret, s)
		}
	}
	return
}

func isSorted(strArr []string) bool {
	sorted := false
	for i := 0; i < len(strArr); i++ {
		if (strArr)[i] < (strArr)[i+1] {
			sorted = true
		} else {
			return false
		}
	}
	return sorted
}
