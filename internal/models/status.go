package models

import (
	"time"

	"github.com/samredai/sludge/internal/api"
)

// Ping SLURM to get the cluster's status and calculate the ping time
func GetSlurmStatus() api.SlurmPingResult {
	start := time.Now()
	status, _ := api.RpcPingSlurm()
	duration := time.Since(start)
	return api.SlurmPingResult{
		Status:        status,
		PingStartTime: time.Now(),
		PingDuration:  duration,
	}
}
