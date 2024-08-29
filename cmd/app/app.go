package app

import (
	"sync"

	"github.com/mingi3442/ibc-monitoring/internal/utils"
	"github.com/mingi3442/ibc-monitoring/pkg/ibcmonitor"
	"github.com/mingi3442/logger"
)

func Start() {
	var wg sync.WaitGroup

	config, err := utils.ReadConfig()
	if err != nil {
		logger.Fatal("Failed to read config file")
	}

	aChainIBCClient, _ := ibc_monitor.IBCClientBuild(config.ChainA)
	bChainIBCClient, _ := ibc_monitor.IBCClientBuild(config.ChainB)

	aChainMonitorUseCase := ibc_monitor.NewMonitorUseCase(aChainIBCClient)
	bChainMonitorUseCase := ibc_monitor.NewMonitorUseCase(bChainIBCClient)

	wg.Add(2)
	go func() {
		defer wg.Done()
		aChainMonitorUseCase.StartMonitoring()

	}()

	go func() {
		defer wg.Done()
		bChainMonitorUseCase.StartMonitoring()
	}()

	wg.Wait()

}
