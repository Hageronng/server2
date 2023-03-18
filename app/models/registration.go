package models

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Reg struct {
	Email    string
	Password string
}

func registration() {
	Server := gin.Default()

	var R Reg

	Server.POST("/registration", PostUserData(ctx, R.Email, R.Password))

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}

	db.Exec("insert into logins (email, password) values ($1, $2)", R.Email, HashPassword(R.Password))
	db.Close()
}
