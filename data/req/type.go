package req

type Type struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}
