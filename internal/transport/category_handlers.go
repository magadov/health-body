package transport

import (
	"healthy_body/internal/models"
	"healthy_body/internal/service"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryHandler struct {
	category service.CategoryServices
	log      *slog.Logger
}

func NewCategoryHandler(category service.CategoryServices, log *slog.Logger) *CategoryHandler {
	return &CategoryHandler{
		category: category,
		log:      log,
	}
}

func (h *CategoryHandler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/category")
	{
		group.POST("/", h.CreateCategory)
		group.GET("/", h.GetList)
		group.GET("/:id", h.GetByID)
		group.PATCH("/:id", h.UpdateCategory)
		group.DELETE("/:id", h.DeleteCategory)
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("invalid input", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat, err := h.category.CreateCategory(input)
	if err != nil {
		h.log.Error("failed to create category", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
    "error": "failed to create category",
	})
		return
	}

	c.JSON(http.StatusCreated, cat)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.log.Warn("invalid id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	cat, err := h.category.GetCategoryByID(uint(id))
	if err != nil {
		h.log.Error("category not found", "id", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, cat)
}

func (h *CategoryHandler) GetList(c *gin.Context) {
	list, err := h.category.GetCategoryList()
	if err != nil {
		h.log.Error("failed to get category list", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get categories"})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.log.Warn("invalid id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("invalid update data", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat, err := h.category.UpdateCategory(uint(id), input)
	if err != nil {
		h.log.Error("failed to update category", "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found or update failed"})
		return
	}

	c.JSON(http.StatusOK, cat)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.log.Warn("invalid id", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.category.DeleteCategory(uint(id)); err != nil {
		h.log.Error("failed to delete category", "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found or delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
