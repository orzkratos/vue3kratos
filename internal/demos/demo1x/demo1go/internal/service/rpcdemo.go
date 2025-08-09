package service

import (
	"context"
	pb "demo1go/api/rpcdemo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yyle88/must"
	"github.com/yyle88/syncmap"
)

type DemoItem struct {
	Code string
	Name string
	Type string
}

type RpcdemoService struct {
	pb.UnimplementedRpcdemoServer
	slog *log.Helper
	sMap *syncmap.Map[string, *DemoItem]
}

func NewRpcdemoService(logger log.Logger) *RpcdemoService {
	return &RpcdemoService{
		slog: log.NewHelper(logger),
		sMap: syncmap.NewMap[string, *DemoItem](),
	}
}

func (s *RpcdemoService) CreateRpcDemo(ctx context.Context, req *pb.CreateRpcDemoRequest) (*pb.CreateRpcDemoReply, error) {
	s.sMap.Store(req.Code, &DemoItem{
		Code: must.Nice(req.Code),
		Name: must.Nice(req.Name),
		Type: must.Nice(req.Type),
	})
	return &pb.CreateRpcDemoReply{
		Code: req.Code,
	}, nil
}
func (s *RpcdemoService) DeleteRpcDemo(ctx context.Context, req *pb.DeleteRpcDemoRequest) (*pb.DeleteRpcDemoReply, error) {
	s.sMap.Delete(must.Nice(req.Code))
	return &pb.DeleteRpcDemoReply{
		Code: req.Code,
	}, nil
}
func (s *RpcdemoService) UpdateRpcDemo(ctx context.Context, req *pb.UpdateRpcDemoRequest) (*pb.UpdateRpcDemoReply, error) {
	res, ok := s.sMap.Load(must.Nice(req.Code))
	must.True(ok)
	res.Name = must.Nice(req.Name)
	return &pb.UpdateRpcDemoReply{
		Code: req.Code,
		Name: req.Name,
	}, nil
}
func (s *RpcdemoService) SelectRpcDemo(ctx context.Context, req *pb.SelectRpcDemoRequest) (*pb.SelectRpcDemoReply, error) {
	var results = make([]*pb.SelectRpcDemoReply_Item, 0)
	s.sMap.Range(func(_ string, value *DemoItem) bool {
		if value.Type == req.Type {
			results = append(results, &pb.SelectRpcDemoReply_Item{
				Code: value.Code,
				Name: value.Name,
				Type: value.Type,
			})
		}
		return true
	})
	return &pb.SelectRpcDemoReply{
		Demos: results,
	}, nil
}
