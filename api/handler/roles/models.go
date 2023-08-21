package Roles

import (
	"service-api-cff/pkg/auth/roles"
)

type RequestRoles struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResponseAllRoles struct {
	Error bool          `json:"error"`
	Data  []*roles.Role `json:"data"`
	Code  int           `json:"code"`
	Type  string        `json:"type"`
	Msg   string        `json:"msg"`
}

type ResponseRoles struct {
	Error bool        `json:"error"`
	Data  *roles.Role `json:"data"`
	Code  int         `json:"code"`
	Type  string      `json:"type"`
	Msg   string      `json:"msg"`
}
