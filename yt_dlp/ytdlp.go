package yt_dlp

import (
	"bufio"
	"fmt"
	"log"
	"sync"
)

type Ytdlp struct {
	semaphore         chan struct{}
	Url               string // uri -> eg -> https://www.bilibili.com/video/BV1PY411P7D1
	CookieFromBrowser string // --cookies-from-browser -> eg -> chrome
	//TODO ... You can add the rest yourself
}

func (y *Ytdlp) Init() {
	y.semaphore = make(chan struct{}, 5)
}

func (y *Ytdlp) MultiThreadedDownload() (err error) {
	// Control the number of goroutines to 5

	// create waitGroup
	var wg sync.WaitGroup
	var total int = 1
	for total > 0 {
		wg.Add(1)
		// When Ctrip starts,
		// write an empty strcut into the semaphore.
		// If it is full, it will block.
		// This will achieve the effect of controlling
		// the number of Ctrip concurrencies.
		y.semaphore <- struct{}{}
		go func(index int) {
			defer func() {
				<-y.semaphore
			}()
			defer wg.Done()

			// get stdout
			cmd := y.toCmdString(index)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				log.Printf("Error getting stdout pipe: %v\n", err)
				return
			}
			defer stdout.Close()

			// get stderr
			stderr, err := cmd.StderrPipe()
			if err != nil {
				log.Printf("Error getting stderr pipe: %v\n", err)
				return
			}
			defer stderr.Close()

			// get scanner for real-time printing
			outScanner := bufio.NewScanner(stdout)
			errScanner := bufio.NewScanner(stderr)
			if err := cmd.Start(); err != nil {
				log.Printf("Error starting command: %v\n", err)
				return
			}

			for outScanner.Scan() {
				// Here is the normal log output
				fmt.Println("out", outScanner.Text())
			}
			for errScanner.Scan() {
				// If err appears here,
				// it is not a network problem,
				// then it is a failure caused by invalid subsequent page numbers.
				// You can judge the specific error by yourself.
				fmt.Println("err", errScanner.Text())
				// Set total to 0 to break out of the loop
				total = 0
			}

			// Wait for cmd to complete
			if err := cmd.Wait(); err != nil {
				log.Printf("Error waiting for command to finish: %v\n", err)
			}
			return
		}(total)

		if total > 0 {
			total++
		}
	}
	// Wait for goroutines to complete
	wg.Wait()
	return nil
}
