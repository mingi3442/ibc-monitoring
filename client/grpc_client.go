package client

import (
  "context"

  "github.com/cosmos/cosmos-sdk/codec"
  grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
  slashingTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
  "github.com/mingi3442/logger"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  "google.golang.org/grpc/metadata"
)

type GrpcClient struct {
  Conn        *grpc.ClientConn
  networkName string
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
  }, nil

}

func (c GrpcClient) DisConnect() {
  logger.Info("Disconnected from cosmos-grpc")
  c.Conn.Close()
}

func (c GrpcClient) GetLatestBlock() string {
  var header metadata.MD
  client := slashingTypes.NewQueryClient(c.Conn)
  _, err := client.Params(context.Background(), &slashingTypes.QueryParamsRequest{}, grpc.Header(&header))
  if err != nil {
    logger.Fatal("Error: %v\n", err)
  }

  blockHeight := header.Get(grpctypes.GRPCBlockHeightHeader)[0]

  logger.Notice("[%s] Block Height: %s", c.networkName, blockHeight)
  return blockHeight
}
