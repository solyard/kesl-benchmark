package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {

	//Check run params
	threads := flag.Int("threads", 4, "Number of gorutines that need to be running")
	scansCount := flag.Int("scansCount", 10, "Number of tasks of scanning files in KESL")
	fileSize := flag.Int("fileSize", 1000000, "File size in bytes that need to generate for KESL scan-file task")
	keslCommand := flag.String("keslCommand", "/opt/kaspersky/kesl/bin/kesl-control --scan-file %s --action Skip", "kesl-control command for file scanning")

	flag.Parse()

	tasks := make(chan *exec.Cmd, 64)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go func() {
			for cmd := range tasks {
				cmd.Stderr = os.Stderr
				start := time.Now()
				_, err := cmd.Output()
				if err != nil {
					fmt.Printf("%v", err)
				}
				elapsed := time.Since(start)
				log.Printf("Scanning took %s", elapsed)
			}
			wg.Done()
		}()
	}

	// generate some tasks
	for i := 0; i < *scansCount; i++ {
		file, err := ioutil.TempFile(".", "*.txt")
		if err != nil {
			log.Fatal(err)
		}
		filename := file.Name()
		bigBuff := make([]byte, *fileSize)
		ioutil.WriteFile(filename, bigBuff, 0666)
		tasks <- exec.Command("/bin/bash", "-c", fmt.Sprintf(*keslCommand, filename))
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
