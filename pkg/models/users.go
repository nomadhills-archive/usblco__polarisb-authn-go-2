package models

import (
	"github.com/google/uuid"
)

type PolarisbUser struct {
	Id           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Passwordhash string
	Fullname     string `json:"fullname"`
	CreateDate   string `json:"createDate"`
	Role         int    `json:"role"`
}
