package common

const (
	UserStatusOnline  = 1
	UserStatusOffline = iota
)

type User struct {
	UserId    int    `json:"user_id" db:"UserId"`
	Passwd    string `json:"passwd"  db:"Passwd"`
	Nick      string `json:"nick"    db:"Nick"`  
	Sex       string `json:"sex"     db:"Sex"`
	Header    string `json:"header"  db:"Header"` 
	LastLogin string `json:"last_login" db:"LastLogin"`
	Status    int    `json:"status"  db:"Status"`
}

