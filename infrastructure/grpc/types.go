package grpc

import (
	cmtService "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	Conn          *grpc.ClientConn
	serviceClient cmtService.ServiceClient
	url           string
	networkName   string
}
