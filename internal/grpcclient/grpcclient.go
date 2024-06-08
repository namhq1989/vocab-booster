package grpcclient

// func NewUserClient(ctx *appcontext.AppContext, addr string) (userpb.UserServiceClient, error) {
// 	conn, err := grpc.DialContext(ctx.Context(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return userpb.NewUserServiceClient(conn), nil
// }
//
// func NewAuthClient(ctx *appcontext.AppContext, addr string) (authpb.AuthServiceClient, error) {
// 	conn, err := grpc.DialContext(ctx.Context(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return authpb.NewAuthServiceClient(conn), nil
// }
