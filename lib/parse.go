// Package shuttler parses provided JSON and maps it to a DB table.
package parse

import (

  "fmt"
  "encoding/json"
  "strings"

)

type Upcoming struct {
    Results []Result `json:"results"`
 }

type Result struct {
  PosterPath string `json:"poster_path"`
  Adult bool `json:"adult"`
  Overview string `json:"overview"`
  ReleaseDate string `json:"release_date"`
  GenreIds []int `json:"genre_ids"`
  Id int `json:"id"`
  OriginalTitle string `json:"original_title"`
  OriginaLanguage string `json:"original_language"`
  Title string `json:"title"`
  BackdropPath string `json:"backdrop_path"`
  Popularity float64 `json:"popularity"`
  VoteCount int `json:"vote_count"`
  Video bool `json:"video"`
  VoteAverage float64 `json:"vote_average"`
}

//Maps a given JSON array to a given Struct
func MapJson(s []byte) {

  //fmt.Printf("%s", s)

  dec := json.NewDecoder(strings.NewReader(string(s)))

    var m Upcoming
    for dec.More() {

      //decode an Array Value
      err := dec.Decode(&m)
      if err != nil {
        panic(err)
      }
      fmt.Printf("%+v\n", m.Results)

    }

   fmt.Printf("%+v\n", m.Results[2])

}
