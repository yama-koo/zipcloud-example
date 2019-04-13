package domain

import "time"

// Zipcode struct
type Zipcode struct {
	ID                     int       `gorm:"AUTO_INCREMENT"`
	PrefectureCode         string    `json:"prefcode"`
	Zipcode                string    `json:"zipcode"`
	Prefecture             string    `json:"address1"`
	Municipalities         string    `json:"address2"`
	Section                string    `json:"address3"`
	PrefecturePhonetic     string    `json:"kana1"`
	MunicipalitiesPhonetic string    `json:"kana2"`
	SectionPhonetic        string    `json:"kana3"`
	CreatedAt              time.Time `sql:"type:datetime"`
	UpdatedAt              time.Time `sql:"type:datetime"`
}
