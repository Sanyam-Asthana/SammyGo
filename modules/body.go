package modules

import (
	"fmt"
	"io"
	"net/http"
)

// ShowBody shows the body of the response
// Param: resp *http.Response (The pointer to the http Response object)
// Returns: void
func ShowBody(resp *http.Response) {
	var body, errBody = io.ReadAll(resp.Body)

	if errBody != nil {
		fmt.Printf("%sError loading request body!\n", ColorRed)
		return
	}

	fmt.Println("-------------BODY-------------")
	fmt.Println(string(body))
	fmt.Println("------------------------------")
}
