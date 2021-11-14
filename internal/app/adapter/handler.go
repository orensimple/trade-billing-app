package adapter

import (
	"net/http"

	"github.com/prometheus/common/log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/orensimple/trade-billing-app/internal/app/application/usecase"
	"github.com/orensimple/trade-billing-app/internal/app/domain"
)

func (ctrl Controller) health(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) accountCreate(c *gin.Context) {
	var req domain.CreateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	account, err := usecase.CreateAccount(ctrl.AccountRepository, &domain.Account{ID: uuid.New(), CurrencyCode: req.CurrencyCode, Name: req.Name})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed create account"})

		return
	}

	c.JSON(http.StatusOK, account)
}

func (ctrl Controller) accountGet(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	res, err := usecase.GetAccount(ctrl.AccountRepository, &domain.Account{ID: id})
	if err != nil && err.Error() != "account not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get account"})

		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "account not found"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) accountBlocked(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	account, err := usecase.GetAccount(ctrl.AccountRepository, &domain.Account{ID: id})
	if err != nil && err.Error() != "account not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get account"})

		return
	}
	if account == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "account not found"})

		return
	}

	var req domain.BlockedRequest
	err = c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	account.BlockedAmount = req.BlockedAmount

	err = usecase.UpdateAccountBlocked(ctrl.AccountRepository, account)
	if err != nil {
		log.Error(err)

		if err.Error() == "Error 3819: Check constraint 'blocked_amount_value' is violated." {
			c.JSON(http.StatusConflict, domain.SimpleResponse{Status: "haven't money"})

			return
		}
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed blocked money"})

		return
	}

	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) accountPay(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	account, err := usecase.GetAccount(ctrl.AccountRepository, &domain.Account{ID: id})
	if err != nil && err.Error() != "account not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get account"})

		return
	}
	if account == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "account not found"})

		return
	}

	var req domain.PayRequest
	err = c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	account.Balance = req.PayAmount

	err = usecase.UpdateAccountPay(ctrl.AccountRepository, account)
	if err != nil {
		log.Error(err)

		if err.Error() == "Error 3819: Check constraint 'blocked_amount_zero' is violated." {
			c.JSON(http.StatusConflict, domain.SimpleResponse{Status: "haven't money"})

			return
		}
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed pay money"})

		return
	}

	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) accountDelete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	err = usecase.DeleteAccount(ctrl.AccountRepository, &domain.Account{ID: id})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed delete account"})

		return
	}

	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}
