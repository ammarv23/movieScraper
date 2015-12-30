package main

import (

  "io/ioutil"
  "encoding/json"
)

  type Config struct {
    MovieDB struct {
      ApiKey string
    }
  }

  func config(fileName string) Config {

    configJSON, _ := ioutil.ReadFile(fileName)

    config := Config{}

    if err := json.Unmarshal(configJSON, &config); err != nil {
      panic(err)
    }
    return config

  }


