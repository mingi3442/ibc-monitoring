package grpc

import (
  cmtService "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
  "github.com/cosmos/cosmos-sdk/codec"

  "github.com/mingi3442/logger"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
)

func Connect(url, networkName string) (*GrpcClient, error) {
  conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())))
  if err != nil {
    logger.Error("did not connect: %v", err)
    return nil, err
  }
  logger.Info("Connected to %s", networkName)
  return &GrpcClient{
    Conn:          conn,
    url:           url,
    networkName:   networkName,
    serviceClient: cmtService.NewServiceClient(conn),
  }, nil

}

func (gc GrpcClient) DisConnect(networkName string) error {
  if gc.Conn != nil {
    logger.Info("Disconnected for %s", networkName)
    gc.Conn.Close()
  }
  logger.Error("Failed to disconnect for %s", networkName)
  return nil
}
