package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

func main() {
	r := gin.Default()
	msg := "hello world"

	// show simple json when get request
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": msg})
	})
	r.LoadHTMLGlob("template/*")

	// change server site's variable from client site by post
	r.POST("/hello", func(ctx *gin.Context) {
		var request struct {
			Msg string `json:"msg"`
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		msg = request.Msg
		ctx.JSON(http.StatusOK, gin.H{"msg": msg})
	})
	msgForHtml := "hello world for html"

	// response a html page
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"msg": msgForHtml})
	})

	// get variable from client site by get
	// sample: 127.0.0.1:8080/user/info?name="Tom"&age="21"
	r.GET("/user/info", func(ctx *gin.Context) {
		name := ctx.Query("name")
		age := ctx.Query("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// restful API
	// sample: 127.0.0.1:8080/user/info/nagisa/25
	r.GET("/user/info/:name/:age", func(ctx *gin.Context) {
		name := ctx.Param("name")
		age := ctx.Param("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	//Resolve json obj
	// sample:127.0.0.1:8080/json  {"name":"nagisa","age":25}
	r.POST("/json", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		var m map[string]interface{}
		if err := json.Unmarshal(data, &m); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"data": m})
	})

	// redirect to another url
	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// create 404 page
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"info": "this is 404 page",
			"error": "404 not found",
		})
	})

	// router group
	userGroup := r.Group("/user")
	userGroup.GET("/add", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"info": "add user"})
	})
	userGroup.GET("/delete", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"info": "delete user"})
	})
	userGroup.GET("/update", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"info": "update user"})
	})
	r.Run("127.0.0.1:8080")
}