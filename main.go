package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "123sa"
	dbname   = "TodoGoGoGo"
)

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
	//db.AutoMigrate(&TodoItem{})
	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Job 1",
		Description: "This is my first todo",
		Status:      "done",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	jsonString := "{\"id\":1,\"title\":\"Job 1\",\"description\":\"This is my first todo\",\"status\":\"done\",\"createdAt\":\"2024-05-11T09:10:59.3062614Z\",\"updatedAt\":null}\n"

	var item2 TodoItem
	if err := json.Unmarshal([]byte(jsonString), &item2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item2)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
