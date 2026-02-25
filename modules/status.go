package modules

import (
	"fmt"
	"net/http"
)

// ShowStatusCode shows the status of the response
// Param: resp *http.Response (The pointer to the http Response)
// Returns: void
func ShowStatusCode(resp *http.Response) {
	fmt.Println()
        fmt.Println("--------------------")
        fmt.Println("Status Code: " + resp.Status)
        fmt.Println("--------------------")
        fmt.Println()
}


