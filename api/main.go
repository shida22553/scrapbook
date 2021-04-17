package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
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

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile("/myapp/my-scrapbook-dev-9f06fe374e05.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Printf("authHeader: %v\n", c.Request.Header)
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			c.Abort()
			return
		}
		log.Printf("Verified ID token: %v\n", token)
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(authMiddleware())

	db := getGormConnect()
	db.AutoMigrate(&User{})
	closeDb, err := db.DB()

	if err != nil {
		panic(err.Error())
	}
	defer closeDb.Close()

	r.GET("/users", func(c *gin.Context) {
		users := findAllUser()
		c.JSON(http.StatusOK, users)
	})
	r.POST("/users", func(c *gin.Context) {
		user := User{
			Name: "test",
		}
		insertUser(&user)
		c.JSON(http.StatusOK, user)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
