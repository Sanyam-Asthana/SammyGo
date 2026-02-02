package modules

import (
	"fmt"
//	"net/http"
//	"log"
//	"bufio"
	"os"
	"path/filepath"
//	"io"
)

func checkForErrors(err error) {
	if err != nil {
		panic(err)
	}
}


func BruteForce() {
	
	dir, err := os.Getwd()

	checkForErrors(err)

	wordlistPath := filepath.Join(dir, "wordlist.txt")
	dat, err := os.ReadFile(wordlistPath);

	checkForErrors(err)

	words := string(dat)

	fmt.Println(words)

}





