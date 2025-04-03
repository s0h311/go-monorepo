package packageManager

import (
  "encoding/json"
  "fmt"
  "io"
  "os"
)

type ComposerJson struct {
  Scripts map[string]string `json:"scripts"`
}

func GetComposerScriptCommand(composerJsonPath string, script string) string {
  scripts := getComposerScripts(composerJsonPath)

  if command, ok := scripts[script]; ok {
    return command
  }

  panic(fmt.Sprintf("Cannot find script %s in file %s", script, composerJsonPath))
}

func getComposerScripts(composerJsonPath string) map[string]string {
  composerJsonRaw, err := os.Open(composerJsonPath)

  if err != nil {
    panic(err)
  }

  byteValue, _ := io.ReadAll(composerJsonRaw)

  var composerJson ComposerJson

  unmarshalErr := json.Unmarshal(byteValue, &composerJson)

  if unmarshalErr != nil {
    panic(unmarshalErr)
  }

  return (&composerJson).Scripts
}
