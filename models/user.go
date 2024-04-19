package models

type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	UserName    string `json:"userName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	ConfirmHash string `json:"confirmHash"`
}
