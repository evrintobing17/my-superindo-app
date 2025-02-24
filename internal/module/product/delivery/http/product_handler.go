package http

import (
	"net/http"
	"strconv"

	"github.com/evrintobing17/my-superindo-app/internal/models"
	"github.com/evrintobing17/my-superindo-app/internal/module/product"
	"github.com/evrintobing17/my-superindo-app/pkg/middleware"
	"github.com/evrintobing17/my-superindo-app/pkg/validation"
	"github.com/evrintobing17/my-superindo-app/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase product.ProductUsecase
	authMiddleware middleware.AuthMiddleware
}

func NewAuthHandler(r *gin.Engine, productUsecase product.ProductUsecase, authMiddleware middleware.AuthMiddleware) {
	handler := &ProductHandler{
		productUsecase: productUsecase,
		authMiddleware: authMiddleware,
	}
	authorized := r.Group("/product", handler.authMiddleware.AuthMiddleware())
	{
		authorized.GET("/list", handler.productList)
		authorized.GET("/list/:category_id", handler.productListByCategory)
		authorized.GET("/detail/:id", handler.productByID)
	}
}

func (h *ProductHandler) productList(c *gin.Context) {
	resp, err := h.productUsecase.GetListProduct(c.Request.Context(), nil)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
			return
		}
		utils.JSONResponse(c, http.StatusInternalServerError, "Internal server error", nil, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Success", resp, nil)
}

func (h *ProductHandler) productListByCategory(c *gin.Context) {
	var req models.GetListProductRequest
	var categoryID *int
	req.CategoryID = c.Param("category_id")
	if err := validation.ValidateRequest(c, &req); err != nil {
		return
	}

	if req.CategoryID != "" {
		category, _ := strconv.Atoi(req.CategoryID)
		categoryID = &category
	}

	resp, err := h.productUsecase.GetListProduct(c.Request.Context(), categoryID)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
			return
		}
		utils.JSONResponse(c, http.StatusInternalServerError, "Internal server error", nil, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Success", resp, nil)
}

func (h *ProductHandler) productByID(c *gin.Context) {
	var req models.GetProductByIDRequest
	var categoryID int
	req.ID = c.Param("id")
	if err := validation.ValidateRequest(c, &req); err != nil {
		return
	}

	categoryID, _ = strconv.Atoi(req.ID)

	resp, err := h.productUsecase.GetProductByID(c.Request.Context(), categoryID)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
			return
		}
		utils.JSONResponse(c, http.StatusInternalServerError, "Internal server error", nil, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Success", resp, nil)
}
