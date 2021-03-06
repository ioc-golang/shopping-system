//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package festival

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	rpc_service "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_service"
	"github.com/ioc-golang/shopping-system/pkg/model/do"
	"github.com/ioc-golang/shopping-system/pkg/model/dto"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &service_{}
		},
	})
	rpc_service.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "github.com/ioc-golang/shopping-system/pkg/service/festival/api.ServiceIOCRPCClient",
		Factory: func() interface{} {
			return &Service{}
		},
	})
}

type service_ struct {
	ListCachedCards_        func(uid int64, num int) ([]dto.Card, error)
	ListCards_              func(uid int64, num int) ([]dto.Card, error)
	getCachedProductDetail_ func(id int64) (*do.ProductDetailDO, error)
	getCachedAds_           func(id int64) (*do.AdvertisementDO, error)
	cacheProductDetail_     func(id int64, product *do.ProductDetailDO) error
	cacheAss_               func(id int64, ads *do.AdvertisementDO) error
	getProductDetail_       func(id int64) (*do.ProductDetailDO, error)
	getAds_                 func(id int64) (*do.AdvertisementDO, error)
}

func (s *service_) ListCachedCards(uid int64, num int) ([]dto.Card, error) {
	return s.ListCachedCards_(uid, num)
}

func (s *service_) ListCards(uid int64, num int) ([]dto.Card, error) {
	return s.ListCards_(uid, num)
}

func (s *service_) getCachedProductDetail(id int64) (*do.ProductDetailDO, error) {
	return s.getCachedProductDetail_(id)
}

func (s *service_) getCachedAds(id int64) (*do.AdvertisementDO, error) {
	return s.getCachedAds_(id)
}

func (s *service_) cacheProductDetail(id int64, product *do.ProductDetailDO) error {
	return s.cacheProductDetail_(id, product)
}

func (s *service_) cacheAss(id int64, ads *do.AdvertisementDO) error {
	return s.cacheAss_(id, ads)
}

func (s *service_) getProductDetail(id int64) (*do.ProductDetailDO, error) {
	return s.getProductDetail_(id)
}

func (s *service_) getAds(id int64) (*do.AdvertisementDO, error) {
	return s.getAds_(id)
}

type ServiceIOCInterface interface {
	ListCachedCards(uid int64, num int) ([]dto.Card, error)
	ListCards(uid int64, num int) ([]dto.Card, error)
	getCachedProductDetail(id int64) (*do.ProductDetailDO, error)
	getCachedAds(id int64) (*do.AdvertisementDO, error)
	cacheProductDetail(id int64, product *do.ProductDetailDO) error
	cacheAss(id int64, ads *do.AdvertisementDO) error
	getProductDetail(id int64) (*do.ProductDetailDO, error)
	getAds(id int64) (*do.AdvertisementDO, error)
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
