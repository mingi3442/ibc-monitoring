package ws

import (
  "context"

  rpcHttp "github.com/cometbft/cometbft/rpc/client/http"
  coreTypes "github.com/cometbft/cometbft/rpc/core/types"
  types "github.com/cometbft/cometbft/types"

  "github.com/mingi3442/ibc-monitoring/utils"
  "github.com/mingi3442/logger"
)

type WsClient struct {
  RpcClient *rpcHttp.HTTP
  url       string
  ctx       context.Context
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
    RpcClient: rpcWsClient,
    url:       url,
    ctx:       context.Background(),
  }

  return client, nil
}

func (wc *WsClient) DisConnect(networkName string) error {
  if wc.RpcClient != nil {
    logger.Info("RPC client stopped for network %s", networkName)
    return wc.RpcClient.Stop()
  }
  logger.Error("Fail to RPC client stopped for network %s", networkName)
  return nil
}

func (wc *WsClient) Subscribe(subscriber, query, networkName string, recentState *int64) (<-chan coreTypes.ResultEvent, error) {

  events, err := wc.RpcClient.Subscribe(wc.ctx, subscriber, query)
  if err != nil {
    logger.Error("Failed to subscribe to events: %v", err)
    return nil, err
  }

  logger.Info("Subscribed to events with query: %s", query)
  wc.wsEventHandler(events, networkName, recentState)
  return events, nil
}

func (wc *WsClient) wsEventHandler(txCh <-chan coreTypes.ResultEvent, networkName string, recentState *int64) {
  for {
    logger.Debug("event handler with recentState: %d in %s", *recentState, networkName)
    select {
    case event := <-txCh:
      if eventBlockData, ok := event.Data.(types.EventDataNewBlock); ok {
        *recentState = int64(eventBlockData.Block.Height)
        logger.Notice("[%s] New block created: %d", networkName, eventBlockData.Block.Height)
      } else if eventTxData, ok := event.Data.(types.EventDataTx); ok {
        utils.SaveTransactionData(eventTxData, "transaction_actions")
        // actions, found := event.Events["message.action"]
        // if !found {
        //   logger.Notice("[%s] No exist message.action in Transaction: %X, at Block: %+v", networkName, types.Tx(eventTxData.Tx).Hash(), eventTxData.TxResult.Height)
        //   logger.Notice("--------------------------------------------------------------------------------")
        // } else {
        //   logger.Notice("[%s] message.action in Transaction: %X, at Block: %+v", networkName, types.Tx(eventTxData.Tx).Hash(), eventTxData.TxResult.Height)
        //   utils.SaveActionData(actions, "transaction_actions")
        //   for _, action := range actions {
        //     logger.Notice(" - %s", action)
        //   }
        //   logger.Notice("--------------------------------------------------------------------------------")
        // }
      } else {
        logger.Warn("Unknown event data type: %T", event.Data)
      }

    case <-wc.ctx.Done():
      logger.Debug("Event processing stopped due to timeout or cancellation")
      return
    }
  }
}

// * eventTxData.Result.Code 를 이용해서 tx 의 성공 여부를 확인할 수 있음
