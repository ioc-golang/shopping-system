package do

import "encoding/json"

type AdvertisementDO struct {
	ID      int64
	Content string
}

func GetAdvertisementDOFromBytes(data []byte) (*AdvertisementDO, error) {
	product := &AdvertisementDO{}
	err := json.Unmarshal(data, product)
	return product, err
}

func (p *AdvertisementDO) ToBytes() []byte {
	data, _ := json.Marshal(p)
	return data
}
