package config

import (

  "io/ioutil"
  "encoding/json"
)

  type Config struct {
    MovieDB struct {
      ApiKey string
    }
  }

  func GetConfig(fileName string) Config {

    configJSON, _ := ioutil.ReadFile(fileName)

    config := Config{}

    if err := json.Unmarshal(configJSON, &config); err != nil {
      panic(err)
    }
    return config

  }


