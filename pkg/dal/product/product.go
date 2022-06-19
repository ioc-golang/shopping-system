package product

import (
	"github.com/alibaba/ioc-golang/extension/normal/mysql"
	"github.com/ioc-golang/shopping-system/pkg/model/do"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type DAL struct {
	ProductTable mysql.ImplIOCInterface `normal:",dev-mysql,product"`
}

func (d *DAL) ListProductByPage(pageIdx, pageSize int) ([]*do.ProductDO, int, error) {
	productDO := make([]*do.ProductDO, 0)
	err := d.ProductTable.SelectWhere("*", &productDO)
	if err != nil {
		return nil, -1, err
	}
	// todo total page
	return productDO, 100, nil
}
