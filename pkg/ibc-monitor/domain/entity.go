package domain

import (
  grpc "github.com/mingi3442/ibc-monitoring/infrastructure/grpc"
  ws "github.com/mingi3442/ibc-monitoring/infrastructure/ws"
  "github.com/mingi3442/logger"
)

func IBCClientBuild(config IBCClientConfig) (*IBCClient, error) {
  wsClient, err := ws.Connect(config.WsUrl, config.NetworkName)
  if err != nil {
    logger.Error("Failed to connect to websocket client: %v", err)
    return nil, err
  }

  grpcClient, err := grpc.Connect(config.GrpcUrl, config.NetworkName)
  if err != nil {
    logger.Error("Failed to connect to grpc client: %v", err)
    return nil, err
  }

  client := &IBCClient{
    WsClient:    wsClient,
    GrpcClient:  grpcClient,
    NetworkName: config.NetworkName,
  }

  client.updateQuery(config.Query)
  client.updateSubscriber(config.Subscriber)

  logger.Info("Connected to %s", config.NetworkName)
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

func (ibc *IBCClient) updateSubscriber(subscriber string) {
  ibc.WsClient.UpdateSubscriber(subscriber)
}

func (ibc *IBCClient) updateQuery(query string) {
  ibc.WsClient.UpdateQuery(query)
}

func (ibc *IBCClient) Monitoring() {

  go func() {

    ibc.WsClient.Subscribe(&ibc.recentState)
    // go func() {
    //   ibc.GrpcClient.GetLatestBlock(ibc.NetworkName)
    // }()
  }()

  select {}

}
