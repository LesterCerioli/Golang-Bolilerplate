package main


import (
    "log"
    "net/http"
    "{{ .ProjectName}}/internal/handler"
)

func main() {
    httpHandleFunc("/", handler.HomeHandler)
    log.Println("{{ .ProjectName }} started on port 8081...")
    log.Fatal(http.ListenAndServe(":8081", nil))
}