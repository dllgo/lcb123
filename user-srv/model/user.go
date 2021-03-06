package model

import (
	"lcb123/pkg/orm"
	"lcb123/pkg/toolfunc"
	"time"

	"github.com/jinzhu/gorm"
)

type AdminUser struct {
	orm.CommonModel
	UserName string `gorm:"column:username;size:32;not null;" json:"username" form:"user_name"`  // 用户名
	Password string `gorm:"column:password;not null;" json:"-" form:"password"`                  // 密码
	RealName string `gorm:"column:real_name;size:32;" json:"real_name" form:"real_name"`         // 真实姓名
	Email    string `gorm:"column:email;size:64;" json:"email" form:"email"`                     // 邮箱
	Phone    string `gorm:"column:phone;type:char(20);" json:"phone" form:"phone"`               // 手机号
	Status   uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"` // 状态(1:启用  2.禁用)
	Salt     string `gorm:"Column:salt" json:"-"`
}

// 设置admin_user的表名为`admin_user`
func (AdminUser) TableName() string {
	return "admin_user"
}

func (bc *AdminUser) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("id", toolfunc.GenerateUUID())
	if err != nil {
		return err
	}
	bc.CreatedAt = orm.JsonTime(time.Now())
	bc.UpdatedAt = orm.JsonTime(time.Now())
	return nil
}

// 更新前
func (bu *AdminUser) BeforeUpdate(scope *gorm.Scope) error {
	bu.UpdatedAt = orm.JsonTime(time.Now())
	return nil
}

// 删除用户及关联数据
func (AdminUser) Delete(uids []string) error {
	tx := orm.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("id in (?)", uids).Delete(&AdminUser{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("user_id in (?)", uids).Delete(&AdminUserRoles{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 根据uid获取用户
func (AdminUser) GetByID(uid string) (AdminUser, error) {
	var user AdminUser
	err := orm.GetDB().Where("id = ?", uid).Find(&user).Error
	return user, err
}
