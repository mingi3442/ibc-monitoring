package ws

import (
	"context"

	rpcHttp "github.com/cometbft/cometbft/rpc/client/http"
	coreTypes "github.com/cometbft/cometbft/rpc/core/types"

	"github.com/mingi3442/logger"
)

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
		ctx:         context.Background(),
		networkName: networkName,
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

func (wc *WsClient) Subscribe(recentState *int64) (<-chan coreTypes.ResultEvent, error) {

	events, err := wc.RpcClient.Subscribe(wc.ctx, wc.subscriber, wc.query)
	if err != nil {
		logger.Error("Failed to subscribe to events: %v", err)
		return nil, err
	}

	logger.Info("Subscribed to events with query: %s", wc.query)
	wc.WsEventHandler(events, wc.networkName, recentState)
	return events, nil
}

func (wc *WsClient) UpdateSubscriber(subscriber string) {
	wc.subscriber = subscriber
	logger.Info("Subscriber updated to %s", subscriber)
}

func (wc *WsClient) UpdateQuery(query string) {
	wc.query = query
	logger.Info("Query updated to %s", query)
}
