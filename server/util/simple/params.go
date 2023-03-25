package simple

import (
	"bbs-go/util/simple/date"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// param error
func paramError(name string) error {
	return errors.New(fmt.Sprintf("unable to find param value '%s'", name))
}

// Bind : bind request dto and auto verify parameters
func ReadForm(c *gin.Context, obj interface{}) error {
	_ = c.ShouldBindUri(obj)
	if err := c.ShouldBind(obj); err != nil {
		var tagErrorMsg []string
		for _, e := range err.(validator.ValidationErrors) {
			if _, has := ValidateErrorMessage[e.Tag()]; has {
				tagErrorMsg = append(tagErrorMsg, fmt.Sprintf(ValidateErrorMessage[e.Tag()], e.Field(), e.Value()))
			} else {
				tagErrorMsg = append(tagErrorMsg, fmt.Sprintf(ValidateErrorMessage["default"], e.Tag(), e.Field(), e.Value()))
			}
		}
		return errors.New(strings.Join(tagErrorMsg, ","))
	}

	return nil
}

//ValidateErrorMessage : customize error messages
var ValidateErrorMessage = map[string]string{
	"default":        "%s - %s is invalid(%s)",
	"customValidate": "%s can not be %s",
	"required":       "%s is required,got empty %#v",
	"pwdValidate":    "%s is not a valid password",
}

func ParamValue(ctx *gin.Context, name string) string {
	return ctx.Param(name)
}

func FormValue(ctx *gin.Context, name string) string {
	return ctx.Request.FormValue(name)
}

func QueryValue(ctx *gin.Context, name string) string {
	return ctx.Query(name)
}

func FormValueRequired(ctx *gin.Context, name string) (string, error) {
	str := FormValue(ctx, name)
	if len(str) == 0 {
		return "", errors.New("参数：" + name + "不能为空")
	}
	return str, nil
}

func FormValueDefault(ctx *gin.Context, name, def string) string {
	str := FormValue(ctx, name)
	if len(str) == 0 {
		return def
	}
	return str
}

func FormValueInt(ctx *gin.Context, name string) (int, error) {
	str := FormValue(ctx, name)
	if str == "" {
		return 0, paramError(name)
	}
	return strconv.Atoi(str)
}

func FormValueIntDefault(ctx *gin.Context, name string, def int) int {
	if v, err := FormValueInt(ctx, name); err == nil {
		return v
	}
	return def
}

func FormValueInt64(ctx *gin.Context, name string) (int64, error) {
	str := FormValue(ctx, name)
	if str == "" {
		return 0, paramError(name)
	}
	return strconv.ParseInt(str, 10, 64)
}

func FormValueInt64Default(ctx *gin.Context, name string, def int64) int64 {
	if v, err := FormValueInt64(ctx, name); err == nil {
		return v
	}
	return def
}

func ParamValueInt64(ctx *gin.Context, name string) (int64, error) {
	str := ParamValue(ctx, name)
	if str == "" {
		return 0, paramError(name)
	}
	return strconv.ParseInt(str, 10, 64)
}

func QueryValueInt64(ctx *gin.Context, name string) (int64, error) {
	str := QueryValue(ctx, name)
	if str == "" {
		return 0, paramError(name)
	}
	return strconv.ParseInt(str, 10, 64)
}

func ParamValueInt64Default(ctx *gin.Context, name string, def int64) int64 {
	if v, err := ParamValueInt64(ctx, name); err == nil {
		return v
	}
	return def
}

func QueryValueInt64Default(ctx *gin.Context, name string, def int64) int64 {
	if v, err := QueryValueInt64(ctx, name); err == nil {
		return v
	}
	return def
}

func FormValueInt64Array(ctx *gin.Context, name string) []int64 {
	str := FormValue(ctx, name)
	if str == "" {
		return nil
	}
	ss := strings.Split(str, ",")
	if len(ss) == 0 {
		return nil
	}
	var ret []int64
	for _, v := range ss {
		item, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ret = append(ret, item)
	}
	return ret
}

func FormValueStringArray(ctx *gin.Context, name string) []string {
	str := FormValue(ctx, name)
	if len(str) == 0 {
		return nil
	}
	ss := strings.Split(str, ",")
	if len(ss) == 0 {
		return nil
	}
	var ret []string
	for _, s := range ss {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		ret = append(ret, s)
	}
	return ret
}

func FormValueBool(ctx *gin.Context, name string) (bool, error) {
	str := FormValue(ctx, name)
	if str == "" {
		return false, paramError(name)
	}
	return strconv.ParseBool(str)
}

// 从请求中获取日期
func FormDate(ctx *gin.Context, name string) *time.Time {
	value := FormValue(ctx, name)
	if IsBlank(value) {
		return nil
	}
	layouts := []string{date.FmtDateTime, date.FmtDate, date.FmtDateTimeNoSeconds}
	for _, layout := range layouts {
		if ret, err := date.Parse(value, layout); err == nil {
			return &ret
		}
	}
	return nil
}

func GetPaging(ctx *gin.Context) *Paging {
	page := FormValueIntDefault(ctx, "page", 1)
	limit := FormValueIntDefault(ctx, "limit", 20)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	return &Paging{Page: page, Limit: limit}
}
