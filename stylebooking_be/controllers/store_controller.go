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

	var viewmodels []vm.Store
	for _, store := range stores {
		viewmodels = append(viewmodels, vm.Store{
			Name:        store.Name,
			Description: store.Description,
			Location:    store.Location,
			Hours:       store.Hours,
		})
	}

	ctx.JSON(http.StatusOK, viewmodels)
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

	viewmodel := vm.Store{
		Name:        store.Name,
		Description: store.Description,
		Location:    store.Location,
		Hours:       store.Hours,
	}

	ctx.JSON(http.StatusOK, viewmodel)
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
