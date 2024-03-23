package models

type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type Annunc struct {
	Name  string `json:"name" db:"name"`
	Body  string `json:"body" db:"body"`
	Image string `json:"image" db:"image"`
	Price int    `json:"price" db:"price"`
}
