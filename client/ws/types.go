package ws

import (
  "context"

  rpcHttp "github.com/cometbft/cometbft/rpc/client/http"
)

type WsClient struct {
  RpcClient   *rpcHttp.HTTP
  url         string
  ctx         context.Context
  networkName string
  subscriber  string
  query       string
}
