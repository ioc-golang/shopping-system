//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package product

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	rpc_service "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_service"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &service_{}
		},
	})
	rpc_service.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "github.com/ioc-golang/shopping-system/pkg/service/product/api.ServiceIOCRPCClient",
		Factory: func() interface{} {
			return &Service{}
		},
	})
}

type service_ struct {
	GetRecommendProductIDs_ func(userID int64, num int) ([]int64, error)
}

func (s *service_) GetRecommendProductIDs(userID int64, num int) ([]int64, error) {
	return s.GetRecommendProductIDs_(userID, num)
}

type ServiceIOCInterface interface {
	GetRecommendProductIDs(userID int64, num int) ([]int64, error)
}

func GetServiceRpc() (*Service, error) {
	i, err := rpc_service.GetImpl(util.GetSDIDByStructPtr(new(Service)))
	if err != nil {
		return nil, err
	}
	impl := i.(*Service)
	return impl, nil
}

func GetServiceIOCInterfaceRpc() (ServiceIOCInterface, error) {
	i, err := rpc_service.GetImplWithProxy(util.GetSDIDByStructPtr(new(Service)))
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceIOCInterface)
	return impl, nil
}