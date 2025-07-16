package models

type Account struct {
	ClientId string `json:"accountId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountNameResponse struct {
	AccountNameResponse []Account `json:"accountNameResponse"`
}
