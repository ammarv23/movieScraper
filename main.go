package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "net/url"

  "github.com/ammarv23/movieScraper/lib"
  "github.com/ammarv23/movieScraper/app"
)


func main() {
  config := config.GetConfig("./config/development.json")

  u, err := url.Parse("http://api.themoviedb.org/3/movie/upcoming")
  if err != nil {
    panic(err)
  }

  q := u.Query()
  q.Set("api_key", config.MovieDB.ApiKey)
  u.RawQuery = q.Encode()
  fmt.Println(u.String())

  client := &http.Client{}

  req, _ := http.NewRequest("GET", u.String(), nil)

  req.Header.Add("Accept", "application/json")

  resp, err := client.Do(req)

  if err != nil {
    fmt.Println("Errored when sending request to the server")
    return
  }

  defer resp.Body.Close()
  resp_body, _ := ioutil.ReadAll(resp.Body)

  parse.MapJson(resp_body)
}
