package api

import (
	"database/sql"
	db "github/promo-news-app/db/sqlc"
	"github/promo-news-app/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.Name,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

type createUserRequest struct {
	Email             string    `json:"email" binding:"required,email"`
	Name              string    `json:"name" binding:"required"`
	Username          string    `json:"username" binding:"required"`
	Password          string    `json:"password" binding:"required,min=6"`
	PasswordChangedAt time.Time `json:"password_changed_at" binding:"required"`
	PhoneNumber       int64     `json:"phone_number" binding:"required"`
	DeviceToken       string    `json:"device_token" binding:"required"`
	Lang              string    `json:"lang" binding:"required"`
	Avatar            string    `json:"avatar" binding:"required"`
	UserLevel         string    `json:"user_level" binding:"required"`
	IsActive          bool      `json:"is_active" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	arg := db.CreateUsersParams{
		Email:             req.Email,
		Name:              req.Name,
		Username:          req.Username,
		Password:          hashedPassword,
		PasswordChangedAt: req.PasswordChangedAt,
		PhoneNumber:       req.PhoneNumber,
		DeviceToken:       req.DeviceToken,
		Lang:              req.Lang,
		Avatar:            req.Avatar,
		UserLevel:         req.UserLevel,
		IsActive:          req.IsActive,
	}

	user, err := server.store.CreateUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.store.GetUsers(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) getListUser(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type toggleUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) toogleUser(ctx *gin.Context) {
	var req toggleUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	store, err := server.store.GetUsers(ctx, req.ID)
	arg := db.ToogleActiveUsersParams{
		ID:       req.ID,
		IsActive: !store.IsActive,
	}
	store1, err := server.store.ToogleActiveUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store1)
}

type updateUserRequest struct {
	ID                int64     `json:"id" binding:"required,min=1"`
	Email             string    `json:"email" binding:"required,email"`
	Name              string    `json:"name" binding:"required"`
	Username          string    `json:"username" binding:"required"`
	Password          string    `json:"password" binding:"required,min=6"`
	PasswordChangedAt time.Time `json:"password_changed_at" binding:"required"`
	PhoneNumber       int64     `json:"phone_number" binding:"required"`
	DeviceToken       string    `json:"device_token" binding:"required"`
	Lang              string    `json:"lang" binding:"required"`
	Avatar            string    `json:"avatar" binding:"required"`
	UserLevel         string    `json:"user_level" binding:"required"`
	IsActive          bool      `json:"is_active" binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	arg := db.UpdateUsersParams{
		ID:                req.ID,
		Email:             req.Email,
		Name:              req.Name,
		Username:          req.Username,
		Password:          hashedPassword,
		PasswordChangedAt: req.PasswordChangedAt,
		PhoneNumber:       req.PhoneNumber,
		DeviceToken:       req.DeviceToken,
		Lang:              req.Lang,
		Avatar:            req.Avatar,
		UserLevel:         req.UserLevel,
		IsActive:          req.IsActive,
	}

	user, err := server.store.UpdateUsers(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.TokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
