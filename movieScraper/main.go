package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
)

func main() {
  client := &http.Client{}

  config := config()

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
