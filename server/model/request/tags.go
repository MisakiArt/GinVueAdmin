package request

import "gin-vue-admin/model"

type TagsSearch struct{
    model.Tags
    PageInfo
}