package domain

import (
  grpc "github.com/mingi3442/ibc-monitoring/infrastructure/grpc"
  ws "github.com/mingi3442/ibc-monitoring/infrastructure/ws"
)

type IBCClient struct {
  WsClient    *ws.WsClient
  GrpcClient  *grpc.GrpcClient
  recentState int64
  NetworkName string
}

type IBCClientConfig struct {
  WsUrl       string
  GrpcUrl     string
  NetworkName string
  Query       string
  Subscriber  string
}
