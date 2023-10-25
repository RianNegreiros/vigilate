package domain

type RealTimeMonitoringUsecase interface {
	StartMonitoring(serverID int) error
}
