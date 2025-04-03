package packageManager

import (
  "encoding/json"
  "fmt"
  "io"
  "os"
)

type PackageJson struct {
  Scripts map[string]string `json:"scripts"`
}

func GetNpmScriptCommand(packageJsonPath string, script string) string {
  scripts := getNpmScripts(packageJsonPath)

  if command, ok := scripts[script]; ok {
    return command
  }

  panic(fmt.Sprintf("Cannot find script %s in file %s", script, packageJsonPath))
}

func getNpmScripts(packageJsonPath string) map[string]string {
  packageJsonRaw, err := os.Open(packageJsonPath)

  if err != nil {
    panic(err)
  }

  byteValue, _ := io.ReadAll(packageJsonRaw)

  var packageJson PackageJson

  unmarshalErr := json.Unmarshal(byteValue, &packageJson)

  if unmarshalErr != nil {
    panic(unmarshalErr)
  }

  return (&packageJson).Scripts
}
