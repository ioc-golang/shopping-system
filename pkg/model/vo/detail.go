package vo

type GetProductDetailRequest struct {
	ID     int64 `schema:"id"`
	UserID int64 `schema:"user_id"`
}

type GetProductDetailResponse struct {
	ID          int64
	Name        string
	ProductType string
	Price       float64
	PictureURI  string
}
