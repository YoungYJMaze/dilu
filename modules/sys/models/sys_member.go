package models

import (
	"time"

	"gorm.io/gorm"
)

// SysMember
type SysMember struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"`      //主键
	DeptId    int            `json:"deptId" gorm:"type:int unsigned;uniqueIndex:idx_did_uid;comment:部门id"` //部门id
	UserId    int            `json:"userId" gorm:"type:int unsigned;uniqueIndex:idx_did_uid;comment:用户id"` //用户id
	Nickname  string         `json:"nickname" gorm:"size:128;comment:昵称"`                                  //昵称
	Name      string         `json:"name" gorm:"size:64;comment:姓名"`                                       //姓名
	Phone     string         `json:"phone" gorm:"type:varchar(11);comment:电话"`                             //电话
	DeptPath  string         `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                       //路径
	PostTag   int            `json:"postTag" gorm:"type:tinyint unsigned;comment:职位标签 1主管 2副主管 3员工"`       //职位标签 1主管 2副主管 3员工
	Status    int            `json:"status" gorm:"type:tinyint;comment:状态 1正常 "`                           //状态 1正常
	CreateBy  int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                        //创建者
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                        //更新者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                       //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                     //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                          //删除时间
}

func (SysMember) TableName() string {
	return "sys_member"
}

func NewSysMember() *SysMember {
	return &SysMember{}
}

func (e *SysMember) SetId(id int) *SysMember {
	e.Id = id
	return e
}
func (e *SysMember) SetDeptId(deptId int) *SysMember {
	e.DeptId = deptId
	return e
}
func (e *SysMember) SetUserId(userId int) *SysMember {
	e.UserId = userId
	return e
}
func (e *SysMember) SetPostTag(postTag int) *SysMember {
	e.PostTag = postTag
	return e
}
func (e *SysMember) SetStatus(status int) *SysMember {
	e.Status = status
	return e
}
func (e *SysMember) SetCreateBy(createBy int) *SysMember {
	e.CreateBy = createBy
	return e
}
func (e *SysMember) SetUpdateBy(updateBy int) *SysMember {
	e.UpdateBy = updateBy
	return e
}
func (e *SysMember) SetCreatedAt(createdAt time.Time) *SysMember {
	e.CreatedAt = createdAt
	return e
}
func (e *SysMember) SetUpdatedAt(updatedAt time.Time) *SysMember {
	e.UpdatedAt = updatedAt
	return e
}