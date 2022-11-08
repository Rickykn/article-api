package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rickykn/article-api/dtos"
	"github.com/Rickykn/article-api/helpers"
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

func (h *Handler) GetAllArticle(c *gin.Context) {

	search := c.Query("search")
	author := c.Query("author")

	query := &helpers.Query{
		Search: search,
		Author: author,
	}

	response, _ := h.articleService.GetAllArticle(query)

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

func (h *Handler) GetArticleById(c *gin.Context) {
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)
	response, _ := h.articleService.GetArticleById(convId)

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

func (h *Handler) DeleteArticleById(c *gin.Context) {
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)
	response, _ := h.articleService.DeleteArticle(convId)

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
func (h *Handler) UpdateArticleById(c *gin.Context) {
	var articleUpdateInput *dtos.ArticleInputUpdateDTO

	err := c.ShouldBindJSON(&articleUpdateInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)
	response, _ := h.articleService.UpdateArticle(convId, articleUpdateInput)

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
