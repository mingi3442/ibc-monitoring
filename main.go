package main

import (
  // ws "github.com/mingi3442/ibc-monitoring/client/ws"
  grpc "github.com/mingi3442/ibc-monitoring/client/grpc"
  "sync"
)

func main() {
  //* Websocket Client
  // cosmosWsClient, _ := ws.Connect("http://localhost:11157", "cosmos")
  // osmosisWsClient, _ := ws.Connect("http://localhost:11257", "osmosis")
  // // polkachuCosmosWsClient, _ := ws.Connect("https://cosmos-rpc.polkachu.com", "Polkachu Cosmos")
  // go func() {
  //   go cosmosWsClient.Subscribe("relayer", "tm.event='Tx'")
  //   go osmosisWsClient.Subscribe("relayer", "tm.event='Tx'")
  //   // polkachuCosmosWsClient.Subscribe("relayer", "tm.event='Tx'")
  // }()
  // defer func() {
  //   cosmosWsClient.DisConnect()
  //   osmosisWsClient.DisConnect()
  //   // polkachuCosmosWsClient.DisConnect()
  // }()
  // select {}

  //* GRPC Client
  grpcClient, _ := grpc.Connect("localhost:11290", "osmosis")
  defer grpcClient.DisConnect()
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    grpcClient.GetLatestBlock()
    wg.Done()
  }()
  wg.Wait()
}
