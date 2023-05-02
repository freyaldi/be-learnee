package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Favorite(ctx *gin.Context) {
	isAdmin := ctx.GetBool("is_admin")
	if isAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse("this feature can access by user only", http.StatusUnauthorized))
		return
	}

	userId := ctx.GetInt("user_id")

	request := &dto.AddFavoriteRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err := util.Validate(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.favoriteUsecase.AddFavorite(userId, request.CourseId)
	if err != nil {
		if errors.Is(err, er.ErrCourseAlreadyFavorited) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, util.SuccessResponse("Course is added into favorites successfully", http.StatusCreated, nil))
}
