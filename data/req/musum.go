package req

type MuseumReq struct {
	Name string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	WorkTime string `json:"work_time" form:"work_time"`
	Country string `json:"country" form:"country"`
	City string `json:"city" form:"city"`
	Street string `json:"street" form:"street"`
	Types []int `json:"types" form:"types"`
	Images []string `json:"images" form:"images"`
}