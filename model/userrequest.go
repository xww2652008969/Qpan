package model

type Updateuser struct {
	Uuid        string
	Oldpassword string
	Newpassword string
}
type UserR struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}
