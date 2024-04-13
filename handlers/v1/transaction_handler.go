package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"lesson/storage"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	db *sql.DB
}

func NewTransactionHandler(db *sql.DB) *TransactionHandler {
	return &TransactionHandler{db: db}
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Creates a new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param user body storage.Transaction true "Transaction information"
// @Success 200 {object} storage.Transaction "Created transaction"
// @Failure 400 {object} storage.ResponseError "Invalid transaction data"
// @Failure 401 {object} storage.ResponseError "Unauthorized"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/transaction/create [post]
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction storage.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTransaction, err := storage.CreateTransaction(h.db, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTransaction)
}

// UpdateTransaction godoc
// @Summary Update an existing transaction
// @Description Updates an existing transaction in the database
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param input body storage.Transaction true "Transaction information"
// @Success 200 {object} storage.Transaction "Updated transaction"
// @Failure 400 {object} storage.ResponseError "Invalid transaction data"
// @Failure 401 {object} storage.ResponseError "Unauthorized"
// @Failure 404 {object} storage.ResponseError "Transaction not found"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/transaction/update/{id} [put]
func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	var transaction storage.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTransaction, err := storage.UpdateTransaction(h.db, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTransaction)
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Soft deletes a transaction from the database
// @Tags transactions
// @Param id path int true "Transaction ID"
// @Success 200 {string} string "Transaction deleted successfully"
// @Failure 400 {object} storage.ResponseError "Invalid transaction ID"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/transaction/delete/{id} [delete]
func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	transactionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	err = storage.DeleteTransaction(h.db, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

// GetTransaction godoc
// @Summary Get a single transaction
// @Description Retrieves a single transaction by its ID from the database
// @Tags transactions
// @Param id path int true "Transaction ID"
// @Produce json
// @Success 200 {object} storage.Transaction "Transaction details"
// @Failure 400 {object} storage.ResponseError "Invalid transaction ID"
// @Failure 404 {object} storage.ResponseError "Transaction not found"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/transaction/{id} [get]
func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	transactionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	transaction, err := storage.GetTransaction(h.db, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// GetTransactionDetailsWithCustomerAndItem godoc
// @Summary Get transaction details with customer and item information
// @Description Retrieves transaction details with customer and item information using INNER JOIN
// @Tags transactions
// @Produce json
// @Success 200 {array} storage.TransactionView "List of transactions with customer and item details"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/transaction/details [get]
func (h *TransactionHandler) GetTransactionDetailsWithCustomerAndItem(c *gin.Context) {
	transactions, err := storage.GetTransactionDetailsWithCustomerAndItem(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// GetTransactions godoc
// @Summary Get all transactions
// @Description Retrieves all transactions from the database
// @Tags transactions
// @Produce json
// @Success 200 {array} storage.Transaction "List of transactions"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/transactions [get]
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	transactions, err := storage.GetTransactions(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// FilterTransactions godoc
// @Summary Filter transactions
// @Description Filters transactions based on parameters like ID, customer name, and item name
// @Tags transactions
// @Produce json
// @Param id query int false "Transaction ID"
// @Param customer_name query string false "Customer name"
// @Param item_name query string false "Item name"
// @Success 200 {array} storage.TransactionView "List of filtered transactions"
// @Failure 400 {object} storage.ResponseError "Invalid parameters"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/transaction/filter [get]
func (h *TransactionHandler) FilterTransactions(c *gin.Context) {
	idStr := c.Query("id")
	customerName := c.Query("customer_name")
	itemName := c.Query("item_name")

	var id int
	var err error
	if idStr != "" {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
			return
		}
	}

	transactions, err := storage.FilterTransactions(h.db, id, customerName, itemName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
