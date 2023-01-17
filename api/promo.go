package api

import (
	db "github/promo-news-app/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPromosRequest struct {
	PromoName        string `json:"promo_name" binding:"required"`
	StoreID          int64  `json:"store_id" binding:"required"`
	PromoCode        string `json:"promo_code" binding:"required"`
	PromoDescription string `json:"promo_description" binding:"required"`
	Quantity         int64  `json:"quantity" binding:"required"`
	StartAt          int64  `json:"start_at" binding:"required"`
	ExpiredAt        int64  `json:"expired_at" binding:"required"`
	IsActive         bool   `json:"is_active" binding:"required"`
	CreatedBy        int64  `json:"created_by" binding:"required"`
}

func (server *Server) createPromo(ctx *gin.Context) {
	var req createPromosRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreatePromosParams{
		PromoName:        req.PromoName,
		StoreID:          req.StoreID,
		PromoCode:        req.PromoCode,
		PromoDescription: req.PromoDescription,
		Quantity:         req.Quantity,
		StartAt:          req.StartAt,
		ExpiredAt:        req.ExpiredAt,
		IsActive:         req.IsActive,
		CreatedBy:        req.CreatedBy,
	}

	promo, err := server.store.CreatePromos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, promo)
}

type getPromoRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPromo(ctx *gin.Context) {
	var req getPromoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	promo, err := server.store.GetPromos(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, promo)
}

type listPromoRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) getListPromo(ctx *gin.Context) {
	var req listPromoRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPromosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListPromos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type togglePromoRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) tooglePromo(ctx *gin.Context) {
	var req togglePromoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	promo, err := server.store.GetPromos(ctx, req.ID)
	arg := db.ToogleActivePromosParams{
		ID:       req.ID,
		IsActive: !promo.IsActive,
	}
	account, err := server.store.ToogleActivePromos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type updatePromosRequest struct {
	ID               int64  `json:"id" binding:"required"`
	PromoName        string `json:"promo_name" binding:"required"`
	StoreID          int64  `json:"store_id" binding:"required"`
	PromoCode        string `json:"promo_code" binding:"required"`
	PromoDescription string `json:"promo_description" binding:"required"`
	Quantity         int64  `json:"quantity" binding:"required"`
	StartAt          int64  `json:"start_at" binding:"required"`
	ExpiredAt        int64  `json:"expired_at" binding:"required"`
	IsActive         bool   `json:"is_active" binding:"required"`
	CreatedBy        int64  `json:"created_by" binding:"required"`
}

func (server *Server) updatePromo(ctx *gin.Context) {
	var req updatePromosRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdatePromosParams{
		ID:               req.ID,
		PromoName:        req.PromoName,
		StoreID:          req.StoreID,
		PromoCode:        req.PromoCode,
		PromoDescription: req.PromoDescription,
		Quantity:         req.Quantity,
		StartAt:          req.StartAt,
		ExpiredAt:        req.ExpiredAt,
		IsActive:         req.IsActive,
		CreatedBy:        req.CreatedBy,
	}

	promo, err := server.store.UpdatePromos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, promo)
}
