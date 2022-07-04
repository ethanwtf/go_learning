package controller

import (
	"gin_vue_project/common"
	"gin_vue_project/dto"
	"gin_vue_project/model"
	"gin_vue_project/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	var requestUser = model.User{}
	c.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号不正确")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不正确")
		return
	}
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "名字不能为空")
		return
	}

	if isPhoneExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	response.Success(c, gin.H{"token": token}, "注册成功")
}

func Login(c *gin.Context) {
	//获取参数
	db := common.GetDB()
	var requestUser = model.User{}
	c.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号不正确")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不正确")
		return
	}
	var user model.User
	db.Where("Telephone = ?", telephone).First(&user)
	//判断手机号是否存在
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断用户密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	response.Success(c, gin.H{"token": token}, "登录成功")

}

func isPhoneExist(db *gorm.DB, Telephone string) bool {
	var user model.User
	//result := db.Limit(1).Where("phone = ?", phone).First(&user)
	//if result.RowsAffected == 0 {
	//	return false
	//} else {
	//	return true
	//}
	db.Where("Telephone = ?", Telephone).First(&user)
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
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}
