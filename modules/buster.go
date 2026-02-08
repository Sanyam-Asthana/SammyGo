package modules

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	ColorRed    = "\x1b[31m"
	ColorGreen  = "\x1b[32m"
	ColorYellow = "\x1b[33m"
	ColorReset  = "\x1b[0m"
)

func worker(jobs <-chan string, wg *sync.WaitGroup, filterCode int, verbose bool) {
	defer wg.Done()
	for url := range jobs {
		checkPath(url, filterCode, verbose)
	}
}

func checkPath(path string, filterCode int, verbose bool) {
	if verbose {
		fmt.Fprintf(os.Stdout, "\rChecking: %-60s", path)
	}

	resp, err := http.Get(path)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		if filterCode == 0 {
			fmt.Printf("%s[+] Found: %s%s %s(%d)%s\n", ColorGreen, ColorYellow, path, ColorRed, resp.StatusCode, ColorReset)
		} else {
			if resp.StatusCode == filterCode {
				fmt.Printf("%s[+] Found: %s%s %s(%d)%s\n", ColorGreen, ColorYellow, path, ColorRed, resp.StatusCode, ColorReset)
			}
		}
	}
}

func BruteForce(baseURL string, numWorkers int, filterCode int, verbose bool) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting directory:", err)
		return
	}

	wordlistPath := filepath.Join(dir, "wordlist.txt")
	dat, err := os.ReadFile(wordlistPath)
	if err != nil {
		fmt.Println("Wordlist not found at:", wordlistPath)
		return
	}

	paths := strings.Split(string(dat), "\n")
	baseURL = strings.TrimSuffix(baseURL, "/")

	jobs := make(chan string, len(paths))
	var wg sync.WaitGroup

	fmt.Printf("Initiating buster with %d workers...\n", numWorkers)

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(jobs, &wg, filterCode, verbose)
	}

	for _, path := range paths {
		if trimmed := strings.TrimSpace(path); trimmed != "" {
			jobs <- fmt.Sprintf("%s/%s", baseURL, trimmed)
		}
	}
	close(jobs) 
	wg.Wait()
	fmt.Println("\nScan complete.")
}
