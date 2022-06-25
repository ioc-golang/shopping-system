package product

import (
	"github.com/alibaba/ioc-golang/extension/db/gorm"
	"github.com/ioc-golang/shopping-system/pkg/model/do"
	"math/rand"
)

// +ioc:autowire=true
// +ioc:autowire:type=rpc

type Service struct {
	GORMDB gorm.GORMDBIOCInterface `normal:",dev-mysql"`
}

func (d *Service) GetRecommendProductIDs(userID int64, num int) ([]int64, error) {
	productDO := make([]do.ProductDO, 0)
	resultIDs := make([]int64, 0)
	err := d.GORMDB.Table("product").Find(&productDO).Error()
	if err != nil {
		return nil, err
	}
	rand.Seed(userID)
	for i := 0; i < num; i++ {
		idx := rand.Intn(len(productDO))
		resultIDs = append(resultIDs, productDO[idx].ID)
	}

	return resultIDs, nil
}
