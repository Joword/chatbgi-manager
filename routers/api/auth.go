package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/Joword/chatbgi-manager/pkg/app"
	"github.com/Joword/chatbgi-manager/pkg/e"
	"github.com/Joword/chatbgi-manager/pkg/util"
	"github.com/Joword/chatbgi-manager/service/auth_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username + " | " + password)
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	checks := authService.Check()
	if checks != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func Authority(c *gin.Context) {
	appC := app.Gin{C: c}
	// valid := validation.Validation{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	// authForm := auth{Username: username, Password: password}
	// ok, _ := valid.Valid(&authForm)

	authService := auth_service.Auth{Username: username, Password: password}
	checks := authService.Check()
	if checks != nil {
		// appC.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		hash, err := util.HashedPassword(password)
		if err != nil {
			return
		}
		isVerified := util.ComparePasswords(checks[len(checks)-1].HashedPassword, hash)
		fmt.Println(checks[len(checks)-1].HashedPassword, hash, isVerified)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appC.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appC.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}
