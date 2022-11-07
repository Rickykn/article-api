package handlers

import (
	"net/http"

	"github.com/Rickykn/article-api/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateArticle(c *gin.Context) {
	var articleInput *dtos.ArticleInputDTO

	err := c.ShouldBindJSON(&articleInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	response, _ := h.articleService.CreateNewArticle(articleInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	}
}
