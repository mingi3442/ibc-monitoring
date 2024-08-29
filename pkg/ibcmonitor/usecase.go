package ibc_monitor

import "github.com/mingi3442/logger"

type UseCase interface {
	StartMonitoring()
	DisConnect() error
}

type usecase struct {
	client *IBCClient
}

func NewMonitorUseCase(client *IBCClient) UseCase {
	return &usecase{client: client}
}

func (uc *usecase) StartMonitoring() {
	go func() {
		uc.client.wsClient.Subscribe(&uc.client.recentState)
		// go func() {
		//   ibc.GrpcClient.GetLatestBlock(ibc.NetworkName)
		// }()
	}()
	defer uc.DisConnect()
	select {}

}

func (uc *usecase) DisConnect() error {
	if err := uc.client.wsClient.DisConnect(uc.client.networkName); err != nil {
		logger.Error("Failed to disconnect from websocket client: %v", err)
		return err
	}
	if err := uc.client.grpcClient.DisConnect(uc.client.networkName); err != nil {
		logger.Error("Failed to disconnect from grpc client: %v", err)
		return err
	}
	logger.Info("Disconnected from %s", uc.client.networkName)
	return nil
}
