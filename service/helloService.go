package service

import (
	"context"

	pb "user/api"
	model "user/model"

	"github.com/biny-go/loong/proto/common"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/anypb"
)

// repo interface
type IHelloRepo interface {
	SaveHello(ctx context.Context, req *pb.SaveHelloRequest) (*model.HelloDevPO, error)
	DeleteHello(ctx context.Context, req *pb.DeleteHelloRequest) (bool, error)
	GetHello(ctx context.Context, req *pb.GetHelloRequest) (*model.HelloDevPO, error)
	ListHello(ctx context.Context, req *pb.ListHelloRequest) ([]*model.HelloDevPO, error)
}

type HelloService struct {
	repo IHelloRepo
}

func NewHelloService(repo IHelloRepo) *HelloService {
	return &HelloService{repo: repo}
}

func (s *HelloService) SaveHello(ctx context.Context, req *pb.SaveHelloRequest) (*common.BaseResult, error) {
	result := &common.BaseResult{Success: true, Error: "操作成功"}
	data, err := s.repo.SaveHello(ctx, req)
	if err != nil {
		result.Success = false
		result.Error = "操作失败"
		return result, err
	}
	var dataPB pb.HelloDevPO
	if err := copier.Copy(&dataPB, &data); err != nil {
		return nil, err
	}
	anyData, err := anypb.New(&dataPB)
	if err != nil {
		return nil, err
	}
	result.Data = anyData
	return result, err
}

func (s *HelloService) DeleteHello(ctx context.Context, req *pb.DeleteHelloRequest) (*common.BaseResult, error) {
	result := &common.BaseResult{Success: true}
	status, err := s.repo.DeleteHello(ctx, req)
	if err != nil {
		result.Success = false
		result.Error = "删除失败: " + err.Error()
		return result, err
	}
	if !status {
		result.Success = false
		result.Error = "删除失败"
	}
	return result, err
}
func (s *HelloService) GetHello(ctx context.Context, req *pb.GetHelloRequest) (*common.BaseResult, error) {
	result := &common.BaseResult{}
	data, err := s.repo.GetHello(ctx, req)
	if err != nil {
		result.Success = false
		result.Error = "查询失败: " + err.Error()
		return result, err
	}
	var dataPB pb.HelloDevPO
	if err := copier.Copy(&dataPB, &data); err != nil {
		return nil, err
	}
	anyData, err := anypb.New(&dataPB)
	if err != nil {
		return nil, err
	}
	result.Data = anyData
	return result, err
}

func (s *HelloService) ListHello(ctx context.Context, req *pb.ListHelloRequest) (*common.BaseResultArray, error) {
	result := &common.BaseResultArray{Success: true, Error: "操作成功"}
	dataArray, err := s.repo.ListHello(ctx, req)
	if err != nil {
		result.Success = false
		result.Error = "操作失败"
		return result, err
	}
	for _, data := range dataArray {
		var dataPB pb.HelloDevPO
		copier.Copy(&dataPB, &data)
		anyData, _ := anypb.New(&dataPB)
		result.Data = append(result.Data, anyData)
	}
	return result, err
}
