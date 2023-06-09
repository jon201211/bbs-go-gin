package service

import (
	"bbs-go/dao"
	"bbs-go/model"
	"bbs-go/util/simple"
	"bbs-go/util/simple/date"
	"net"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var OperateLogService = newOperateLogService()

func newOperateLogService() *operateLogService {
	return &operateLogService{}
}

type operateLogService struct {
}

func (s *operateLogService) Get(id int64) *model.OperateLog {
	return dao.OperateLogDao.Get(simple.DB(), id)
}

func (s *operateLogService) Take(where ...interface{}) *model.OperateLog {
	return dao.OperateLogDao.Take(simple.DB(), where...)
}

func (s *operateLogService) Find(cnd *simple.SqlCnd) []model.OperateLog {
	return dao.OperateLogDao.Find(simple.DB(), cnd)
}

func (s *operateLogService) FindOne(cnd *simple.SqlCnd) *model.OperateLog {
	return dao.OperateLogDao.FindOne(simple.DB(), cnd)
}

func (s *operateLogService) FindPageByParams(params *simple.QueryParams) (list []model.OperateLog, paging *simple.Paging) {
	return dao.OperateLogDao.FindPageByParams(simple.DB(), params)
}

func (s *operateLogService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.OperateLog, paging *simple.Paging) {
	return dao.OperateLogDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *operateLogService) Count(cnd *simple.SqlCnd) int64 {
	return dao.OperateLogDao.Count(simple.DB(), cnd)
}

func (s *operateLogService) Create(t *model.OperateLog) error {
	return dao.OperateLogDao.Create(simple.DB(), t)
}

func (s *operateLogService) Update(t *model.OperateLog) error {
	return dao.OperateLogDao.Update(simple.DB(), t)
}

func (s *operateLogService) Updates(id int64, columns map[string]interface{}) error {
	return dao.OperateLogDao.Updates(simple.DB(), id, columns)
}

func (s *operateLogService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.OperateLogDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *operateLogService) Delete(id int64) {
	dao.OperateLogDao.Delete(simple.DB(), id)
}

func (s *operateLogService) AddOperateLog(userId int64, opType, dataType string, dataId int64,
	description string, r *http.Request) {

	operateLog := &model.OperateLog{
		UserId:      userId,
		OpType:      opType,
		DataType:    dataType,
		DataId:      dataId,
		Description: description,
		CreateTime:  date.NowTimestamp(),
	}
	if r != nil {
		operateLog.Ip = ClientIP(r)
		operateLog.UserAgent = r.Header.Get("User-Agent")
		operateLog.Referer = r.Header.Get("Referer")
	}
	if err := dao.OperateLogDao.Create(simple.DB(), operateLog); err != nil {
		logrus.Error(err)
	}
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
