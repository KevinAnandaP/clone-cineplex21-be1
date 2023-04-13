package entity

type TheaterList struct {
	ID			uint `json:"id" gorm:"primary_key"`
	TheaterID	uint `json:"theater_id"`
	Theater		Theater `gorm:"foreignKey:TheaterID"`
	FilmID		uint `json:"film_id"`
	Film Film `gorm:"foreignKey:FilmID"`
}