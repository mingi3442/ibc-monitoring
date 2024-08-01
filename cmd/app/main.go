package main

import (
  "sync"

  "github.com/mingi3442/ibc-monitoring/pkg/ibc-monitor/domain"
  "github.com/mingi3442/ibc-monitoring/pkg/ibc-monitor/usecase"
  "github.com/mingi3442/logger"
)

func main() {
  var wg sync.WaitGroup

  cosmosConfig := domain.IBCClientConfig{
    WsUrl:       "http://localhost:11157",
    GrpcUrl:     "http://localhost:11290",
    NetworkName: "cosmos",
    Query:       "tm.event='NewBlock'",
    Subscriber:  "relayer",
  }

  osmosisConfig := domain.IBCClientConfig{
    WsUrl:       "http://localhost:11257",
    GrpcUrl:     "http://localhost:11290",
    NetworkName: "osmosis",
    Query:       "tm.event='NewBlock'",
    Subscriber:  "relayer",
  }

  cosmosIBCClient, _ := domain.IBCClientBuild(cosmosConfig)
  osmosisIBCClient, _ := domain.IBCClientBuild(osmosisConfig)

  cosmosMonitorUseCase := usecase.NewMonitorUseCase(cosmosIBCClient)
  osmosisMonitorUseCase := usecase.NewMonitorUseCase(osmosisIBCClient)

  wg.Add(2)
  go func() {
    defer wg.Done()
    if err := cosmosMonitorUseCase.StartMonitoring(); err != nil {
      logger.Fatal("Failed to start monitoring: %v", err)
    }
  }()

  go func() {
    defer wg.Done()
    if err := osmosisMonitorUseCase.StartMonitoring(); err != nil {
      logger.Fatal("Failed to start monitoring: %v", err)
    }
  }()

  wg.Wait()

}
