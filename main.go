package main

import (
	"flag"
	"fmt"
	"github.com/sheghun/xp/lib"
	"io"
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
	memoryLimit := flag.Int("memoryLimit", 1, "File threshold size in MB")

	flag.Parse()

	if *filePath == "" {
		log.Fatalln("please specify a file path with --file")
	}

	fileStats, err := os.Stat(*filePath)

	if os.IsNotExist(err) {
		log.Fatalln("file does not exists", err)
	}

	os.RemoveAll("./output") // Delete output folder

	// Maximum size of bytes to read at a time
	memoryLimitBytes := *memoryLimit * 1_000_000 // Convert mb to bytes

	file, err := os.OpenFile(*filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("error occurred when opening file")
	}

	buf := make([]byte, memoryLimitBytes)
	fileSize := fileStats.Size()
	readTimes := math.Ceil(float64(fileSize) / float64(memoryLimitBytes))
	reminder := fileSize % int64(memoryLimitBytes)
	firstDataSorted := false

	err = os.Mkdir("./output", os.ModePerm)

	if err != nil {
		log.Fatalln("error occurred when creating output directory", err)

	}

	for i := 1; i <= int(readTimes); i++ {

		if i == int(readTimes) && reminder > 0 {
			buf = make([]byte, reminder)
		}

		_, err := file.Read(buf)
		if err != nil {
			log.Fatalln("error occured reading file", err)
		}

		strArr := lib.FilterStrings(strings.Split(string(buf), "\n"))

		// Check if the first two strings read are sorted
		if i <= 2 && lib.IsSorted(strArr) {
			if i == 1 {
				firstDataSorted = true
			}

			if firstDataSorted && i == 2 {
				fmt.Println("File is already sorted")
				os.Exit(0)
			}
		}

		wg.Add(1)
		go sortAndWriteFiles(strArr, i, memoryLimitBytes, wg)
	}

	wg.Wait()

}

func sortAndWriteFiles(strArr []string, index int, maxFileSizeBytes int, wg *sync.WaitGroup) {
	defer wg.Done()

	smallerFilesSizeBytes := 500_000

	err := os.MkdirAll(fmt.Sprintf("./output/%d", index), os.ModePerm)
	if err != nil {
		log.Fatalf("error occured creating file output/%d.txt\n", index)
	}

	sort.Strings(strArr)

	filesInDirectory := math.Floor(float64(maxFileSizeBytes) / float64(smallerFilesSizeBytes))

	for j := 1; j <= int(filesInDirectory); j++ {
		file, err := os.OpenFile(fmt.Sprintf("./output/%d/%d.%d.txt", index, index, j), os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln("error occurred")
		}

		io.WriteString(file, strings.Join(strArr, "\n"))
	}

}
