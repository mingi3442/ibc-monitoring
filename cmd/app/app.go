package app

import (
  "sync"

  "github.com/mingi3442/ibc-monitoring/pkg/ibcmonitor"
)

func Start() {
  var wg sync.WaitGroup

  cosmosConfig := ibc_monitor.IBCClientConfigParamsBuild("http://localhost:11157", "http://localhost:11190", "cosmos", "tm.event='NewBlock' OR tm.event='Tx'", "relayer")
  osmosisConfig := ibc_monitor.IBCClientConfigParamsBuild("http://localhost:11257", "http://localhost:11290", "osmosis", "tm.event='NewBlock'", "relayer")

  cosmosIBCClient, _ := ibc_monitor.IBCClientBuild(cosmosConfig)
  osmosisIBCClient, _ := ibc_monitor.IBCClientBuild(osmosisConfig)

  cosmosMonitorUseCase := ibc_monitor.NewMonitorUseCase(cosmosIBCClient)
  osmosisMonitorUseCase := ibc_monitor.NewMonitorUseCase(osmosisIBCClient)

  wg.Add(2)
  go func() {
    defer wg.Done()
    cosmosMonitorUseCase.StartMonitoring()

  }()

  go func() {
    defer wg.Done()
    osmosisMonitorUseCase.StartMonitoring()
  }()

  wg.Wait()

}
