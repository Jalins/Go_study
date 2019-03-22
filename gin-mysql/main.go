package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(193.112.89.144:3307)/test?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works\n")
	})

	Add(db, r)

	r.Run(":8080")

}


func Add (db *sql.DB, r *gin.Engine){
	r.POST("add", func(c *gin.Context) {
		first_name := c.Request.FormValue("first_name")
		last_name := c.Request.FormValue("last_name")

		rs, err := db.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", first_name, last_name)

		if err != nil {
			log.Fatal(err.Error())
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("insert person Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})

	})
}