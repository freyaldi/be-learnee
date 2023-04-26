package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"

	er "git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/error"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(ctx *gin.Context) {
	request := &dto.UserRegisterRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "BAD REQUEST")
		return
	}

	err := util.Validate(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	err = h.userUsecase.Register(request)
	if err != nil {
		if errors.Is(err, er.ErrUserAlreadyExists) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse("email is already used", http.StatusBadRequest))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, util.SuccessResponse("User is created successfully", http.StatusCreated, nil))
}

func (h *Handler) Login(ctx *gin.Context) {
	request := &dto.UserLoginRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "BAD REQUEST")
		return
	}
	
	err := util.Validate(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	token, err := h.userUsecase.Login(request)
	if err != nil {
		if errors.Is(err, er.ErrIncorrectCredentials) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.UnauthorizedError())
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("user log in successfully", http.StatusOK, token))
}
