package proto

import "day10/chat/common"

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User common.User `json:"user"`
}

type LoginCmdRes struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}

type UserStatusNotify struct {
	UserId int `json:"user_id"`
	Status int `json:"user_status"`
}

type UserSendMessageReq struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}

type UserRecvMessageReq struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}

type UserSendMessageP2PReq struct {
        From int    `json:"from"`
        To int    `json:"to"`
        Data   string `json:"data"`
}

type UserRecvMessageP2PReq struct {
        UserId int    `json:"user_id"`
        Data   string `json:"data"`
}
