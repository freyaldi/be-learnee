package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Vouchers(ctx *gin.Context) {
	vouchers, err := h.voucherUsecase.GetVouchers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrorResponse(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("get vouchers' data success", http.StatusOK, vouchers))
}