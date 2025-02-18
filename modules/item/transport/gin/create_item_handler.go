package gin_item

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}