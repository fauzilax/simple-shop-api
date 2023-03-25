package handler

import (
	"simple-shop-api/features/user"
)

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(input interface{}) *user.UserCore {
	res := user.UserCore{}
	switch input.(type) {
	case RegisterRequest:
		cnv := input.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Password = cnv.Password
	default:
		return nil
	}
	return &res
}
