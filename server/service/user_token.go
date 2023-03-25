package service

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple/date"
	"time"

	"bbs-go/util/simple"

	"bbs-go/cache"
	"bbs-go/dao"
	"bbs-go/model"

	"github.com/gin-gonic/gin"
)

var UserTokenService = newUserTokenService()

func newUserTokenService() *userTokenService {
	return &userTokenService{}
}

type userTokenService struct {
}

func (s *userTokenService) Get(id int64) *model.UserToken {
	return dao.UserTokenDao.Get(simple.DB(), id)
}

func (s *userTokenService) Take(where ...interface{}) *model.UserToken {
	return dao.UserTokenDao.Take(simple.DB(), where...)
}

func (s *userTokenService) Find(cnd *simple.SqlCnd) []model.UserToken {
	return dao.UserTokenDao.Find(simple.DB(), cnd)
}

func (s *userTokenService) FindOne(cnd *simple.SqlCnd) *model.UserToken {
	return dao.UserTokenDao.FindOne(simple.DB(), cnd)
}

func (s *userTokenService) FindPageByParams(params *simple.QueryParams) (list []model.UserToken, paging *simple.Paging) {
	return dao.UserTokenDao.FindPageByParams(simple.DB(), params)
}

func (s *userTokenService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.UserToken, paging *simple.Paging) {
	return dao.UserTokenDao.FindPageByCnd(simple.DB(), cnd)
}

// 获取当前登录用户的id
func (s *userTokenService) GetCurrentUserId(ctx *gin.Context) int64 {
	user := s.GetCurrent(ctx)
	if user != nil {
		return user.Id
	}
	return 0
}

// 获取当前登录用户
func (s *userTokenService) GetCurrent(ctx *gin.Context) *model.User {
	token := s.GetUserToken(ctx)
	userToken := cache.UserTokenCache.Get(token)
	// 没找到授权
	if userToken == nil || userToken.Status == constants.StatusDeleted {
		return nil
	}
	// 授权过期
	if userToken.ExpiredAt <= date.NowTimestamp() {
		return nil
	}
	user := cache.UserCache.Get(userToken.UserId)
	if user == nil || user.Status != constants.StatusOk {
		return nil
	}
	return user
}

// CheckLogin 检查登录状态
func (s *userTokenService) CheckLogin(ctx *gin.Context) (*model.User, *simple.CodeError) {
	user := s.GetCurrent(ctx)
	if user == nil {
		return nil, simple.ErrorNotLogin
	}
	return user, nil
}

// 退出登录
func (s *userTokenService) Signout(ctx *gin.Context) error {
	token := s.GetUserToken(ctx)
	userToken := dao.UserTokenDao.GetByToken(simple.DB(), token)
	if userToken == nil {
		return nil
	}
	return dao.UserTokenDao.UpdateColumn(simple.DB(), userToken.Id, "status", constants.StatusDeleted)
}

// 从请求体中获取UserToken
func (s *userTokenService) GetUserToken(ctx *gin.Context) string {
	userToken := simple.FormValue(ctx, "userToken")
	if len(userToken) > 0 {
		return userToken
	}
	return ctx.GetHeader("X-User-Token")
}

// 生成
func (s *userTokenService) Generate(userId int64) (string, error) {
	token := simple.UUID()
	tokenExpireDays := SysConfigService.GetTokenExpireDays()
	expiredAt := time.Now().Add(time.Hour * 24 * time.Duration(tokenExpireDays))
	userToken := &model.UserToken{
		Token:      token,
		UserId:     userId,
		ExpiredAt:  date.Timestamp(expiredAt),
		Status:     constants.StatusOk,
		CreateTime: date.NowTimestamp(),
	}
	err := dao.UserTokenDao.Create(simple.DB(), userToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 禁用
func (s *userTokenService) Disable(token string) error {
	t := dao.UserTokenDao.GetByToken(simple.DB(), token)
	if t == nil {
		return nil
	}
	err := dao.UserTokenDao.UpdateColumn(simple.DB(), t.Id, "status", constants.StatusDeleted)
	if err != nil {
		cache.UserTokenCache.Invalidate(token)
	}
	return err
}
