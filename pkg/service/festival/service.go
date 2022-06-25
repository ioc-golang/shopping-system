package festival

import (
	"github.com/alibaba/ioc-golang/extension/db/gorm"
	"github.com/alibaba/ioc-golang/extension/state/redis"
	"github.com/ioc-golang/shopping-system/pkg/model/do"
	"github.com/ioc-golang/shopping-system/pkg/model/dto"
	adsAPI "github.com/ioc-golang/shopping-system/pkg/service/advertisement/api"
	productAPI "github.com/ioc-golang/shopping-system/pkg/service/product/api"
	"log"
	"strconv"
	"time"
)

// +ioc:autowire=true
// +ioc:autowire:type=rpc

type Service struct {
	ProductService productAPI.ServiceIOCRPCClient `rpc-client:""`
	AdsService     adsAPI.ServiceIOCRPCClient     `rpc-client:""`
	GORMDB         gorm.GORMDBIOCInterface        `normal:",dev-mysql"`
	RedisClient    redis.RedisIOCInterface        `normal:",dev-redis"`
}

// ListCachedCards tries to list cards containing product and advertisement cards from cached
func (s *Service) ListCachedCards(uid int64, num int) ([]dto.Card, error) {
	returnCards := make([]dto.Card, 0)
	recommendedProductIDs, err := s.ProductService.GetRecommendProductIDs(uid, num)
	if err != nil {
		return nil, err
	}
	for _, productID := range recommendedProductIDs {
		product, err := s.getCachedProductDetail(productID)
		if err != nil {
			continue
		}
		returnCards = append(returnCards, dto.Card{
			CardType: 1,
			Product:  *product,
		})
	}

	recommendedAdsIDs, err := s.AdsService.GetRecommendedAds(uid, num/2)
	if err != nil {
		return nil, err
	}
	for _, adsID := range recommendedAdsIDs {
		ads, err := s.getCachedAds(adsID)
		if err != nil {
			continue
		}
		returnCards = append(returnCards, dto.Card{
			CardType:   2,
			ADsContent: ads.Content,
		})
	}
	return returnCards, nil
}

// ListCards list cards containing product and advertisement cards
func (s *Service) ListCards(uid int64, num int) ([]dto.Card, error) {
	returnCards := make([]dto.Card, 0)
	recommendedProductIDs, err := s.ProductService.GetRecommendProductIDs(uid, num)
	if err != nil {
		return nil, err
	}
	for _, productID := range recommendedProductIDs {
		product, err := s.getProductDetail(productID)
		if err != nil {
			continue
		}
		returnCards = append(returnCards, dto.Card{
			CardType: 1,
			Product:  *product,
		})
	}

	recommendedAdsIDs, err := s.AdsService.GetRecommendedAds(uid, num/2)
	if err != nil {
		return nil, err
	}
	for _, adsID := range recommendedAdsIDs {
		ads, err := s.getAds(adsID)
		if err != nil {
			continue
		}
		returnCards = append(returnCards, dto.Card{
			CardType:   2,
			ADsContent: ads.Content,
		})
	}
	return returnCards, nil
}

func (s *Service) getCachedProductDetail(id int64) (*do.ProductDetailDO, error) {
	val, err := s.RedisClient.Get("product_" + strconv.Itoa(int(id))).Result()
	if err != nil {
		log.Println("get cached product detail failed, error is ", err)
		return s.getProductDetail(id)
	}
	prod, err := do.GetProductDetailDOFromBytes([]byte(val))
	if err != nil {
		log.Println("get cached product detail failed, error is ", err)
		return s.getProductDetail(id)
	}
	return prod, nil
}

func (s *Service) getCachedAds(id int64) (*do.AdvertisementDO, error) {
	val, err := s.RedisClient.Get("advertisement_" + strconv.Itoa(int(id))).Result()
	if err != nil {
		log.Println("get cached advertisement failed, error is ", err)
		return s.getAds(id)
	}
	adv, err := do.GetAdvertisementDOFromBytes([]byte(val))
	if err != nil {
		log.Println("get cached advertisement failed, error is ", err)
		return s.getAds(id)
	}
	return adv, nil
}

func (s *Service) cacheProductDetail(id int64, product *do.ProductDetailDO) error {
	return s.RedisClient.Set("product_"+strconv.Itoa(int(id)), product.ToBytes(), time.Hour).Err()
}

func (s *Service) cacheAss(id int64, ads *do.AdvertisementDO) error {
	return s.RedisClient.Set("advertisement_"+strconv.Itoa(int(id)), ads.ToBytes(), time.Hour).Err()
}

func (s *Service) getProductDetail(id int64) (*do.ProductDetailDO, error) {
	productDetail := do.ProductDetailDO{
		ID: id,
	}
	err := s.GORMDB.Table("product_detail").First(&productDetail).Error()
	if err != nil {
		return nil, err
	}
	if err := s.cacheProductDetail(id, &productDetail); err != nil {
		log.Println("cache product detail failed, error is ", err)
	}
	return &productDetail, nil
}

func (s *Service) getAds(id int64) (*do.AdvertisementDO, error) {
	ads := do.AdvertisementDO{
		ID: id,
	}
	err := s.GORMDB.Table("advertisement").First(&ads).Error()
	if err != nil {
		return nil, err
	}
	if err := s.cacheAss(id, &ads); err != nil {
		log.Println("cache advertisement failed, error is ", err)
	}
	return &ads, nil
}
