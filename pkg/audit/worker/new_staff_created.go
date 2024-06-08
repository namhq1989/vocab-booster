package worker

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/domain"
)

type NewStaffCreatedHandler struct {
	auditRepository domain.AuditRepository
}

func NewNewStaffCreatedHandler(auditRepository domain.AuditRepository) NewStaffCreatedHandler {
	return NewStaffCreatedHandler{
		auditRepository: auditRepository,
	}
}

func (w NewStaffCreatedHandler) NewStaffCreated(ctx *appcontext.AppContext, payload domain.QueueNewStaffCreatedAuditLog) error {
	ctx.Logger().Text("new audit domain model")
	domainAudit, err := domain.NewAudit(domain.ActionCreate.String(), payload.ActorID, payload.StaffID, database.Collections.Staff, payload.SourceIp)
	if err != nil {
		ctx.Logger().Error("failed to create audit domain model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("create audit in db")
	if err = w.auditRepository.CreateAudit(ctx, *domainAudit); err != nil {
		ctx.Logger().Error("failed to create audit in db", err, appcontext.Fields{"audit": domainAudit})
		return err
	}

	return nil
}
