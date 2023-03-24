package models

type Product struct {
	Id    int64  `json:"id" gorm:"productId"`
	Name  string `json:"productName" gorm:"unique"`
	Price int64  `json:"price"`
	Total int64  `json:"total"`
}
