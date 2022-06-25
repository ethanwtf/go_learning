package controller

import (
	"gin_vue_project/common"
	"gin_vue_project/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//验证数据
	//fmt.Println(name, password, telephone)

	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号不正确",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不正确",
		})
		return
	}
	if len(name) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "名字不能为空",
		})
		return
	}

	if isPhoneExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "加密错误",
		})
	}

	newUser := model.User{
		Name:     name,
		Phone:    telephone,
		Password: string(hasedPassword),
	}
	db.Create(&newUser)

	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

func Login(c *gin.Context) {
	//获取参数
	db := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号不正确",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不正确",
		})
		return
	}
	if len(name) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "名字不能为空",
		})
		return
	}
	var user model.User
	db.Where("phone = ?", telephone).First(&user)
	//判断手机号是否存在
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断用户密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generate error : %v", err)
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})

}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	//result := db.Limit(1).Where("phone = ?", phone).First(&user)
	//if result.RowsAffected == 0 {
	//	return false
	//} else {
	//	return true
	//}
	db.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		return false
	} else {
		return true
	}
}
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
	})
}
