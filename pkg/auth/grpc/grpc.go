package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/authpb"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"google.golang.org/grpc"
)

type server struct {
	staffHub domain.StaffHub
	authpb.UnimplementedAuthServiceServer
}

var _ authpb.AuthServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar, staffHub domain.StaffHub) error {
	authpb.RegisterAuthServiceServer(registrar, server{staffHub: staffHub})
	return nil
}

func (s server) IsAdmin(ctx context.Context, req *authpb.IsAdminRequest) (*authpb.IsAdminResponse, error) {
	h := NewIsAdminHandler(s.staffHub)
	return h.IsAdmin(appcontext.NewGRPC(ctx), req)
}
