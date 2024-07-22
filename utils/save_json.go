package utils

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "time"
)

type ActionData struct {
  NetworkName string   `json:"network"`
  Action      []string `json:"action"`
}

func ensureDir(dir string) error {
  if _, err := os.Stat(dir); os.IsNotExist(err) {
    err := os.MkdirAll(dir, os.ModePerm)
    if err != nil {
      return fmt.Errorf("failed to create directory: %w", err)
    }
  }
  return nil
}

func SaveActionData(actionData []string, networkName, dir string) error {
  formatData := ActionData{
    Action:      actionData,
    NetworkName: networkName,
  }
  err := ensureDir(dir)
  if err != nil {
    return err
  }

  data, err := json.MarshalIndent(formatData, "", "  ")
  if err != nil {
    return fmt.Errorf("failed to marshal action: %w", err)
  }

  filename := fmt.Sprintf("%s_action_%s.json", networkName, time.Now().Format("2006-01-02T15-04-05"))

  filepath := filepath.Join(dir, filename)

  file, err := os.Create(filepath)
  if err != nil {
    return fmt.Errorf("failed to create file: %w", err)
  }
  defer file.Close()

  _, err = file.Write(data)
  if err != nil {
    return fmt.Errorf("failed to write data to file: %w", err)
  }

  return nil
}
