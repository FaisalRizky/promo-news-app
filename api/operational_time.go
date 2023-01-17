package api

import (
	db "github/promo-news-app/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createOperationalTimeRequest struct {
	OpeningTime     string `json:"opening_time" binding:"required"`
	ClosingTime     string `json:"closing_time" binding:"required"`
	OperationalDays string `json:"operational_days" binding:"required"`
	OffDays         string `json:"off_days" binding:"required"`
	IsActive        bool   `json:"is_active" binding:"required"`
}

func (server *Server) createOperationalTime(ctx *gin.Context) {
	var req createOperationalTimeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateOperationalTimeParams{
		OpeningTime:     req.OpeningTime,
		ClosingTime:     req.ClosingTime,
		OperationalDays: req.OperationalDays,
		OffDays:         req.OffDays,
		IsActive:        req.IsActive,
	}

	operationalTime, err := server.store.CreateOperationalTime(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, operationalTime)
}

type getOperationalTimeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getOperationalTime(ctx *gin.Context) {
	var req getOperationalTimeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	operationalTime, err := server.store.GetOperationalTime(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, operationalTime)
}

type listOperationalTimeRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) getListOperationalTime(ctx *gin.Context) {
	var req listOperationalTimeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListOperationalTimeParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	operationalTime, err := server.store.ListOperationalTime(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, operationalTime)
}

type toggleOperationalTimeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) toogleOperationalTime(ctx *gin.Context) {
	var req toggleOperationalTimeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	operationalTime, err := server.store.GetOperationalTime(ctx, req.ID)
	arg := db.ToogleActiveOperationalTimeParams{
		ID:       req.ID,
		IsActive: !operationalTime.IsActive,
	}
	operationalTime1, err := server.store.ToogleActiveOperationalTime(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, operationalTime1)
}

type updateOperationalTimeRequest struct {
	ID              int64  `json:"id" binding:"required,min=1"`
	OpeningTime     string `json:"opening_time" binding:"required"`
	ClosingTime     string `json:"closing_time" binding:"required"`
	OperationalDays string `json:"operational_days" binding:"required"`
	OffDays         string `json:"off_days" binding:"required"`
	IsActive        bool   `json:"is_active" binding:"required"`
}

func (server *Server) updateOperationalTime(ctx *gin.Context) {
	var req updateOperationalTimeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateOperationalTimeParams{
		ID:              req.ID,
		OpeningTime:     req.OpeningTime,
		ClosingTime:     req.ClosingTime,
		OperationalDays: req.OperationalDays,
		OffDays:         req.OffDays,
		IsActive:        req.IsActive,
	}

	operationalTime1, err := server.store.UpdateOperationalTime(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, operationalTime1)
}
