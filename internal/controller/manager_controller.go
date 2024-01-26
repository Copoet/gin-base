package controllers

import (
	"gin-base/internal/model"
	services "gin-base/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MangerController struct {
	ManagerService *services.ManagerService
}

func (m *MangerController) GetManagerList(c *gin.Context) {

	var query model.ManagerQuery
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")

	query.Keyword = &keyword

	data, err := m.ManagerService.GetManagerList(query, page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
