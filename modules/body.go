package modules

import (
	"fmt"
	"net/http"
	"io"
	"log"
)

func ShowBody(resp *http.Response){
	defer resp.Body.Close()
	var body, errBody = io.ReadAll(resp.Body)

	if errBody != nil {
		log.Fatal(errBody)
	}
	
	fmt.Println("-------------BODY-------------")
	fmt.Println(string(body))
	fmt.Println("------------------------------")
}

