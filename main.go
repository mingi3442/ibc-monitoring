package main

import (
  grpc "github.com/mingi3442/ibc-monitoring/client"
  "sync"
)

func main() {
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
