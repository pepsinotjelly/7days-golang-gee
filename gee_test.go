package gee

import (
	"net/http"
	"testing"
)

func TestRun(t *testing.T) {
	r := New()
	r.GET("/", func(c *Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *Context) {
		c.String(200, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *Context) {
		c.JSON(http.StatusOK, H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
