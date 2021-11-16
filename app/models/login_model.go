package models

type LoginModel struct {
	Username string
}

type UserTokens struct {
	UserID      string
	Token       string
	Exp         int64
	CreatedDate string
}
