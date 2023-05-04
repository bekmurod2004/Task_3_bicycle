package handler

import (
	"app/api/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePromo(c *gin.Context) {
	var createPromo models.PromoCreate

	err := c.ShouldBindJSON(&createPromo)
	if err != nil {
		h.handlerResponse(c, "create promo", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Code().Create(context.Background(), &createPromo)
	if err != nil {
		h.handlerResponse(c, "storage.promo.create", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create product", http.StatusCreated, id)
}

func (h *Handler) GetByIdPromo(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.Code", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Code().GetByID(context.Background(), &models.PromoPrimaryKey{Promo_id: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get product by id", http.StatusCreated, resp)

}

func (h *Handler) GetListPromo(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list product", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list promo", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Code().GetList(context.Background(), &models.Query{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})

	if err != nil {
		h.handlerResponse(c, "_", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list promo response", http.StatusOK, resp)
}

func (h *Handler) DeletePromo(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "_", http.StatusBadRequest, "id incorrect")
		return
	}
	rowsAffected, err := h.storages.Code().Delete(context.Background(), &models.PromoPrimaryKey{Promo_id: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.promo.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.promo.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete promo", http.StatusNoContent, nil)

}
