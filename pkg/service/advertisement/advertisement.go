package product

import (
	"math/rand"

	"github.com/alibaba/ioc-golang/extension/db/gorm"

	"github.com/ioc-golang/shopping-system/pkg/model/do"
)

// +ioc:autowire=true
// +ioc:autowire:type=rpc

type Service struct {
	GORMDB gorm.GORMDBIOCInterface `normal:",dev-mysql"`
}

func (d *Service) GetRecommendedAds(userID int64, num int) ([]int64, error) {
	adsDO := make([]do.AdvertisementDO, 0)
	resultIDs := make([]int64, 0)
	err := d.GORMDB.Table("advertisement").Find(&adsDO).Error()
	if err != nil {
		return nil, err
	}
	rand.Seed(userID)
	for i := 0; i < num; i++ {
		idx := rand.Intn(len(adsDO))
		resultIDs = append(resultIDs, adsDO[idx].ID)
	}
	return resultIDs, nil
}
