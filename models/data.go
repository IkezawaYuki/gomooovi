package models

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:@/mooovi-exam3_development?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := Db.Query("SELECT id, uuid, nickname, email, password, created_at FROM users_go")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var person User //構造体Person型の変数personを定義
		err := rows.Scan(&person.Id, &person.Uuid, &person.Nickname, &person.Email, &person.Password, &person.CreatedAt)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(person.Id, person.Nickname, person.CreatedAt) //結果　1 yamada 2 suzuki

	}
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
