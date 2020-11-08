package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateTags
// @description   create a Tags
// @param     tags               model.Tags
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateTags(tags model.Tags) (err error) {
	err = global.GVA_DB.Create(&tags).Error
	return err
}

// @title    DeleteTags
// @description   delete a Tags
// @auth                     （2020/04/05  20:22）
// @param     tags               model.Tags
// @return                    error

func DeleteTags(tags model.Tags) (err error) {
	err = global.GVA_DB.Delete(tags).Error
	return err
}

// @title    DeleteTagsByIds
// @description   delete Tagss
// @auth                     （2020/04/05  20:22）
// @param     tags               model.Tags
// @return                    error

func DeleteTagsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Tags{},"id in ?",ids.Ids).Error
	return err
}

// @title    UpdateTags
// @description   update a Tags
// @param     tags          *model.Tags
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateTags(tags *model.Tags) (err error) {
	err = global.GVA_DB.Save(tags).Error
	return err
}

// @title    GetTags
// @description   get the info of a Tags
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Tags        Tags

func GetTags(id uint) (err error, tags model.Tags) {
	err = global.GVA_DB.Where("id = ?", id).First(&tags).Error
	return
}

// @title    GetTagsInfoList
// @description   get Tags list by pagination, 分页获取Tags
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetTagsInfoList(info request.TagsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Tags{})
    var tagss []model.Tags
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tagss).Error
	return err, tagss, total
}