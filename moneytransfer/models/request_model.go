package models

//UserTransferRequest ...
type UserTransferRequest struct {
	FromUserID            int  `json:"from_user_id"`
	ToUserID         int  `json:"to_user_id"`
	Amount     float64  `json:"amount"`
	
}
