package controller

import (
	"context"

	pb "user/api"
	svc "user/service"

	common "github.com/biny-go/loong/proto/common"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type HelloController struct {
	pb.UnimplementedHelloServer
	s *svc.HelloService
}

func NewHelloController(svc *svc.HelloService) *HelloController {
	return &HelloController{s: svc}
}

func (c *HelloController) Register(srv *http.Server) {
	pb.RegisterHelloHTTPServer(srv, c)
}

func (c *HelloController) SaveHello(ctx context.Context, req *pb.SaveHelloRequest) (*common.BaseResult, error) {
	result, err := c.s.SaveHello(ctx, req)
	return result, err
}
func (c *HelloController) DeleteHello(ctx context.Context, req *pb.DeleteHelloRequest) (*common.BaseResult, error) {
	result, err := c.s.DeleteHello(ctx, req)
	return result, err
}
func (c *HelloController) GetHello(ctx context.Context, req *pb.GetHelloRequest) (*common.BaseResult, error) {
	result, err := c.s.GetHello(ctx, req)
	return result, err
}
func (c *HelloController) ListHello(ctx context.Context, req *pb.ListHelloRequest) (*common.BaseResultArray, error) {
	result, err := c.s.ListHello(ctx, req)
	return result, err
}
