package gin_item

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}

		storage := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(storage)

		data, err := business.GetItemById(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}