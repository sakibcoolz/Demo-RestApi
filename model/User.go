package model

import (
	"database/sql"
	"log"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
}

func UserInsert(u *User, db *sql.DB) {
		tx,_:=db.Begin()
		insForm, err := tx.Prepare("INSERT INTO user(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(u.Name, u.Email)
		log.Println("INSERT: Name: " + u.Name + " | Email: " + u.Email)
		tx.Commit()
}




