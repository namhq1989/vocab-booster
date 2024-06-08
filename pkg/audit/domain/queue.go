package domain

import "time"

type QueueNewStaffCreatedAuditLog struct {
	ActorID   string
	StaffID   string
	SourceIp  string
	CreatedAt time.Time
}
