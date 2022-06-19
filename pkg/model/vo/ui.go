package vo

import "github.com/ioc-golang/shopping-system/pkg/model/dto"

type GetFestivalHomepageRequest struct {
	UserID    int64
	PageSize  int
	PageIndex int
}

type GetFestivalHomepageResponse struct {
	Cards     []dto.Card
	TotalPage int
	PageIndex int
}
