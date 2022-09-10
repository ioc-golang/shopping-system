// Code generated by iocli, run 'iocli gen' to re-generate

package api

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	rpc_client "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_client"
)

func init() {
	rpc_client.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceIOCRPCClient{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceIOCRPCClient_{}
		},
	})
}

type serviceIOCRPCClient_ struct {
	GetRecommendedAds_ func(userID int64, num int) ([]int64, error)
}

func (s *serviceIOCRPCClient_) GetRecommendedAds(userID int64, num int) ([]int64, error) {
	return s.GetRecommendedAds_(userID, num)
}

type ServiceIOCRPCClient interface {
	GetRecommendedAds(userID int64, num int) ([]int64, error)
}

type serviceIOCRPCClient struct {
	GetRecommendedAds func(userID int64, num int) ([]int64, error)
}
