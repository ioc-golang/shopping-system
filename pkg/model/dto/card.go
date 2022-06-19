package dto

import "github.com/ioc-golang/shopping-system/pkg/model/do"

type Card struct {
	CardType   int64 // 1. product, 2. advertisement
	ADsContent string
	Product    do.ProductDO
}
