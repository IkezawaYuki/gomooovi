package models

import (
	"fmt"
	"time"
)

type AdminUser struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminSession struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *AdminUser) CreateAdminSession() (session AdminSession, err error) {
	statement := "insert into sessions_admin_go (uuid, email, user_id, created_at) values(?, ?, ?, ?)"
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


func AdminUserByEmail(email string) (user AdminUser, err error) {
	user = AdminUser{}
	err = Db.QueryRow("select id, uuid, nickname, email, password, created_at from users_admin_go where email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func AdminUserByUUID(uuid string) (user AdminUser, err error) {
	user = AdminUser{}
	err = Db.QueryRow("select id, uuid, nickname, email, password, created_at from users_admin_go where uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func AdminUserByID(id string) (user AdminUser, err error) {
	user = AdminUser{}
	err = Db.QueryRow("select id, uuid, nickname, email, password, created_at from users_admin_go where id = ?", id).
		Scan(&user.Id, &user.Uuid, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt)
	return
}


func AdminUsers() (users []AdminUser, err error) {
	rows, err := Db.Query("SELECT id, uuid, nickname, email, password, created_at FROM users_admin_go")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var person AdminUser //構造体Person型の変数personを定義
		err := rows.Scan(&person.Id, &person.Uuid, &person.Nickname, &person.Email, &person.Password, &person.CreatedAt)

		if err != nil {
			panic(err.Error())
		}
		users = append(users, person)
	}
	return
}

func (user *AdminUser) CreateSession() (session AdminSession, err error) {
	statement := "insert into sessions_admin_go (uuid, email, user_id, created_at) values(?, ?, ?, ?)"
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


func (user *AdminUser) AdminSession() (session AdminSession, err error) {
	session = AdminSession{}
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from users_admin_go where user_id = ?", user.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (session *Session) AdminCheck() (valid bool, err error) {
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions_admin_go where uuid = ?", session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *AdminSession) DeleteByUUID() (err error) {
	statement := "delete from sessions_admin_go where uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("session is none")
	}
	_, err = stmt.Exec(session.Uuid)
	return
}

func (user *AdminUser) Create() (err error) {
	statement := "insert into users_admin_go (uuid, email, password, nickname, created_at) values (?, ?, ?, ?, ?)"
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

func (user *AdminUser) Delete() (err error) {
	statement := "Delete from users_admin_go where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return
}

func (user *AdminUser) Update() (err error) {
	statement := "update users_admin_go set nickname = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Nickname, user.Email)
	return
}

func (user *AdminUser) DeleteAll() (err error) {
	statement := "delete from users_admin_go"
	_, err = Db.Exec(statement)
	return
}
