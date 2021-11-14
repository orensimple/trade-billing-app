package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/orensimple/trade-billing-app/internal/app/adapter/mysql"
	"github.com/orensimple/trade-billing-app/internal/app/adapter/repository"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// Controller is a controller
type Controller struct {
	AccountRepository repository.Account
}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	db := mysql.Connection()

	// init prometheus metrics
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	accountRepository := repository.NewAccountRepo(db)

	ctrl := Controller{
		AccountRepository: accountRepository,
	}

	r.GET("/health", ctrl.health)

	api := r.Group("/api")
	api.POST("/account", ctrl.accountCreate)
	api.GET("/account/:id", ctrl.accountGet)
	api.POST("/account/:id/block", ctrl.accountBlocked)
	api.POST("/account/:id/pay", ctrl.accountPay)
	api.DELETE("/account/:id", ctrl.accountDelete)

	return r
}
