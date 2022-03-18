package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {

	var maximumNumberOfConcurrentConnections int
	flag.IntVar(&maximumNumberOfConcurrentConnections, "c", 2, "maximum number of concurrent connections")
	flag.Parse()

	urls := flag.Args()

	// for i := 0; i < 5; i++ {
	// 	go func(){

	// 	}()
		
	// }
	result := fetchURLS(urls, maximumNumberOfConcurrentConnections)

	for url := range result{
		fmt.Println("Done", url)
	}

}

type Metadata struct {
	URL string 
}

func fetchURLS(urls []string, concurrency int) chan Metadata {

	processQueue := make(chan string, concurrency)

	outChan := make(chan Metadata)


	var wg sync.WaitGroup
	go func() {
		log.Println("Coordinator start")

		for _, url := range urls {

			wg.Add(1)
			processQueue <- url

			go func(url string) {

				defer wg.Done()
				pingURL(url)
				<- processQueue

				outChan <- Metadata{url}
			}(url)
		}
		wg.Wait()
		log.Println("Coordinator finished")
		close(outChan)
	}()

	return outChan
}


func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

