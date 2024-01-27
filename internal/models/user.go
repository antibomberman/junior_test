package models

type User struct {
	ID          int64
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Age         uint8  `json:"age"`
}
type UserCreate struct {
	Name        string `validate:"required" json:"name"`
	Surname     string `validate:"required" json:"surname"`
	Patronymic  string `validate:"required" json:"patronymic"`
	Gender      string `validate:"oneof=male female prefer_not_to"`
	Nationality string `validate:"required"`
	Age         string `validate:"gte=0,lte=150" json:"age"`
}

type UserUpdate struct {
	Name        string `validate:"required" json:"name"`
	Surname     string `validate:"required" json:"surname"`
	Patronymic  string `validate:"required" json:"patronymic"`
	Gender      string `validate:"oneof=male female prefer_not_to"`
	Nationality string `validate:"required"`
	Age         string `validate:"gte=0,lte=150" json:"age"`
}
