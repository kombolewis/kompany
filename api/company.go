package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kombolewis/kompani/db/sqlc"
)

type createCompanyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Amount      int32  `json:"amount" binding:"required"`
	Registered  bool   `json:"registered" binding:"required"`
	Type        string `json:"type" binding:"required,oneof=Corporations NonProfit Cooperative Sole_Proprietorship)"`
}

func (server *Server) createCompany(ctx *gin.Context) {
	var req createCompanyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCompanyParams{
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: true},
		Amount:      req.Amount,
		Registered:  req.Registered,
		Type:        req.Type,
	}

	company, err := server.store.CreateCompany(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, company)
}

type getCompanyRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCompany(ctx *gin.Context) {
	var req getCompanyRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	company, err := server.store.GetCompany(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, company)
}

type listCompanyRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCompanies(ctx *gin.Context) {
	var req listCompanyRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListCompaniesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	companies, err := server.store.ListCompanies(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, companies)
}
