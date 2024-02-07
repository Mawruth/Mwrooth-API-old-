package req

type PieceReq struct {
	Name string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Master_piece bool `json:"master_piece" form:"master_piece"`
	CategoryID uint `json:"type_id" form:"category_id"`
	MuseumID uint `json:"museum_id" form:"museum_id"`
}