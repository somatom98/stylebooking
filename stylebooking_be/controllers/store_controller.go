package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sb "github.com/somatom98/stylebooking/stylebooking_be"
	vm "github.com/somatom98/stylebooking/stylebooking_be/viewmodels"
)

type StoreController struct {
	storeService sb.StoreService
}

func NewStoreController(storeService sb.StoreService) *StoreController {
	return &StoreController{
		storeService: storeService,
	}
}

func (c *StoreController) GetAll(ctx *gin.Context) {
	stores, err := c.storeService.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, stores)
}

func (c *StoreController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	store, err := c.storeService.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, store)
}

func (c *StoreController) Create(ctx *gin.Context) {
	var store vm.Store
	if err := ctx.ShouldBindJSON(&store); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.storeService.Create(ctx.Request.Context(), store); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, store)
}

func (c *StoreController) AddService(ctx *gin.Context) {
	storeId := ctx.Param("id")

	var service vm.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.storeService.AddService(ctx.Request.Context(), storeId, service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, service)
}

func (c *StoreController) UpdateService(ctx *gin.Context) {
	storeId := ctx.Param("id")
	serviceId := ctx.Param("serviceId")

	var service vm.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.storeService.UpdateService(ctx.Request.Context(), storeId, serviceId, service); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, service)
}

func (c *StoreController) DeleteService(ctx *gin.Context) {
	storeId := ctx.Param("id")
	serviceId := ctx.Param("serviceId")

	if err := c.storeService.DeleteService(ctx.Request.Context(), storeId, serviceId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
