package modules

import (
	"fmt"
)

// ShowBody shows the body of the response
// Param: resp *http.Response (The pointer to the http Response object)
// Returns: void
func ShowBody(body string) {
	fmt.Println("-------------BODY-------------")
	fmt.Println(body)
	fmt.Println("------------------------------")
}
