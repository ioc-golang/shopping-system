package vo

import "github.com/ioc-golang/shopping-system/pkg/model/dto"

type GetFestivalHomepageRequest struct {
	UserID int64 `form:"user_id"`
	Num    int   `form:"num"`
}

type GetFestivalHomepageResponse struct {
	Cards []dto.Card
}
