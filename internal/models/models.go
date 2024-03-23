package models

type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type Annunc struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Body  string `json:"body" db:"body"`
	Image string `json:"image" db:"image"`
	Price int    `json:"price" db:"price"`
	Data  int64  `json:"date" db:"data"`
}

type AnnuncRes struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Body  string `json:"body" db:"body"`
	Image string `json:"image" db:"image"`
	Price int    `json:"price" db:"price"`
	Data  string `json:"date" db:"data"`
}

type AnnuncA struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Body   string `json:"body" db:"body"`
	Image  string `json:"image" db:"image"`
	Price  int    `json:"price" db:"price"`
	Data   int64  `json:"date" db:"data"`
	IsYour bool   `json:"your"`
}

type AnnuncARes struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Body   string `json:"body" db:"body"`
	Image  string `json:"image" db:"image"`
	Price  int    `json:"price" db:"price"`
	Data   string `json:"date" db:"data"`
	IsYour bool   `json:"your"`
}
