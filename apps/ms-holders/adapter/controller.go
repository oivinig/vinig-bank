package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/vinigbank/apps/msholders/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/msholders/adapter/repository"
	"github.com/oivinig/vinigbank/apps/msholders/application/usecase"
)

var (
	holderRepository = repository.Holder{}
)

// Controller is a controller
type Controller struct{}

// Router is routing
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	r.POST("/holder", ctrl.CreateHolder)
	r.GET("/holder", ctrl.FetchHolders)
	// r.GET("/holder/document/:document", holder.FindHolderByDocument)
	r.PATCH("/holder/id/:id", ctrl.EditHolder)
	// r.DELETE("/holder/document/:document", holder.DeleteHolder)

	return r
}

func (ctrl Controller) CreateHolder(c *gin.Context) {
	var h model.Holder
	if err := c.ShouldBindJSON(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	holder := usecase.CreateHolder(holderRepository, &h)
	c.JSON(http.StatusOK, holder)
}

func (ctrl Controller) FetchHolders(c *gin.Context) {
	holders := usecase.FetchHolders(holderRepository)
	c.JSON(http.StatusOK, holders)
}

func (ctrl Controller) EditHolder(c *gin.Context) {
	var h model.Holder
	if err := c.ShouldBindJSON(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	holder := usecase.EditHolder(holderRepository, &h)
	c.JSON(http.StatusOK, holder)
}
