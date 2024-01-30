package services

import (
	"gin-base/internal/model"
	"github.com/google/uuid"
	"time"
)

type ManagerService struct {
}

// 根据指定条件获取列表
func (s *ManagerService) GetManagerList(query model.ManagerQuery, page int, pageSize int) (map[string]interface{}, error) {

	data, err := model.GetManagerList(query, page, pageSize)
	return data, err
}

// 获取单条信息
func (s *ManagerService) GetManagerInfo(query model.ManagerQuery) (*model.AuthManager, error) {
	data, err := model.GetManagerInfo(query)
	return data, err
}

func (s *ManagerService) AddManager(data *model.AuthManager) (id int, err error) {
	// 默认字段添加值
	date := time.Now().Format("2006-01-02 15:04:05")
	if data.CreateTime == "" {
		data.CreateTime = date
	}
	if data.UpdateTime == "" {
		data.UpdateTime = date
	}
	if data.LastTime == "" {
		data.LastTime = date
	}
	if data.Uuid == "" {
		// 生成uuid
		uuid := uuid.New()
		data.Uuid = uuid.String()
	}
	if data.UpIp == "" {
		data.UpIp = "127.0.0.1"
	}

	id, err = model.AddManager(data)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新管理员信息
func (s *ManagerService) UpdateManager(id int, data *model.AuthManager) (rid int, err error) {
	id, err = model.UpdateManager(id, data)
	if err != nil {
		return 0, err
	}
	return rid, nil
}

// 删除管理员
func (s *ManagerService) DeleteManager(id int) error {
	id, err := model.DeleteManager(id)
	if err != nil {
		return err
	}
	return nil
}
