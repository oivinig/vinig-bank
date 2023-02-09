package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/vinigbank/apps/ms-accounts/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/ms-accounts/adapter/repository"
	"github.com/oivinig/vinigbank/apps/ms-accounts/application/usecase"
)

var (
	accountRepository = repository.Account{}
)

// Controller is a controller
type Controller struct{}

// Router is routing
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	r.POST("/account", ctrl.OpenAccount)

	return r
}

func (ctrl Controller) OpenAccount(c *gin.Context) {
	var a model.Account
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	holder := usecase.OpenAccount(accountRepository, &a)
	c.JSON(http.StatusOK, holder)
}
