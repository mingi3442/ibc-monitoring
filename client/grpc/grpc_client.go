package grpc

import (
  "context"

  "github.com/cosmos/cosmos-sdk/codec"
  grpcTypes "github.com/cosmos/cosmos-sdk/types/grpc"
  slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
  "github.com/mingi3442/logger"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  "google.golang.org/grpc/metadata"
)

type GrpcClient struct {
  Conn        *grpc.ClientConn
  networkName string
  url         string
}

func Connect(url, networkName string) (*GrpcClient, error) {
  // conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()),)
  conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())))
  if err != nil {
    logger.Error("did not connect: %v", err)
    return nil, err
  }
  logger.Info("Connected to cosmos-grpc")
  return &GrpcClient{
    Conn:        conn,
    networkName: networkName,
    url:         url,
  }, nil

}

func (gc GrpcClient) DisConnect() error {
  if gc.Conn != nil {
    logger.Info("Disconnected from cosmos-grpc for network %s", gc.networkName)
    gc.Conn.Close()
  }
  logger.Error("Failed to disconnect from cosmos-grpc for network %s", gc.networkName)
  return nil
}

func (gc GrpcClient) GetLatestBlock() string {
  var header metadata.MD
  client := slashingTypes.NewQueryClient(gc.Conn)
  _, err := client.Params(context.Background(), &slashingTypes.QueryParamsRequest{}, grpc.Header(&header))
  if err != nil {
    logger.Fatal("Error: %v\n", err)
  }

  blockHeight := header.Get(grpcTypes.GRPCBlockHeightHeader)[0]

  logger.Notice("[%s] Block Height: %s", gc.networkName, blockHeight)
  return blockHeight
}
