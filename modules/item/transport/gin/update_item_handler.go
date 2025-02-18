package gin_item

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}

		storage := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(storage)

		if err := business.UpdateItemById(ctx, id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}