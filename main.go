package main

import (
    "context"
    "fmt"
    "github.com/google/go-github/v74/github"
    "github.com/sirupsen/logrus"
    "github.com/gorilla/mux"
    "github.com/spf13/viper"
    "github.com/go-resty/resty/v2"
    "math/rand"
    "time"
    "encoding/json"
    "os"
    "strings"
    "net/http"
)

func main() {
    client := github.NewClient(nil)
    fmt.Println("Using go-github library for GitHub API interaction")

    // math/rand
    fmt.Println("Random number:", rand.Intn(100))

    // time
    fmt.Println("Current time:", time.Now())

    // encoding/json
    data, _ := json.Marshal(map[string]string{"hello": "world"})
    fmt.Println("JSON marshal:", string(data))

    // os
    fmt.Println("Current working directory:", func() string { d, _ := os.Getwd(); return d }())

    // strings
    fmt.Println("To upper:", strings.ToUpper("go-gen-two"))

    // net/http
    fmt.Println("HTTP package available:", http.MethodGet)

    // logrus
    logrus.Info("This is a logrus info message!")

    // gorilla/mux
    r := mux.NewRouter()
    fmt.Println("Gorilla/mux router created:", r != nil)

    // viper
    viper.Set("appName", "go-gen-two")
    fmt.Println("Viper config appName:", viper.GetString("appName"))

    // resty
    clientResty := resty.New()
    fmt.Println("Resty client created:", clientResty != nil)
}
