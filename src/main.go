package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func getAll(c *gin.Context) {
	customers := SelectAllCustomers()
	c.IndentedJSON(http.StatusOK, customers)
}

func getById(c *gin.Context) {
	id := c.Param("id")

	customer, err := SelectCustomerById(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func create(c *gin.Context) {
	var customer Customer
	err := c.BindJSON(&customer)
	if err != nil {
		fmt.Println(err)
	}

	InsertCustomer(customer)

	c.IndentedJSON(http.StatusOK, gin.H{"result": "OK"})

}

func update(c *gin.Context) {
	var customer Customer
	err := c.BindJSON(&customer)
	if err != nil {
		fmt.Println(err)
	}

	updateCustomer(customer);

	c.IndentedJSON(http.StatusOK, gin.H{"result": "OK"})
}

func delete(c *gin.Context) {
	var customer Customer
	err := c.BindJSON(&customer)
	if err != nil {
		fmt.Println(err)
	}

	deleteCustomer(customer.ID)

	c.IndentedJSON(http.StatusOK, gin.H{"result": "OK"})
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/", getAll)
	router.GET("/:id", getById)
	router.POST("/", create)
	router.PUT("/", update)
	router.DELETE("/", delete)

	router.Run(":8080")
}
