package api

import (
	"database/sql"
	v1 "lesson/handlers/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/v1")

	customerHandler := v1.NewCustomerHandler(db)
	api.GET("/customers", customerHandler.GetCustomers)
	api.POST("/customer/create", customerHandler.CreateCustomer)
	api.PUT("/customer/update/:id", customerHandler.UpdateCustomer)
	api.DELETE("/customer/delete/:id", customerHandler.DeleteCustomer)
	api.GET("/customer/get/:id", customerHandler.GetCustomer)

	itemHandler := v1.NewItemHandler(db)
	api.GET("/items", itemHandler.GetItems)
	api.POST("/item/create", itemHandler.CreateItem)
	api.PUT("/item/update/:id", itemHandler.UpdateItem)
	api.DELETE("/item/delete/:id", itemHandler.DeleteItem)
	api.GET("/item/get/:id", itemHandler.GetItem)

	transactionHandler := v1.NewTransactionHandler(db)
	api.GET("/transactions", transactionHandler.GetTransactions)
	api.POST("/transaction/create", transactionHandler.CreateTransaction)
	api.PUT("/transaction/update/:id", transactionHandler.UpdateTransaction)
	api.DELETE("/transaction/delete/:id", transactionHandler.DeleteTransaction)
	api.GET("/transaction/get/:id", transactionHandler.GetTransaction)
	api.GET("/transaction/details", transactionHandler.GetTransactionDetailsWithCustomerAndItem)
	api.GET("/transaction/filter", transactionHandler.FilterTransactions)
	return r
}
