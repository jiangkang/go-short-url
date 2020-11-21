package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiangkang/go-short-url/db"
	"net/http"
	"strconv"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/shorten_url", shortenUrl)
	router.GET("/:hash", recoverUrl)
	router.Run(":80")
}

// 将短链接重定向到原始url
func recoverUrl(context *gin.Context) {
	hash := context.Param("hash")
	url, err := db.RedisDB.Get(context, hash).Result()
	if err != nil || len(url) < 1 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "URL Not Found",
		})
	}
	context.Redirect(http.StatusMovedPermanently, url)
}

// 将url转换成短链接
func shortenUrl(context *gin.Context) {
	originUrl := context.PostForm("url")
	// 过期时间，单位秒
	expiredString := context.PostForm("expired")
	if len(expiredString) < 1 {
		expiredString = strconv.Itoa(30 * 24 * 3600)
	}
	expired, err := strconv.Atoi(expiredString)
	if err != nil {
		panic(err)
	}

	id := db.GetDbCount(context)
	hash := decToB62String(id)
	// key： hash， value： urlMd5
	error := db.RedisDB.Set(context, hash, originUrl, time.Duration(expired)*time.Second).Err()
	if error != nil {
		panic(error)
	}

	context.JSON(http.StatusOK, gin.H{
		"short_url": fmt.Sprintf("http://%s/%s", context.Request.Host, hash),
	})
}
