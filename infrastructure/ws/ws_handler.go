package ws

import (
  coreTypes "github.com/cometbft/cometbft/rpc/core/types"
  types "github.com/cometbft/cometbft/types"
  "github.com/mingi3442/ibc-monitoring/internal/utils"
  "github.com/mingi3442/logger"
)

func (wc *WsClient) WsEventHandler(txCh <-chan coreTypes.ResultEvent, networkName string, recentState *int64) {
  for {
    logger.Debug("event handler with recentState: %d in %s", *recentState, networkName)
    select {
    case event := <-txCh:
      if eventBlockData, ok := event.Data.(types.EventDataNewBlock); ok {
        *recentState = int64(eventBlockData.Block.Height)
        logger.Notice("[%s] New block created: %d", networkName, eventBlockData.Block.Height)
      } else if eventTxData, ok := event.Data.(types.EventDataTx); ok {
        // utils.SaveTransactionData(eventTxData, "transaction_actions")
        actions, found := event.Events["message.action"]
        if !found {
          logger.Notice("[%s] No exist message.action in Transaction: %X, at Block: %+v", networkName, types.Tx(eventTxData.Tx).Hash(), eventTxData.TxResult.Height)
          logger.Notice("--------------------------------------------------------------------------------")
        } else {
          logger.Notice("[%s] message.action in Transaction: %X, at Block: %+v", networkName, types.Tx(eventTxData.Tx).Hash(), eventTxData.TxResult.Height)
          utils.SaveActionData(actions, "transaction_actions", "")
          for _, action := range actions {
            logger.Notice(" - %s", action)
          }
          logger.Notice("--------------------------------------------------------------------------------")
        }
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
