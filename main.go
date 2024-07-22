package main

import (
  ibc "github.com/mingi3442/ibc-monitoring/client"
  // ws "github.com/mingi3442/ibc-monitoring/client/ws"
  // grpc "github.com/mingi3442/ibc-monitoring/client/grpc"
  // "sync"
)

func main() {
  //* Websocket Client
  cosmosIBCClient, _ := ibc.Connect("http://localhost:11157", "http://localhost:11290", "cosmos")
  osmosisIBCClient, _ := ibc.Connect("http://localhost:11257", "http://localhost:11290", "osmosis")
  // cosmosWsClient, _ := ws.Connect("http://localhost:11157", "cosmos")
  // osmosisWsClient, _ := ws.Connect("http://localhost:11257", "osmosis")
  // polkachuCosmosWsClient, _ := ws.Connect("https://cosmos-rpc.polkachu.com", "Polkachu Cosmos")
  cosmosIBCClient.UpdateQuery("tm.event='NewBlock'")
  osmosisIBCClient.UpdateQuery("tm.event='NewBlock'")
  cosmosIBCClient.UpdateSubscriber("relayer")
  osmosisIBCClient.UpdateSubscriber("relayer")
  go cosmosIBCClient.Subscribe()
  go osmosisIBCClient.Subscribe()
  // osmosisWsClient.Subscribe("relayer", "tm.event='Tx'")

  defer func() {
    cosmosIBCClient.DisConnect()
    osmosisIBCClient.DisConnect()
    // polkachuCosmosWsClient.DisConnect()
  }()
  select {}

  //* GRPC Client
  // grpcClient, _ := grpc.Connect("localhost:11290", "osmosis")
  // defer grpcClient.DisConnect()
  // var wg sync.WaitGroup
  // wg.Add(1)
  // go func() {
  //   grpcClient.GetLatestBlock()
  //   wg.Done()
  // }()
  // wg.Wait()
}
