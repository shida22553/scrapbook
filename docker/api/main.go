package main

import "myapp/infrastructure"

func main() {
	infrastructure.Router.Run()
}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"

// 	firebase "firebase.google.com/go"
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/api/option"
// )

// func insertCutting(cutting *Cutting) {
// 	db := getGormConnect()
// 	// insert文
// 	db.Create(&cutting)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// }

// func updateCutting(cutting *Cutting) {
// 	db := getGormConnect()
// 	db.Save(&cutting)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// }

// func deleteCutting(cutting *Cutting) {
// 	db := getGormConnect()
// 	db.Delete(&cutting)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// }

// func findAllCuttings(user User, db *gorm.DB) []Cutting {
// 	var cuttings []Cutting

// 	// select文
// 	db.Where("user_id = ?", user.ID).Find(&cuttings)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// 	return cuttings
// }

// func findCutting(user User, id string) Cutting {
// 	db := getGormConnect()
// 	var cutting Cutting

// 	// select文
// 	db.Where("user_id = ? and id = ?", user.ID, id).Find(&cutting)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// 	return cutting
// }

// func paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		switch {
// 		case pageSize > 100:
// 			pageSize = 100
// 		case pageSize <= 0:
// 			pageSize = 10
// 		}

// 		offset := (page - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }

// func authMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Firebase SDK のセットアップ
// 		opt := option.WithCredentialsFile("/myapp/my-scrapbook-dev-9f06fe374e05.json")
// 		app, err := firebase.NewApp(context.Background(), nil, opt)
// 		if err != nil {
// 			fmt.Printf("error: %v\n", err)
// 			os.Exit(1)
// 		}
// 		auth, err := app.Auth(context.Background())
// 		if err != nil {
// 			fmt.Printf("error: %v\n", err)
// 			os.Exit(1)
// 		}

// 		// クライアントから送られてきた JWT 取得
// 		authHeader := c.Request.Header.Get("Authorization")
// 		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

// 		// JWT の検証
// 		token, err := auth.VerifyIDToken(context.Background(), idToken)
// 		if err != nil {
// 			// JWT が無効なら Handler に進まず別処理
// 			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
// 			c.Abort()
// 			return
// 		}
// 		log.Printf("Verified ID token: %v\n", token)
// 		c.Set("uid", token.UID)
// 		c.Next()
// 	}
// }

// func main() {
// 	r := gin.Default()
// 	config := cors.DefaultConfig()
// 	config.AllowAllOrigins = true
// 	config.AllowHeaders = []string{"Authorization", "Content-Type"}
// 	r.Use(cors.New(config))
// 	r.Use(authMiddleware())

// 	db := getGormConnect()
// 	db.AutoMigrate(&User{})
// 	db.AutoMigrate(&Cutting{})
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()

// 	r.GET("/users", func(c *gin.Context) {
// 		users := findAllUser()
// 		c.JSON(http.StatusOK, users)
// 	})
// 	r.POST("/users", func(c *gin.Context) {
// 		user := User{
// 			Name: "test",
// 			Uid:  c.MustGet("uid").(string),
// 		}
// 		insertUser(&user)
// 		c.JSON(http.StatusOK, user)
// 	})
// 	r.GET("/cuttings", func(c *gin.Context) {
// 		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
// 		db := getGormConnect().Scopes(paginate(page, pageSize))
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		cuttings := findAllCuttings(user, db.Order("ID desc"))
// 		c.JSON(http.StatusOK, cuttings)
// 	})
// 	r.GET("/cuttings/:id", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		cutting := findCutting(user, c.Param("id"))
// 		c.JSON(http.StatusOK, cutting)
// 	})
// 	r.POST("/cuttings", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		var cutting Cutting
// 		c.BindJSON(&cutting)
// 		cutting.UserID = user.ID
// 		insertCutting(&cutting)
// 		c.JSON(http.StatusOK, cutting)
// 	})
// 	r.PUT("/cuttings/:id", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		cutting := findCutting(user, c.Param("id"))
// 		requestBody := CuttingPutRequest{}
// 		c.BindJSON(&requestBody)
// 		cutting.Note = requestBody.Note
// 		updateCutting(&cutting)
// 		c.JSON(http.StatusOK, cutting)
// 	})
// 	r.DELETE("/cuttings/:id", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		cutting := findCutting(user, c.Param("id"))
// 		deleteCutting(&cutting)
// 		c.JSON(http.StatusOK, "ok")
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }
