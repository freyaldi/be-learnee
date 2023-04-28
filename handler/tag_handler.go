package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Tags(ctx *gin.Context) {
	tags, err := h.tagUsecase.GetTags()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("get tags' data success", http.StatusOK, tags))
}