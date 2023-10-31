package usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/RianNegreiros/vigilate/internal/domain"
	"github.com/google/uuid"
	"github.com/pusher/pusher-http-go"
)

type remoteServerUsecase struct {
	remoteServerRepo domain.RemoteServerRepository
	contextTimeout   time.Duration
	pusherClient     *pusher.Client
}

func NewRemoteServerUsecase(u domain.RemoteServerRepository, contextTimeout time.Duration, pusherClient *pusher.Client) *remoteServerUsecase {
	return &remoteServerUsecase{
		remoteServerRepo: u,
		contextTimeout:   contextTimeout,
		pusherClient:     pusherClient,
	}
}

func (s *remoteServerUsecase) Create(ctx context.Context, req *domain.CreateRemoteServer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	exists, err := s.remoteServerRepo.Exists(ctx, req.Address)
	if err != nil {
		log.Println("Error checking if server exists: ", err)
		return
	}

	if exists {
		return domain.ErrDuplicateAddress
	}

	remoteServer := &domain.RemoteServer{
		UserID:        req.UserID,
		Name:          req.Name,
		Address:       req.Address,
		IsActive:      isServerUp(req.Address),
		LastCheckTime: time.Now(),
		NextCheckTime: time.Now().Add(time.Minute * 5),
		Notified:      false,
	}

	err = s.remoteServerRepo.Create(ctx, remoteServer)

	return
}

func (s *remoteServerUsecase) GetByUserID(ctx context.Context, userID int) (servers []domain.RemoteServer, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	servers, err = s.remoteServerRepo.GetByUserID(ctx, userID)

	return servers, err
}

func (s *remoteServerUsecase) GetByID(ctx context.Context, id int) (server domain.RemoteServer, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	server, err = s.remoteServerRepo.GetByID(ctx, id)

	return server, err
}

var monitoringStarted sync.Map

func (s *remoteServerUsecase) StartMonitoring(serverID int) error {
	_, alreadyStarted := monitoringStarted.LoadOrStore(serverID, true)
	if alreadyStarted {
		log.Printf("Monitoring already started for server %d", serverID)
		return nil
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Get the initial server status
	server, err := s.getServerInfo(serverID)
	if err != nil {
		log.Printf("Error getting server info: %v\n", err)
		return err
	}

	initialStatus := isServerUp(server.Address)

	prevStatus := initialStatus

	for range ticker.C {
		isServerUp := isServerUp(server.Address)

		if isServerUp != prevStatus {
			s.notifyStatusChange(serverID, isServerUp)
			prevStatus = isServerUp
		}
	}

	return nil
}

func (s *remoteServerUsecase) getServerInfo(serverID int) (domain.RemoteServer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	server, err := s.remoteServerRepo.GetByID(ctx, serverID)
	if err != nil {
		log.Printf("Error getting server by ID: %v\n", err)
	}
	return server, err
}

func (s *remoteServerUsecase) notifyStatusChange(serverID int, isServerUp bool) {
	eventID := uuid.New().String()
	s.pusherClient.Trigger(fmt.Sprintf("server-%d", serverID), "status-changed", map[string]interface{}{
		"id":         eventID,
		"isServerUp": isServerUp,
	})
}

func isServerUp(address string) bool {
	resp, err := http.Get(address)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
