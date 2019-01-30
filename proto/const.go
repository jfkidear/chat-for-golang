package proto

const (
	UserLogin           = "user_login"
	UserLoginRes        = "user_login_res"
	UserRegister        = "user_register"
	UserStatusNotifyRes = "user_status_notify"
	UserSendMessageCmd  = "user_send_message"
	UserRecvMessageCmd  = "user_recv_message"
        UserSendMessageP2PCmd  = "user_send_message_p2p"
        UserRecvMessageP2PCmd  = "user_recv_message_p2p"
)

const (
	UserOnline  = 1
	UserOffline = 2
)

