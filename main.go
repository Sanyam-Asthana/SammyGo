package main

import (
	"fmt"
	"net/http"
	"log"
	"flag"
	"com.github/sammygo/modules"
)

func main() {
	fmt.Println(" ____                                   ____       ")
	fmt.Println("/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___  ")
	fmt.Println("\\___ \\ / _` | '_ ` _ \\| '_ ` _ \\| | | | |  _ / _ \\ ")
	fmt.Println(" ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |")
	fmt.Println("|____/ \\__,_|_| |_| |_|_| |_| |_|\\__, |\\____|\\___/ ")
	fmt.Println("                                 |___/             ")
	
	fmt.Println("\nSanyam Asthana, 2026")

	textBoolPtr := flag.Bool("text", false, "Get the response text")
	headBoolPtr := flag.Bool("head", false, "Get the response header")
	bustBoolPtr := flag.Bool("bust", false, "Brute-force web directories")
	workerIntPtr := flag.Int("workers", 20, "Number of workers for directory busting")
	filterCodeIntPtr := flag.Int("filter", 0, "Only show brute-forced paths with a specific status code")
	verboseBoolPtr := flag.Bool("verbose", false, "Verbose output for brute-forcing")

	flag.Parse()

	url := flag.Arg(0)	
	fmt.Println("\nInitiated SammyGo on " + url)

	resp, err := http.Get(url);

	if err != nil {
		fmt.Println("SammyGo terminated with an error. Does the website exist?")
		log.Fatal(err)
	}

	modules.ShowStatusCode(resp)	
	
	if *headBoolPtr == true {
		modules.ShowHeader(resp)
	}
	 
	if *textBoolPtr == true {
		modules.ShowBody(resp)
	}
	
	if *bustBoolPtr == true {
		modules.BruteForce(url, *workerIntPtr, *filterCodeIntPtr, *verboseBoolPtr)
	}

	resp.Body.Close()
}



