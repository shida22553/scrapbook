package infrastructure

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
	"myapp/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))
	router.Use(authMiddleware())

	userController := controllers.NewUserController(NewSqlHandler())
	looseLeafController := controllers.NewLooseLeafController(NewSqlHandler())
	binderController := controllers.NewBinderController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.PUT("/users/:id", func(c *gin.Context) { userController.Update(c) })
	router.GET("/user", func(c *gin.Context) { userController.ShowCurrentUser(c) })
	router.POST("/loose_leafs", func(c *gin.Context) { looseLeafController.Create(c) })
	router.PUT("/loose_leafs/:id", func(c *gin.Context) { looseLeafController.Update(c) })
	router.GET("/loose_leafs", func(c *gin.Context) { looseLeafController.Index(c) })
	router.GET("/loose_leafs/:id", func(c *gin.Context) { looseLeafController.Show(c) })
	router.PUT("/loose_leafs/:id/binder", func(c *gin.Context) { looseLeafController.UpdateBinderID(c) })
	router.DELETE("/loose_leafs/:id", func(c *gin.Context) { looseLeafController.Delete(c) })
	router.POST("/binders", func(c *gin.Context) { binderController.Create(c) })
	router.PUT("/binders/:id", func(c *gin.Context) { binderController.Update(c) })
	router.GET("/binders", func(c *gin.Context) { binderController.Index(c) })
	router.GET("/binders/:id", func(c *gin.Context) { binderController.Show(c) })
	router.DELETE("/binders/:id", func(c *gin.Context) { binderController.Delete(c) })

	Router = router
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile("my-scrapbook-dev-9f06fe374e05.json")
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
		c.Set("uid", token.UID)
		c.Next()
	}
}
