package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/RianNegreiros/vigilate/internal/domain"
	"github.com/pusher/pusher-http-go"
)

type realTimeMonitoringUsecase struct {
	pusherClient     *pusher.Client
	remoteServerRepo domain.RemoteServerRepository
	contextTimeout   time.Duration
}

func NewRealTimeMonitoringUsecase(pusherClient *pusher.Client, remoteServerRepo domain.RemoteServerRepository, contextTimeout time.Duration) *realTimeMonitoringUsecase {
	return &realTimeMonitoringUsecase{
		pusherClient:     pusherClient,
		remoteServerRepo: remoteServerRepo,
		contextTimeout:   contextTimeout,
	}
}

func (rtm *realTimeMonitoringUsecase) StartMonitoring(serverID int) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	prevStatus := false
	server, _ := rtm.getServerInfo(serverID)

	for range ticker.C {
		isServerUp := isServerUp(server.Address)

		if isServerUp != prevStatus {
			rtm.notifyStatusChange(serverID, isServerUp)
		}

		prevStatus = isServerUp
	}

	return nil
}

func (rtm *realTimeMonitoringUsecase) getServerInfo(serverID int) (domain.RemoteServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rtm.contextTimeout)
	defer cancel()

	server, err := rtm.remoteServerRepo.GetByID(ctx, serverID)
	if err != nil {
		log.Printf("Error getting server by ID: %v\n", err)
	}
	return server, err
}

func (rtm *realTimeMonitoringUsecase) notifyStatusChange(serverID int, isServerUp bool) {
	rtm.pusherClient.Trigger("server-"+strconv.Itoa(serverID), "server-status-changed", isServerUp)
}
