package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Tags
// @Summary 创建Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "创建Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tags/createTags [post]
func CreateTags(c *gin.Context) {
	var tags model.Tags
	_ = c.ShouldBindJSON(&tags)
	err := service.CreateTags(tags)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Tags
// @Summary 删除Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "删除Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tags/deleteTags [delete]
func DeleteTags(c *gin.Context) {
	var tags model.Tags
	_ = c.ShouldBindJSON(&tags)
	err := service.DeleteTags(tags)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Tags
// @Summary 批量删除Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tags/deleteTagsByIds [delete]
func DeleteTagsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteTagsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Tags
// @Summary 更新Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "更新Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tags/updateTags [put]
func UpdateTags(c *gin.Context) {
	var tags model.Tags
	_ = c.ShouldBindJSON(&tags)
	err := service.UpdateTags(&tags)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Tags
// @Summary 用id查询Tags
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Tags true "用id查询Tags"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tags/findTags [get]
func FindTags(c *gin.Context) {
	var tags model.Tags
	_ = c.ShouldBindQuery(&tags)
	err, retags := service.GetTags(tags.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"retags": retags}, c)
	}
}

// @Tags Tags
// @Summary 分页获取Tags列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TagsSearch true "分页获取Tags列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tags/getTagsList [get]
func GetTagsList(c *gin.Context) {
	var pageInfo request.TagsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTagsInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
