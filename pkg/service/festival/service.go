package festival

import (
	"github.com/ioc-golang/shopping-system/pkg/dal/product"
	"github.com/ioc-golang/shopping-system/pkg/model/dto"
)

// +ioc:autowire=true
// +ioc:autowire:type=rpc

type Service struct {
	ProductDAL product.DALIOCInterface `singleton:""`
}

// ListCards list cards containing product and advertisement cards
func (s *Service) ListCards(pageIndex, pageSize int) ([]dto.Card, int, error) {
	returnCards := make([]dto.Card, 0)
	products, totalPage, err := s.ProductDAL.ListProductByPage(pageIndex, pageSize)
	if err != nil {
		return nil, -1, err
	}
	for _, p := range products {
		returnCards = append(returnCards, dto.Card{
			CardType: 1,
			Product:  *p,
		})
	}
	return returnCards, totalPage, nil
}
