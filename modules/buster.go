package modules

import (
	"fmt"
	"net/http"
//	"log"
//	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"
//	"io"
)

func checkForErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func checkPath(path string) (string, bool) {

	fmt.Fprintf(os.Stdout, "\rChecking %s", path)

	resp, err := http.Get(path)

	if err != nil {
		return path, false
	}

	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		return path, true
	}

	return path, false
}

func printResult(path string, wg *sync.WaitGroup) {

	defer wg.Done()

	url, result := checkPath(path)

	if result {
		fmt.Printf("Found: %s\n", url)
	}
}


func BruteForce(baseURL string) {
	
	dir, err := os.Getwd()

	checkForErrors(err)

	wordlistPath := filepath.Join(dir, "wordlist.txt")
	dat, err := os.ReadFile(wordlistPath);

	checkForErrors(err)

	paths := strings.Split(string(dat), "\n")

	var wg sync.WaitGroup

	fmt.Println("Initiating buster...")

	for _, path := range paths {
		fullPath := fmt.Sprintf("%s/%s", baseURL, path)
		wg.Add(1)
		go printResult(fullPath, &wg)	
	}	

	wg.Wait()
}





