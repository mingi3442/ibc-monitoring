package client

import (
  grpc "github.com/mingi3442/ibc-monitoring/client/grpc"
  ws "github.com/mingi3442/ibc-monitoring/client/ws"
  "github.com/mingi3442/logger"
)

type IBCClient struct {
  WsClient    *ws.WsClient
  GrpcClient  *grpc.GrpcClient
  recentState int64
  NetworkName string
}

func Connect(wsUrl, grpcUrl, networkName string) (*IBCClient, error) {
  wsClient, err := ws.Connect(wsUrl, networkName)
  if err != nil {
    logger.Error("Failed to connect to websocket client: %v", err)
    return nil, err
  }

  grpcClient, err := grpc.Connect(grpcUrl, networkName)
  if err != nil {
    logger.Error("Failed to connect to grpc client: %v", err)
    return nil, err
  }

  client := &IBCClient{
    WsClient:    wsClient,
    GrpcClient:  grpcClient,
    NetworkName: networkName,
  }
  logger.Info("Connected to %s", networkName)
  return client, nil

}

func (ibc *IBCClient) DisConnect() error {
  if err := ibc.WsClient.DisConnect(ibc.NetworkName); err != nil {
    logger.Error("Failed to disconnect from websocket client: %v", err)
    return err
  }
  if err := ibc.GrpcClient.DisConnect(ibc.NetworkName); err != nil {
    logger.Error("Failed to disconnect from grpc client: %v", err)
    return err
  }
  logger.Info("Disconnected from %s", ibc.NetworkName)
  return nil
}

func (ibc *IBCClient) UpdateSubscriber(subscriber string) {
  ibc.WsClient.UpdateSubscriber(subscriber)
}

func (ibc *IBCClient) UpdateQuery(query string) {
  ibc.WsClient.UpdateQuery(query)
}

func (ibc *IBCClient) Subscribe() {
  func() {
    go ibc.WsClient.Subscribe(&ibc.recentState)
    // go func() {
    //   ibc.GrpcClient.GetLatestBlock(ibc.NetworkName)
    // }()
  }()
}
