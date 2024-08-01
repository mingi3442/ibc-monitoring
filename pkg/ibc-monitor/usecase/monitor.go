package usecase

import "github.com/mingi3442/ibc-monitoring/pkg/ibc-monitor/domain"

type MonitorUseCase struct {
  client *domain.IBCClient
}

func NewMonitorUseCase(client *domain.IBCClient) *MonitorUseCase {
  return &MonitorUseCase{client: client}
}

func (uc *MonitorUseCase) StartMonitoring() error {

  uc.client.Monitoring()

  defer uc.client.DisConnect()

  return nil
}
