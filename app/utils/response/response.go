package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"component-combine/app/global/consts"
	"component-combine/app/global/my_errors"
	"component-combine/app/utils/validator_translation"
)

func ReturnJson(context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

func ReturnJsonFromString(context gin.Context, httpCode int, jsonStr string) {
	context.Header("Content-Type", "application/json; charset=utf-8")
	context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// Success 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

// Fail 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// ErrorTokenBaseInfo token 基本的格式错误
func ErrorTokenBaseInfo(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsTokenBaseInfo, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// ErrorTokenAuthFail token 权限校验失败
func ErrorTokenAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// ErrorTokenRefreshFail token 不符合刷新条件
func ErrorTokenRefreshFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusBadRequest, my_errors.ErrorsRefreshTokenFail, "")
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// TokenErrorParam token 参数校验错误
func TokenErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusUnauthorized, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// casbin 鉴权失败，返回 405 方法不允许访问
//func ErrorCasbinAuthFail(c *gin.Context, msg interface{}) {
//    ReturnJson(c, http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, my_errors.ErrorsCasbinNoAuthorization, msg)
//    c.Abort()
//}

// ErrorParam 参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// ErrorSystem 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, data)
	c.Abort()
}

// ValidatorError 翻译表单参数验证器出现的校验错误
func ValidatorError(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		wrongParam := validator_translation.RemoveTopStruct(errs.Translate(validator_translation.Trans))
		ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	} else {
		errStr := err.Error()
		// multipart:nextpart:eof 错误表示验证器需要一些参数，但是调用者没有提交任何参数
		if strings.ReplaceAll(strings.ToLower(errStr), " ", "") == "multipart:nextpart:eof" {
			ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, gin.H{"tips": my_errors.ErrorNotAllParamsIsBlank})
		} else {
			ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, gin.H{"tips": errStr})
		}
	}
	c.Abort()
}
