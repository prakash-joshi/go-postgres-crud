package main

import (
    "go-postgres-crud/router"
    "fmt"
    "log"
    "net/http"
)

func main() {

    router:= router.Router()

    fmt.Println("Starting server on Port 8080....")

    log.Fatal(http.ListenAndServe(":8080",router))
}
