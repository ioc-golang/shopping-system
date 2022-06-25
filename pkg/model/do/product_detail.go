package do

import "encoding/json"

type ProductDetailDO struct {
	ID          int64
	Name        string
	ProductType string
	Price       float64
	PictureURI  string
}

func GetProductDetailDOFromBytes(data []byte) (*ProductDetailDO, error) {
	product := &ProductDetailDO{}
	err := json.Unmarshal(data, product)
	return product, err
}

func (p *ProductDetailDO) ToBytes() []byte {
	data, _ := json.Marshal(p)
	return data
}
