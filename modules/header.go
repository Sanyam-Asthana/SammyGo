package modules

import (
        "fmt"
        "net/http"
)

func ShowHeader(resp *http.Response) {
        fmt.Println("----------HEADER----------")
        for k := range resp.Header {
                fmt.Printf("%s: %s\n", k, resp.Header[k][0])
        }
        fmt.Println("--------------------------")
}
