package model

type User struct {
	ID      int      `json:"id" gorm:"primaryKey"`
	Name    string   `json:"name" gorm:"not null;"`
	Email   string   `json:"email" gorm:"unique; not null;"`
	Reviews []Review `json:"reviews,omitempty" gorm:"foreignKey:UserID;"`
}
