package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/queue"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/domain"
)

type (
	Handlers interface {
		NewStaffCreated(ctx *appcontext.AppContext, payload domain.QueueNewStaffCreatedAuditLog) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		NewStaffCreatedHandler
	}
	Worker struct {
		queue *queue.Queue
		workerHandlers
	}
)

var _ Instance = (*Worker)(nil)

func New(queue *queue.Queue, auditRepository domain.AuditRepository) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			NewStaffCreatedHandler: NewNewStaffCreatedHandler(auditRepository),
		},
	}
}

func (w Worker) Start() {
	w.queue.Server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.AuditNewStaffCreated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueNewStaffCreatedAuditLog](bgCtx, t, queue.ParsePayload[domain.QueueNewStaffCreatedAuditLog], w.NewStaffCreated)
	})
}
