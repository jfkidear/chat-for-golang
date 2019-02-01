package model

import (
    "fmt"
    "time"
    //"encoding/json"
    "day10/chat/common"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var (
	UserTable = "users"
)

type UserMgr struct {
    db *sqlx.DB	
}

func NewUserMgr(db *sqlx.DB) (mgr *UserMgr) {

	mgr = &UserMgr{
		db: db,
	}
	return
}

func (p *UserMgr) getUser(id int) (user []common.User, err error) {

        err = p.db.Select(&user, "select UserId, Passwd, Nick, Sex, Header, LastLogin, Status from users where UserId=?", id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
        if len(user) == 0 {
            err = ErrUserNotExist  
            fmt.Println("select empty!")
            return
        }
	/*result, err := redis.String(conn.Do("HGet", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExist
		}
		return
	}

	user = &common.User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return
	}*/
	return
}

func (p *UserMgr) Login(id int, passwd string) (user []common.User, err error) {

	/*conn := p.pool.Get()
	defer conn.Close()*/

	user, err = p.getUser(id)
        fmt.Println("user:", user)
	if err != nil {
		return
	}

	if user[0].UserId != id || user[0].Passwd != passwd {
		err = ErrInvalidPasswd
		return
	}

	user[0].Status = common.UserStatusOnline
	user[0].LastLogin = fmt.Sprintf("%v", time.Now())

	return
}

func (p *UserMgr) Register(user *common.User) (err error) {

	/*conn := p.pool.Get()
	defer conn.Close()*/

	if user == nil {
		fmt.Println("invalid user")
		err = ErrInvalidParams
		return
	}

	_, err = p.getUser(user.UserId)
	if err == nil {
		err = ErrUserExist
		return
	}

	if err != ErrUserNotExist {
		return
	}

	/*data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", UserTable, fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		return
	}*/
	r, err := p.db.Exec("insert into users(UserId, Passwd, Nick, Sex, Header, LastLogin, Status)values(?, ?, ?, ?, ?, ?, ?)", user.UserId, user.Passwd, user.Nick, user.Sex, user.Header, user.LastLogin, user.Status)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
	return
}

