package controller

// Context interface
type Context interface {
	Param(string) string
	JSON(code int, obj interface{})
}
