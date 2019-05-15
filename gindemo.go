package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "肉体",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

// ---------------------
// 作者：Vtamins
// 来源：CSDN
// 原文：https://blog.csdn.net/u014762921/article/details/77934205
// 版权声明：本文为博主原创文章，转载请附上博文链接！
