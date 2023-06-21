package api

import (
	"fmt"
	"github.com/892294101/jxlife/server/constant"
	"github.com/892294101/jxlife/server/models/web"
	"github.com/892294101/jxlife/server/response"
	"github.com/892294101/jxlife/server/service"

	"github.com/gin-gonic/gin"
)

type WebStatistics struct {
	service.WebStatisticsService
}

func GetWebStatistics() *WebStatistics {
	return &WebStatistics{}
}

func (s *WebStatistics) GetTodayData(c *gin.Context) {
	var param web.DataParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		response.Failed(constant.ParamInvalid, c)
		return
	}
	todayData := s.TodayData(param)
	response.Success(constant.Selected, todayData, c)
}

func (s *WebStatistics) GetOrderData(c *gin.Context) {
	var param web.DataParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	orderData := s.OrderData(param)
	response.Success(constant.Selected, orderData, c)
}

func (s *WebStatistics) GetShopData(c *gin.Context) {
	var param web.DataParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	shopData := s.ShopData(param)
	response.Success(constant.Selected, shopData, c)
}
