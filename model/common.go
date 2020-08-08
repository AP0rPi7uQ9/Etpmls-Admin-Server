package model

import (
	"Etpmls-Admin-Server/library"
	"github.com/gin-gonic/gin"
	"strconv"
)


// 根据PageNO和PageSize获取分页
func CommonGetPageByQuery(c *gin.Context) (limit int, offset int) {
	limit = -1
	offset = -1

	pn := c.Query(library.Config.App.Api.Pagination.Field.PageNo)
	pageNo, err1 := strconv.Atoi(pn)

	ps := c.Query(library.Config.App.Api.Pagination.Field.PageSize)
	pageSize, err2 := strconv.Atoi(ps)

	if err1 == nil && err2 == nil && pageSize > 0 && pageNo > 0 {
		limit = pageSize
		offset = (pageNo - 1) * limit
	}

	return limit, offset
}

// Check if 1(Admin User/Admin Role) is included in ids
// 查看ids中是否包含admin
func Common_CheckIfOneIsIncludeInIds(ids []uint) bool {
	for _, v := range ids {
		if v == 1 {
			return true
		}
	}

	return false
}
