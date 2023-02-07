package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type GrpcAuther struct {
}

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

func NewAuthUnaryServerInterceptor() {

}

func (a *GrpcAuther) Auth(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp interface{}, err error) {
	// 1.读取凭证,凭证放在meta信息中
	md, _ := metadata.FromIncomingContext(ctx)
	// 从meta data中获取客户端传来的凭证
	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	if err := a.ValidateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	// 继续往下层进行
	return handler(ctx, req)
}
func (a *GrpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (
	clientId, clientSecret string) {
	cidList := md[ClientHeaderKey]
	if len(cidList) > 0 {
		clientId = cidList[0]
	}

	cskList := md[ClientSecretKey]
	if len(cskList) > 0 {
		clientSecret = cskList[0]
	}

	return
}

func (a *GrpcAuther) ValidateServiceCredential(clientId, clientSecret string) error {
	if clientId == "admin" && clientSecret == "123456" {
		return nil
	} else {
		return status.Errorf(codes.Unauthenticated, "clientId or clientSecret fail")
	}
}
