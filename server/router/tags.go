package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTagsRouter(Router *gin.RouterGroup) {
	TagsRouter := Router.Group("tags").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		TagsRouter.POST("createTags", v1.CreateTags)   // 新建Tags
		TagsRouter.DELETE("deleteTags", v1.DeleteTags) // 删除Tags
		TagsRouter.DELETE("deleteTagsByIds", v1.DeleteTagsByIds) // 批量删除Tags
		TagsRouter.PUT("updateTags", v1.UpdateTags)    // 更新Tags
		TagsRouter.GET("findTags", v1.FindTags)        // 根据ID获取Tags
		TagsRouter.GET("getTagsList", v1.GetTagsList)  // 获取Tags列表
	}
}
