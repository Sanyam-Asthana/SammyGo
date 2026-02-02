package modules

import (
	"fmt"
	"net/http"
)

func ShowStatusCode(resp *http.Response) {
	fmt.Println()
        fmt.Println("--------------------")
        fmt.Println("Status Code: " + resp.Status)
        fmt.Println("--------------------")
        fmt.Println()
}


