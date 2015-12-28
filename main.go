package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "encoding/json"
)

func main() {
  client := &http.Client{}

  configJSON, err := ioutil.ReadFile("./config/development.json")

  type Config struct {
    MovieDB struct {
      ApiKey string
    }
  }

  config := Config{}

  if err2 := json.Unmarshal(configJSON, &config); err2 != nil {
    panic(err2)
  }
  fmt.Printf("%+v\n", config)

  req, _ := http.NewRequest("GET", "http://api.themoviedb.org/3/movie/upcoming", nil)

  req.Header.Add("Accept", "application/json")

  resp, err := client.Do(req)

  if err != nil {
    fmt.Println("Errored when sending request to the server")
    return
  }

  defer resp.Body.Close()
  resp_body, _ := ioutil.ReadAll(resp.Body)

  fmt.Println(resp.Status)
  fmt.Println(string(resp_body))
}
