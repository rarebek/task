package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"lesson/storage"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	db *sql.DB
}

func NewItemHandler(db *sql.DB) *ItemHandler {
	return &ItemHandler{db: db}
}

// GetItems godoc
// @Summary Get all items
// @Description Retrieves all items from the database
// @Tags items
// @Produce json
// @Success 200 {array} storage.Item "List of items"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/items [get]
func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := storage.GetItems(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// CreateItem godoc
// @Summary Create a new item
// @Description Creates a new item in the database
// @Tags items
// @Accept json
// @Produce json
// @Param input body storage.Item true "Item information"
// @Success 201 {object} storage.Item "Created item"
// @Failure 400 {object} storage.ResponseError "Invalid item data"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/item/create [post]
func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item storage.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdItem, err := storage.CreateItem(h.db, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdItem)
}

// @Summary Update an existing item
// @Description Updates an existing item in the database
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param input body storage.Item true "Item information"
// @Success 200 {object} storage.Item "Updated item"
// @Failure 400 {object} storage.ResponseError "Invalid item data"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/item/update/{id} [put]
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	var item storage.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedItem, err := storage.UpdateItem(h.db, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedItem)
}

// DeleteItem godoc
// @Summary Delete an item
// @Description Soft deletes an item from the database
// @Tags items
// @Param id path int true "Item ID"
// @Success 200 {string} string "Item deleted successfully"
// @Failure 400 {object} storage.ResponseError "Invalid item ID"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/item/delete/{id} [delete]
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}
	deletedItem, err := storage.DeleteItem(h.db, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deletedItem)
}

// GetItem godoc
// @Summary Get a single item
// @Description Retrieves a single item by its ID from the database
// @Tags items
// @Param id path int true "Item ID"
// @Produce json
// @Success 200 {object} storage.Item "Item details"
// @Failure 400 {object} storage.ResponseError "Invalid item ID"
// @Failure 404 {object} storage.ResponseError "Item not found"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/item/{id} [get]
func (h *ItemHandler) GetItem(c *gin.Context) {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}
	item, err := storage.GetItem(h.db, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}
