package grpcclient

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/staffpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewStaffClient(_ *appcontext.AppContext, addr string) (staffpb.StaffServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return staffpb.NewStaffServiceClient(conn), nil
}

// func NewAuthClient(ctx *appcontext.AppContext, addr string) (authpb.AuthServiceClient, error) {
// 	conn, err := grpc.DialContext(ctx.Context(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return authpb.NewAuthServiceClient(conn), nil
// }
