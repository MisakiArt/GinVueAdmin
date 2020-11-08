// 自动生成模板Tags
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type Tags struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:char;"`
      Mid  int `json:"mid" form:"mid" gorm:"column:mid;comment:;type:int;size:10;"`
      CategoryId  int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;type:int;size:10;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:;type:int;size:10;"`
      IsDelete  int `json:"isDelete" form:"isDelete" gorm:"column:is_delete;comment:;type:int;size:10;"`
      TagFollowCount  int `json:"tagFollowCount" form:"tagFollowCount" gorm:"column:tag_follow_count;comment:;type:int;size:10;"`
      LatestTaggedDate  time.Time `json:"latestTaggedDate" form:"latestTaggedDate" gorm:"column:latest_tagged_date;comment:;type:datetime;"`
      IsDefault  *bool `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:;type:tinyint;"`
      CreateTime  time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;type:datetime;"`
      TagCount  int `json:"tagCount" form:"tagCount" gorm:"column:tag_count;comment:tag被打的次数;type:int;size:10;"`
      UpdateTime  time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;type:datetime;"`
      Level  *bool `json:"level" form:"level" gorm:"column:level;comment:菜单层级;type:tinyint;"`
      ParentTagId  int `json:"parentTagId" form:"parentTagId" gorm:"column:parent_tag_id;comment:父标签id;type:int;size:10;"`
      UserSourceId  int `json:"userSourceId" form:"userSourceId" gorm:"column:user_source_id;comment:标签渠道来源;type:int;size:10;"`
}


func (Tags) TableName() string {
  return "tags"
}
