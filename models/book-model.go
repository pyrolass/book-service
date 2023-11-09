package models

type Book struct {
	BookID int    `json:"id" gorm:"column:book_id;primaryKey"`
	Title  string `json:"title" gorm:"column:title"`
	Author string `json:"author" gorm:"column:author"`
	Price  int    `json:"price" gorm:"column:price"`
}
