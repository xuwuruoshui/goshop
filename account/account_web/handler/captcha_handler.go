package handler

import (
	"account/custom_error"
	"account/internal"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"time"
)

func CaptchaHandler(c *gin.Context){
	mobile,ok := c.GetQuery("mobile")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"参数错误",
		})
		return
	}
	

	fileName := "data.png"
	f, err := os.Create(fileName)
	if err != nil {
		zap.S().Error("GenCaptcha 失败")
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":errors.New(custom_error.GenCaptchaError),
		})
	}
	defer f.Close()

	var w io.WriterTo
	d := captcha.RandomDigits(captcha.DefaultLen)
	w = captcha.NewImage("", d, captcha.StdWidth, captcha.StdWidth)
	_, err = w.WriteTo(f)

	if err != nil {
		zap.S().Error("GenCaptcha 失败")
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":errors.New(custom_error.GenCaptchaError),
		})
		return
	}
	fmt.Println(d)
	captcha:=""
	for _, item := range d {
		captcha+=fmt.Sprintf("%d",item)
	}
	fmt.Println(captcha)


	b64, err := GetBase64(fileName)
	if err != nil {
		zap.S().Error("GenCaptcha base64失败")
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":errors.New(custom_error.GenCaptchaBase64Error),
		})
		return
	}
	// 设置120秒过期时间放入redis
	internal.RedisClient.Set(context.Background(),mobile,captcha,120*time.Second)
	c.JSON(http.StatusOK,gin.H{
		"captcha":b64,
	})
}

func GetBase64(fileName string)(string,error){
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "",err
	}
	str := base64.StdEncoding.EncodeToString(file)

	return str,nil
}
