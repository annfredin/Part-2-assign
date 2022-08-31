package models

//UserBalanceResp ...
type UserBalanceResp struct {
	UserID            int  `json:"user_id"`
	Balance     float64  `json:"balance"`
}
