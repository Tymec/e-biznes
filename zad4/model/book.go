package model

type Book struct {
	ID      int      `json:"id" gorm:"primaryKey"`
	Title   string   `json:"title" gorm:"not null;"`
	Year    int      `json:"publication_year" gorm:"not null;"`
	ISBN    string   `json:"isbn" gorm:"unique; not null;"`
	Pages   int      `json:"pages" gorm:"not null;"`
	Genres  []Genre  `json:"genres,omitempty" gorm:"many2many:book_genres;"`
	Authors []Author `json:"authors,omitempty" gorm:"many2many:book_authors;"`
}
