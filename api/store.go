package api

import (
	db "github/promo-news-app/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createStoreRequest struct {
	Name          string `json:"name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	Description   string `json:"description" binding:"required"`
	PhoneNumber   int64  `json:"phone_number" binding:"required"`
	OperationalID int64  `json:"operational_id" binding:"required"`
	IsActive      bool   `json:"is_active" binding:"required"`
}

func (server *Server) createStore(ctx *gin.Context) {
	var req createStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateStoresParams{
		Name:          req.Name,
		Address:       req.Address,
		Description:   req.Description,
		PhoneNumber:   req.PhoneNumber,
		OperationalID: req.OperationalID,
		IsActive:      req.IsActive,
	}

	store, err := server.store.CreateStores(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store)
}

type getStoreRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getStore(ctx *gin.Context) {
	var req getStoreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.store.GetStores(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store)
}

type listStoreRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) getListStore(ctx *gin.Context) {
	var req listStoreRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListStoresParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListStores(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type toggleStoreRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) toogleStore(ctx *gin.Context) {
	var req toggleStoreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	store, err := server.store.GetStores(ctx, req.ID)
	arg := db.ToogleActiveStoresParams{
		ID:       req.ID,
		IsActive: !store.IsActive,
	}
	store1, err := server.store.ToogleActiveStores(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store1)
}

type updateStoreRequest struct {
	ID            int64  `json:"id" binding:"required,min=1"`
	Name          string `json:"name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	Description   string `json:"description" binding:"required"`
	PhoneNumber   int64  `json:"phone_number" binding:"required"`
	OperationalID int64  `json:"operational_id" binding:"required"`
	IsActive      bool   `json:"is_active" binding:"required"`
}

func (server *Server) updateStore(ctx *gin.Context) {
	var req updateStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateStoresParams{
		ID:            req.ID,
		Name:          req.Name,
		Address:       req.Address,
		Description:   req.Description,
		PhoneNumber:   req.PhoneNumber,
		OperationalID: req.OperationalID,
		IsActive:      req.IsActive,
	}

	promo, err := server.store.UpdateStores(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, promo)
}
