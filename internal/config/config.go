package config

import (
  "encoding/json"
  "github.com/s0h311/go-monorepo/internal/utils"
  "io"
  "os"
)

type Config struct {
  Packages []string `json:"packages"`
}

func GetConfig() *Config {
  configJson, err := os.Open("./mono.config.json")

  if err != nil {
    panic(err)
  }

  defer utils.TryOrThrow(configJson.Close)

  byteValue, _ := io.ReadAll(configJson)

  var config Config

  unmarshalErr := json.Unmarshal(byteValue, &config)

  if unmarshalErr != nil {
    panic(unmarshalErr)
  }

  return &config
}
