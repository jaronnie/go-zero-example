// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Username string `json:username`
}

type Response struct {
	Password string `json:"password"`
}

type AllUser struct {
	Alluser []string `json:"alluser"`
}
