package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"
	"github.com/Sanyam-Asthana/sammygo/modules"
)

func main() {
	fmt.Println(" ____                                   ____       ")
	fmt.Println("/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___  ")
	fmt.Println("\\___ \\ / _` | '_ ` _ \\| '_ ` _ \\| | | | |  _ / _ \\ ")
	fmt.Println(" ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |")
	fmt.Println("|____/ \\__,_|_| |_| |_|_| |_| |_|\\__, |\\____|\\___/ ")
	fmt.Println("                                 |___/             ")
	
	textBoolFlagPtr := flag.Bool("text", false, "Get the response text")
	headBoolFlagPtr := flag.Bool("head", false, "Get the response header")
	bustBoolFlagPtr := flag.Bool("bust", false, "Brute-force web directories")
	workerIntFlagPtr := flag.Int("workers", 20, "Number of workers for directory busting")
	filterCodeIntFlagPtr := flag.Int("filter", 0, "Only show brute-forced paths with a specific status code")
	verboseBoolFlagPtr := flag.Bool("verbose", false, "Verbose output for brute-forcing")
	wordlistPathStringFlagPtr := flag.String("wordlist", "../../wordlists/wordlist.txt", "The path for the wordlist for busting")

	flag.Parse()

	url := flag.Arg(0)

	if(url == "") {
		fmt.Println("Please specify a URL!")
		return
	}

	fmt.Println("\nInitiated SammyGo on " + url)

	resp, err := http.Get(url);

	if err != nil {
		fmt.Println("SammyGo terminated with an error. Does the website exist?")
		log.Fatal(err)
	}

	modules.ShowStatusCode(resp)	
	
	if *headBoolFlagPtr {
		modules.ShowHeader(resp)
	}
	 
	if *textBoolFlagPtr {
		modules.ShowBody(resp)
	}
	
	if *bustBoolFlagPtr {
		modules.BruteForce(url, *workerIntFlagPtr, *filterCodeIntFlagPtr, *verboseBoolFlagPtr, *wordlistPathStringFlagPtr)
	}

	resp.Body.Close()
}



