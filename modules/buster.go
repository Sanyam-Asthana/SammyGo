package modules

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
)

func worker(jobsChannel <-chan string, wg *sync.WaitGroup, filterCode int, verbose bool, progressPtr *int64) {
	defer wg.Done()
	for url := range jobsChannel {
		checkPath(url, filterCode, verbose, progressPtr)
	}
}

func checkPath(path string, filterCode int, verbose bool, progressPtr *int64) {
	
	currentProgress := atomic.AddInt64(progressPtr, 1)

	if verbose {
		fmt.Fprintf(os.Stdout, "\rChecking: %d", currentProgress)
	}

	resp, err := http.Get(path)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		if verbose {
			fmt.Println("\n")
		}
		if filterCode == 0 || resp.StatusCode == filterCode {
			fmt.Printf("%s[+] Found: %s%s %s(%d)%s\n", ColorGreen, ColorYellow, path, ColorRed, resp.StatusCode, ColorReset)
		}
	}
}

// BruteForce initiates a dictionary attack on the provided URL
// Param: baseURL string (The base URL to initiate the attack on)
// Param: numWorkers int (The number of workers to work together for busting)
// Param: filterCode int (The status code by which the outputs have to be filtered. 0 for no filtration)
// Param: verbose bool (Whether to show the progress while busting)
// Param: wordlistPath string (The path to the dictionary for the dictionary attack)
// Returns: void
func BruteForce(baseURL string, numWorkers int, filterCode int, verbose bool, wordlistPath string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting directory:", err)
		return
	}

	if(wordlistPath == "wordlist.txt") {
		wordlistPath = filepath.Join(dir, "wordlist.txt")
	}

	dat, err := os.ReadFile(wordlistPath)
	if err != nil {
		fmt.Println("Wordlist not found at:", wordlistPath)
		return
	}

	paths := strings.Split(string(dat), "\n")
	baseURL = strings.TrimSuffix(baseURL, "/")

	jobsChannel := make(chan string, len(paths))
	var wg sync.WaitGroup

	fmt.Printf("Initiating buster with %d workers...\n", numWorkers)

	var progress int64
	progress = 0

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(jobsChannel, &wg, filterCode, verbose, &progress)
	}

	for _, path := range paths {
		if trimmed := strings.TrimSpace(path); trimmed != "" {
			jobsChannel <- fmt.Sprintf("%s/%s", baseURL, trimmed)
		}
	}
	close(jobsChannel) 
	wg.Wait()
	fmt.Println("\nScan complete.")
}
