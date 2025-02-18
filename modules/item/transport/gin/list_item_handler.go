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

func ListItems(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter model.Filter

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}
		paging.Process()

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}

		storage := storage.NewSQLStore(db)
		business := biz.NewListItemBiz(storage)

		data, err := business.ListAllItem(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}