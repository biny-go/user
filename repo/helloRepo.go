package service

import (
	"context"

	pb "user/api"

	model "user/model"
	service "user/service"
)

// var _ HelloRepo = (*HelloRepo)(nil)

func NewHelloRepo() service.IHelloRepo {
	return &HelloRepo{}
}

type HelloRepo struct {
}

func (m *HelloRepo) SaveHello(ctx context.Context, req *pb.SaveHelloRequest) (*model.HelloDevPO, error) {
	data := &model.HelloDevPO{}
	return data, nil
}

func (m *HelloRepo) DeleteHello(ctx context.Context, req *pb.DeleteHelloRequest) (bool, error) {
	return true, nil
}

func (m *HelloRepo) GetHello(ctx context.Context, req *pb.GetHelloRequest) (*model.HelloDevPO, error) {
	data := &model.HelloDevPO{}
	return data, nil
}

func (m *HelloRepo) ListHello(ctx context.Context, req *pb.ListHelloRequest) ([]*model.HelloDevPO, error) {
	return []*model.HelloDevPO{}, nil
}
