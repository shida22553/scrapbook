package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	MustGet(key string) interface{}
	BindJSON(obj interface{}) error
	DefaultQuery(key, defaultValue string) string
	Query(key string) string
	Abort()
}
