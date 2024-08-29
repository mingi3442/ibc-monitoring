package grpc

import (
	"context"
	cmtService "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/mingi3442/logger"
)

func (gc GrpcClient) GetLatestBlock(networkName string) int64 {

	req := &cmtService.GetLatestBlockRequest{}

	res, err := gc.serviceClient.GetLatestBlock(context.Background(), req)
	if err != nil {
		logger.Fatal("could not get latest block: %v", err)
	}

	logger.Notice("[%s] Latest Block Height: %d", networkName, res.SdkBlock.Header.Height)
	return res.SdkBlock.Header.Height
}
