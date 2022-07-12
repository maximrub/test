package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {

    // handle `/` route to `http.DefaultServeMux`
    http.HandleFunc( "/", func( res http.ResponseWriter, req *http.Request ) {

        // get response headers
        header := res.Header()

        // set content type header
        header.Set("Content-Type","application/json")

        // reset date header (inline call)
        res.Header().Set("Date","01/01/2020")
        
        // respond with a JSON string
        fmt.Fprint(res, req)
    } )

    // listen and serve using `http.DefaultServeMux`
    log.Fatal(http.ListenAndServe(":80", nil))

}
