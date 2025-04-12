package model

type Review struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	BookID     int    `json:"book_id" gorm:"not null;"`
	UserID     int    `json:"user_id" gorm:"not null;"`
	Rating     int    `json:"rating" gorm:"not null; check:rating >= 1 AND rating <= 5;"`
	ReviewText string `json:"review_text" gorm:"not null; type:text;"`
	Book       Book   `json:"book,omitempty" gorm:"foreignKey:BookID;"`
	User       User   `json:"user,omitempty" gorm:"foreignKey:UserID;"`
}
