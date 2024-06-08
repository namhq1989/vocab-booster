package grpc

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/auditpb"
	"google.golang.org/grpc"
)

type server struct {
	auditpb.UnimplementedAuditServiceServer
}

var _ auditpb.AuditServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar) error {
	auditpb.RegisterAuditServiceServer(registrar, server{})
	return nil
}
