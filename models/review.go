package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserId   uint    `json:"user_id" validate:"required"`
	MuseumID uint    `json:"museum_id" validate:"required"`
	Content  string  `json:"content" validate:"required"`
	Creator  *User   `json:"creator" gorm:"-"`
	Rating   float32 `json:"rating" validate:"required"`
}

func (r Review) AfterCreate(tx *gorm.DB) error {
	var newRating float64
	err := tx.Raw("select AVG(rating) from reviews where museum_id = ?", r.MuseumID).Scan(&newRating).Error
	if err != nil {
		return err
	}
	err = tx.Model(&Museum{}).Where("id = ?", r.MuseumID).Update("rating", newRating).Error
	if err != nil {
		return err
	}
	return nil
}
