package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sb "github.com/somatom98/stylebooking/stylebooking_be"
	vm "github.com/somatom98/stylebooking/stylebooking_be/viewmodels"
)

type CustomerController struct {
	customerService sb.CustomerService
}

func NewCustomerController(customerService sb.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (c *CustomerController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	customers, err := c.customerService.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (c *CustomerController) SignUp(ctx *gin.Context) {
	var request vm.SignUpRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := c.customerService.SignUp(ctx.Request.Context(), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *CustomerController) LogIn(ctx *gin.Context) {
	var request vm.SignInRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := c.customerService.LogIn(ctx.Request.Context(), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
