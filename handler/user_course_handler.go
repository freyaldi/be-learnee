package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetUserCourses(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	userCourses, err := h.userCourseUsecase.GetUserCourses(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse("no course is found", http.StatusBadRequest))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse("get user course success", http.StatusOK, userCourses))
}
