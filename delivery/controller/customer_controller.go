package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/go-simple-smtp/model"
	"github.com/jutionck/go-simple-smtp/service"
)

type CustomerController struct {
	router     *gin.Engine
	cstService service.RegisterCustomerUseCase
}

func (cc *CustomerController) createHandler(c *gin.Context) {
	var payload model.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}
	result, err := cc.cstService.RegisterCustomer(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func NewCustomerController(router *gin.Engine, cstService service.RegisterCustomerUseCase) *CustomerController {
	controller := CustomerController{
		router:     router,
		cstService: cstService,
	}

	router.POST("/customers", controller.createHandler)
	return &controller
}
