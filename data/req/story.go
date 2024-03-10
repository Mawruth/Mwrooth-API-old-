package req

type StoryReq struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	MuseumID    uint   `json:"museum_id" form:"museum_id"`
}
