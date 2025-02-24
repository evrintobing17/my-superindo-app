package http

import (
	"net/http"

	"github.com/evrintobing17/my-superindo-app/internal/models"
	"github.com/evrintobing17/my-superindo-app/internal/module/cart"
	"github.com/evrintobing17/my-superindo-app/pkg/middleware"
	"github.com/evrintobing17/my-superindo-app/pkg/validation"
	"github.com/evrintobing17/my-superindo-app/utils"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase    cart.CartUsecase
	authMiddleware middleware.AuthMiddleware
}

func NewCartHandler(r *gin.Engine, cartUsecase cart.CartUsecase, authMiddleware middleware.AuthMiddleware) {
	handler := &CartHandler{
		cartUsecase:    cartUsecase,
		authMiddleware: authMiddleware,
	}
	authorized := r.Group("/cart", handler.authMiddleware.AuthMiddleware())
	{
		authorized.POST("/create", handler.addToCart)
	}
}

func (h *CartHandler) addToCart(c *gin.Context) {
	var request models.AddToCardRequest
	if err := validation.ValidateRequest(c, &request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
		return
	}

	userID := c.GetInt("userID")

	err := h.cartUsecase.AddToCart(c.Request.Context(), userID, request)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.JSONResponse(c, http.StatusBadRequest, "Bad Request", nil, err.Error())
			return
		}
		utils.JSONResponse(c, http.StatusInternalServerError, "Internal server error", nil, err.Error())
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Success", nil, nil)
}
