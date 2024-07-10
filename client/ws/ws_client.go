package ws

import (
  "context"

  rpcHttp "github.com/cometbft/cometbft/rpc/client/http"
  coreTypes "github.com/cometbft/cometbft/rpc/core/types"
  types "github.com/cometbft/cometbft/types"
  "github.com/mingi3442/logger"
)

type WsClient struct {
  RpcClient   *rpcHttp.HTTP
  networkName string
  url         string
  ctx         context.Context
}

func Connect(url, networkName string) (*WsClient, error) {
  rpcWsClient, err := rpcHttp.New(url, "/websocket")
  if err != nil {
    return nil, err
  }

  if err := rpcWsClient.Start(); err != nil {
    logger.Error("Failed to start RPC client for network %s: %v", networkName, err)
    return nil, err
  }
  logger.Info("RPC client started for network %s", networkName)

  client := &WsClient{
    RpcClient:   rpcWsClient,
    url:         url,
    networkName: networkName,
    ctx:         context.Background(),
  }

  return client, nil
}

func (wc *WsClient) DisConnect() error {
  if wc.RpcClient != nil {
    logger.Info("RPC client stopped for network %s", wc.networkName)
    return wc.RpcClient.Stop()
  }
  logger.Error("Fail to RPC client stopped for network %s", wc.networkName)
  return nil
}

func (wc *WsClient) Subscribe(subscriber, query string) (<-chan coreTypes.ResultEvent, error) {

  events, err := wc.RpcClient.Subscribe(wc.ctx, subscriber, query)
  if err != nil {
    return nil, err
  }

  logger.Info("Subscribed to events with query: %s", query)
  return events, nil
}
