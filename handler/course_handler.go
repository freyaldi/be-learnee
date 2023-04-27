package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetCourse(ctx *gin.Context) {
	ctx.GetInt("user_id")
	slug := ctx.Param("slug")
	response, err := h.courseUsecase.GetCourseBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse("course is not found", http.StatusBadRequest))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("get course's detail success", http.StatusOK, response))
}

func (h *Handler) CreateCourse(ctx *gin.Context) {

	request := &dto.CreateCourseRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "BAD REQUEST")
		return
	}

	err := util.Validate(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	isAdmin := ctx.GetBool("is_admin")

	if !isAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse("this feature can access by admin only", http.StatusUnauthorized))
			return
	}

	err = h.courseUsecase.CreateCourse(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, util.SuccessResponse("Course is created successfully", http.StatusCreated, nil))
}
