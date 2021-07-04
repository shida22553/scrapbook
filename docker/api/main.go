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

// func insertLooseLeaf(looseLeaf *LooseLeaf) {
// 	db := getGormConnect()
// 	// insert文
// 	db.Create(&looseLeaf)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// }

// func updateLooseLeaf(looseLeaf *LooseLeaf) {
// 	db := getGormConnect()
// 	db.Save(&looseLeaf)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// }

// func deleteLooseLeaf(looseLeaf *LooseLeaf) {
// 	db := getGormConnect()
// 	db.Delete(&looseLeaf)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// }

// func findAllLooseLeafs(user User, db *gorm.DB) []LooseLeaf {
// 	var looseLeafs []LooseLeaf

// 	// select文
// 	db.Where("user_id = ?", user.ID).Find(&looseLeafs)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// 	return looseLeafs
// }

// func findLooseLeaf(user User, id string) LooseLeaf {
// 	db := getGormConnect()
// 	var looseLeaf LooseLeaf

// 	// select文
// 	db.Where("user_id = ? and id = ?", user.ID, id).Find(&looseLeaf)
// 	closeDb, err := db.DB()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer closeDb.Close()
// 	return looseLeaf
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
// 		opt := option.WithCredentialsFile("/myapp/my-scrapbinder-dev-9f06fe374e05.json")
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
// 	db.AutoMigrate(&LooseLeaf{})
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
// 	r.GET("/looseLeafs", func(c *gin.Context) {
// 		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
// 		db := getGormConnect().Scopes(paginate(page, pageSize))
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		looseLeafs := findAllLooseLeafs(user, db.Order("ID desc"))
// 		c.JSON(http.StatusOK, looseLeafs)
// 	})
// 	r.GET("/looseLeafs/:id", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		looseLeaf := findLooseLeaf(user, c.Param("id"))
// 		c.JSON(http.StatusOK, looseLeaf)
// 	})
// 	r.POST("/looseLeafs", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		var looseLeaf LooseLeaf
// 		c.BindJSON(&looseLeaf)
// 		looseLeaf.UserID = user.ID
// 		insertLooseLeaf(&looseLeaf)
// 		c.JSON(http.StatusOK, looseLeaf)
// 	})
// 	r.PUT("/looseLeafs/:id", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		looseLeaf := findLooseLeaf(user, c.Param("id"))
// 		requestBody := LooseLeafPutRequest{}
// 		c.BindJSON(&requestBody)
// 		looseLeaf.Content = requestBody.Content
// 		updateLooseLeaf(&looseLeaf)
// 		c.JSON(http.StatusOK, looseLeaf)
// 	})
// 	r.DELETE("/looseLeafs/:id", func(c *gin.Context) {
// 		uid := c.MustGet("uid").(string)
// 		user := findUser(uid)
// 		looseLeaf := findLooseLeaf(user, c.Param("id"))
// 		deleteLooseLeaf(&looseLeaf)
// 		c.JSON(http.StatusOK, "ok")
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }
