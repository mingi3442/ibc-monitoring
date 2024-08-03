package ibc_monitor

import (
  grpc "github.com/mingi3442/ibc-monitoring/infrastructure/grpc"
  ws "github.com/mingi3442/ibc-monitoring/infrastructure/ws"
  "github.com/mingi3442/logger"
)

type IBCClient struct {
  wsClient    *ws.WsClient
  grpcClient  *grpc.GrpcClient
  recentState int64
  networkName string
}

func IBCClientBuild(config IBCClientConfigParams) (*IBCClient, error) {
  wsClient, err := ws.Connect(config.wsUrl, config.networkName)
  if err != nil {
    logger.Error("Failed to connect to websocket client: %v", err)
    return nil, err
  }

  grpcClient, err := grpc.Connect(config.grpcUrl, config.networkName)
  if err != nil {
    logger.Error("Failed to connect to grpc client: %v", err)
    return nil, err
  }

  client := &IBCClient{
    wsClient:    wsClient,
    grpcClient:  grpcClient,
    networkName: config.networkName,
  }

  client.updateQuery(config.query)
  client.updateSubscriber(config.subscriber)

  logger.Info("Connected to %s", config.networkName)
  return client, nil

}

func (ibc *IBCClient) updateSubscriber(subscriber string) {
  ibc.wsClient.UpdateSubscriber(subscriber)
}

func (ibc *IBCClient) updateQuery(query string) {
  ibc.wsClient.UpdateQuery(query)
}

type IBCClientConfigParams struct {
  wsUrl       string
  grpcUrl     string
  networkName string
  query       string
  subscriber  string
}

func IBCClientConfigParamsBuild(wsUrl, grpcUrl, networkName, query, subscriber string) IBCClientConfigParams {
  return IBCClientConfigParams{
    wsUrl:       wsUrl,
    grpcUrl:     grpcUrl,
    networkName: networkName,
    query:       query,
    subscriber:  subscriber,
  }
}
