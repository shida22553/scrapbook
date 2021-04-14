package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UId  string
	Name string
}

func insertUser(user *User) {
	db := getGormConnect()

	// insert文
	db.Create(&user)
	closeDb, err := db.DB()

	if err != nil {
		panic(err.Error())
	}
	defer closeDb.Close()
}

func findAllUser() []User {
	db := getGormConnect()
	var users []User

	// select文
	db.Order("ID asc").Find(&users)
	closeDb, err := db.DB()

	if err != nil {
		panic(err.Error())
	}
	defer closeDb.Close()
	return users
}

func getGormConnect() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/my-scrapbook"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)

	return db
}

func main() {
	r := gin.Default()
	db := getGormConnect()
	db.AutoMigrate(&User{})
	closeDb, err := db.DB()

	if err != nil {
		panic(err.Error())
	}
	defer closeDb.Close()

	// user := User{
	// 	UId:  "EDqv7BAIG7XpGBaCU1YKCUUx3WP2",
	// 	Name: "test",
	// }
	// insertUser(&user)

	r.GET("/users", func(c *gin.Context) {
		users := findAllUser()
		c.JSON(http.StatusOK, users)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
