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
  return saveJsonFile(formatData, dir, "action_transaction")
}

func SaveTransactionData(transaction interface{}, dir string) error {
  return saveJsonFile(transaction, dir, "basic_transaction")
}

func saveJsonFile(saveData interface{}, dir, filename string) error {
  err := ensureDir(dir)
  if err != nil {
    return err
  }

  data, err := json.MarshalIndent(saveData, "", "  ")
  if err != nil {
    return fmt.Errorf("failed to marshal transaction: %w", err)
  }

  jsonFileName := fmt.Sprintf("%s_transaction_%s.json", filename, time.Now().Format("2006-01-02T15-04-05"))

  filepath := filepath.Join(dir, jsonFileName)

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
