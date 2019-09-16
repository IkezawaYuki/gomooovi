package models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions_go (uuid, email, user_id, created_at) values(?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()

	stmt.Exec(uuid, user.Email, user.Id, time.Now())

	session.Uuid = uuid
	session.Email = user.Email
	session.UserId = user.Id
	session.CreatedAt = time.Now()

	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&session.Id)

	return
}

//&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt

func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from users_go where user_id = ?", user.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions_go where uuid = ?", session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	if err != nil {
		valid = false
		fmt.Println(err)
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions_go where uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("session is none")
	}
	_, err = stmt.Exec(session.Uuid)
	return
}

func (user *User) Create() (err error) {
	statement := "insert into users_go (uuid, email, password, nickname, created_at) values (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.Exec(createUUID(), user.Email, Encrypt(user.Password), user.Nickname, time.Now())

	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.Id)
	if err != nil {
		panic(err)
	}
	user.CreatedAt = time.Now()
	fmt.Println(user)

	return
}

func (user *User) Delete() (err error) {
	statement := "Delete from users_go where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return
}

func (user *User) Update() (err error) {
	statement := "update users_go set nickname = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Nickname, user.Email)
	return
}

func (user *User) DeleteAll() (err error) {
	statement := "delete from users_go"
	_, err = Db.Exec(statement)
	return
}

func Users() (users []User, err error) {
	statement := "select id, uuid, nickname, email, password, created_at from users_go"
	rows, err := Db.Query(statement)
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, nickname, email, password, created_at from users_go where email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, nickname, email, password, created_at from users_go where uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByID(id string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, nickname, email, password, created_at from users_go where id = ?", id).
		Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}



func (session *Session) User() (user User, err error) {
	user = User{}
	Db.QueryRow("select id, uuid, nickname, email, created_at from users_go where id = ?", session.UserId).Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.CreatedAt)
	return
}

func UserDeleteAll() (err error) {
	statement := "delete from users_go"
	_, err = Db.Exec(statement)
	return
}

func SessionDeleteAll() (err error) {
	statement := "delete from sessions_go"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return
}
