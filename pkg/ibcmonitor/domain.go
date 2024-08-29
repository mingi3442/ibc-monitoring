package ibc_monitor

import (
	grpc "github.com/mingi3442/ibc-monitoring/infrastructure/grpc"
	ws "github.com/mingi3442/ibc-monitoring/infrastructure/ws"
	"github.com/mingi3442/ibc-monitoring/internal/types"
	"github.com/mingi3442/logger"
)

type IBCClient struct {
	wsClient    *ws.WsClient
	grpcClient  *grpc.GrpcClient
	recentState int64
	networkName string
}

func IBCClientBuild(config types.ConfigParams) (*IBCClient, error) {
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
		wsClient:    wsClient,
		grpcClient:  grpcClient,
		networkName: config.NetworkName,
	}

	client.updateQuery(config.Query)
	client.updateSubscriber(config.Subscriber)

	logger.Info("Connected to %s", config.NetworkName)
	return client, nil

}

func (ibc *IBCClient) updateSubscriber(subscriber string) {
	ibc.wsClient.UpdateSubscriber(subscriber)
}

func (ibc *IBCClient) updateQuery(query string) {
	ibc.wsClient.UpdateQuery(query)
}
