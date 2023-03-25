package middleware

import (
	"bbs-go/model/constants"
	"bbs-go/service"
	"bbs-go/util/simple"
	"bbs-go/util/urls"
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	config = []PathRole{
		{Pattern: "/api/admin/sys-config/**", Roles: []string{constants.RoleOwner}},
		{Pattern: "/api/admin/user/create", Roles: []string{constants.RoleOwner}},
		{Pattern: "/api/admin/user/update", Roles: []string{constants.RoleOwner}},
		{Pattern: "/api/admin/topic-node/create", Roles: []string{constants.RoleOwner}},
		{Pattern: "/api/admin/topic-node/update", Roles: []string{constants.RoleOwner}},
		{Pattern: "/api/admin/tag/create", Roles: []string{constants.RoleOwner}},
		{Pattern: "/api/admin/tag/update", Roles: []string{constants.RoleOwner}},
		{Pattern: "/api/admin/**", Roles: []string{constants.RoleOwner, constants.RoleAdmin}},
	}
	antPathMatcher = urls.NewAntPathMatcher()
)

// AdminAuth 后台权限
func AdminAuth(ctx *gin.Context) {
	roles := getPathRoles(ctx)

	// 不需要任何角色既能访问
	if len(roles) == 0 {
		return
	}

	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		notLogin(ctx)
		return
	}
	if !user.HasAnyRole(roles...) {
		noPermission(ctx)
		return
	}

	ctx.Next()
}

// getPathRoles 获取请求该路径所需的角色
func getPathRoles(ctx *gin.Context) []string {
	p := ctx.Request.URL.Path
	for _, pathRole := range config {
		if antPathMatcher.Match(pathRole.Pattern, p) {
			return pathRole.Roles
		}
	}
	return nil
}

// notLogin 未登录返回
func notLogin(ctx *gin.Context) {
	err := simple.ErrorNotLogin
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    err.Code,
		"message": err.Message,
	})
}

// noPermission 无权限返回
func noPermission(ctx *gin.Context) {
	err := simple.JsonErrorCode(2, "无权限")
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    err.ErrorCode,
		"message": err.Message,
	})

}

type PathRole struct {
	Pattern string   // path pattern
	Roles   []string // roles
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func BodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	//statusCode := c.Writer.Status()
	//if statusCode >= 400
	{
		//ok this is an request with error, let's make a record for it
		// now print body (or log in your preferred way)
		fmt.Println("Response body: " + blw.body.String())
	}
}
