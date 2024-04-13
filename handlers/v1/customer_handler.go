package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"lesson/storage"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	db *sql.DB
}

func NewCustomerHandler(db *sql.DB) *CustomerHandler {
	return &CustomerHandler{db: db}
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Creates a new customer in the database
// @Tags customers
// @Accept json
// @Produce json
// @Param input body storage.Customer true "Customer information"
// @Success 201 {object} storage.Customer "Created customer"
// @Failure 400 {object} storage.ResponseError "Invalid customer data"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/customer/create [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer storage.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdCustomer, err := storage.CreateCustomer(h.db, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdCustomer)
}

// GetCustomers godoc
// @Summary Get all customers
// @Description Retrieves all customers from the database
// @Tags customers
// @Produce json
// @Success 200 {array} storage.Customer "List of customers"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/customers [get]
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	customers, err := storage.GetCustomers(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// UpdateCustomer godoc
// @Summary Update an existing customer
// @Description Updates an existing customer in the database
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param input body storage.Customer true "Customer information"
// @Success 200 {object} storage.Customer "Updated customer"
// @Failure 400 {object} storage.ResponseError "Invalid customer data"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/customer/update/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	var customer storage.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedCustomer, err := storage.UpdateCustomer(h.db, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCustomer)
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Soft deletes a customer from the database
// @Tags customers
// @Param id path int true "Customer ID"
// @Success 200 {string} string "Customer deleted successfully"
// @Failure 400 {object} storage.ResponseError "Invalid customer ID"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/customer/delete/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	customerID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	deletedCustomer, err := storage.DeleteCustomer(h.db, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deletedCustomer)
}

// GetCustomer godoc
// @Summary Get a single customer
// @Description Retrieves a single customer by its ID from the database
// @Tags customers
// @Param id path int true "Customer ID"
// @Produce json
// @Success 200 {object} storage.Customer "Customer details"
// @Failure 400 {object} storage.ResponseError "Invalid customer ID"
// @Failure 404 {object} storage.ResponseError "Customer not found"
// @Failure 500 {object} storage.ResponseError "Internal server error"
// @Router /v1/customer/{id} [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	customerID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	customer, err := storage.GetCustomer(h.db, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}
