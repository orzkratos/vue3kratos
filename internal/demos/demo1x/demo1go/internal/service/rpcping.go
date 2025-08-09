package service

import (
	"context"
	pb "demo1go/api/rpcping"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type RpcpingService struct {
	pb.UnimplementedRpcpingServer
	slog *log.Helper
}

func NewRpcpingService(logger log.Logger) *RpcpingService {
	return &RpcpingService{
		slog: log.NewHelper(logger),
	}
}

func (s *RpcpingService) Ping(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	s.slog.Infof("receive-ping-message: %s", req.GetValue())
	return wrapperspb.String(req.GetValue()), nil
}
