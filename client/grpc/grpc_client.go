package grpc

import (
  "context"

  cmtService "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
  "github.com/cosmos/cosmos-sdk/codec"

  "github.com/mingi3442/logger"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
  Conn *grpc.ClientConn
  url  string
}

func Connect(url, networkName string) (*GrpcClient, error) {
  conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())))
  if err != nil {
    logger.Error("did not connect: %v", err)
    return nil, err
  }
  logger.Info("Connected to cosmos-grpc")
  return &GrpcClient{
    Conn: conn,
    url:  url,
  }, nil

}

func (gc GrpcClient) DisConnect(networkName string) error {
  if gc.Conn != nil {
    logger.Info("Disconnected from cosmos-grpc for network %s", networkName)
    gc.Conn.Close()
  }
  logger.Error("Failed to disconnect from cosmos-grpc for network %s", networkName)
  return nil
}

func (gc GrpcClient) GetLatestBlock(networkName string) int64 {

  client := cmtService.NewServiceClient(gc.Conn)

  req := &cmtService.GetLatestBlockRequest{}

  res, err := client.GetLatestBlock(context.Background(), req)
  if err != nil {
    logger.Fatal("could not get latest block: %v", err)
  }

  logger.Notice("[%s] Latest Block Height: %d", networkName, res.SdkBlock.Header.Height)
  return res.SdkBlock.Header.Height
}
