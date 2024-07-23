package main

import (
  ibc "github.com/mingi3442/ibc-monitoring/client"
  // ws "github.com/mingi3442/ibc-monitoring/client/ws"
  // grpc "github.com/mingi3442/ibc-monitoring/client/grpc"
  // "sync"
)

func main() {
  //* IBC Client
  cosmosIBCClient, _ := ibc.Connect("http://localhost:11157", "http://localhost:11290", "cosmos")
  osmosisIBCClient, _ := ibc.Connect("http://localhost:11257", "http://localhost:11290", "osmosis")

  cosmosIBCClient.UpdateQuery("tm.event='NewBlock'")
  osmosisIBCClient.UpdateQuery("tm.event='NewBlock'")
  cosmosIBCClient.UpdateSubscriber("relayer")
  osmosisIBCClient.UpdateSubscriber("relayer")
  go cosmosIBCClient.Subscribe()
  go osmosisIBCClient.Subscribe()

  defer func() {
    cosmosIBCClient.DisConnect()
    osmosisIBCClient.DisConnect()
  }()
  select {}

}
