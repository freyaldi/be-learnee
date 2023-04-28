package handler

import (
	"errors"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetCourse(ctx *gin.Context) {
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

func (h *Handler) GetCourses(ctx *gin.Context) {
	query := &dto.CourseRequestQuery{}
	err := ctx.ShouldBindQuery(query)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}
	query = util.FormatCourseQuery(query)
	courses, err := h.courseUsecase.GetCourses(query)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse("no course is found", http.StatusBadRequest))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse("get course success", http.StatusOK, courses))
}

func (h *Handler) CreateCourse(ctx *gin.Context) {
	isAdmin := ctx.GetBool("is_admin")
	if !isAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse("this feature can access by admin only", http.StatusUnauthorized))
		return
	}

	request := &dto.CreateCourseRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err := util.Validate(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.courseUsecase.CreateCourse(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, util.SuccessResponse("Course is created successfully", http.StatusCreated, nil))
}

func (h *Handler) UpdateCourse(ctx *gin.Context) {
	isAdmin := ctx.GetBool("is_admin")
	if !isAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse("this feature can access by admin only", http.StatusUnauthorized))
		return
	}

	request := &dto.UpdateCourseRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err := util.Validate(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	cid := ctx.Param("id")
	courseId, err := strconv.Atoi(cid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.courseUsecase.UpdateCourse(courseId, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Course is updated successfully", http.StatusOK, nil))
}

func (h *Handler) DeleteCourse(ctx *gin.Context) {
	isAdmin := ctx.GetBool("is_admin")
	if !isAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse("this feature can access by admin only", http.StatusUnauthorized))
		return
	}

	request := &dto.DeleteCourseRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err := util.Validate(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.courseUsecase.DeleteCourse(request.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Course is deleted successfully", http.StatusOK, nil))
}
