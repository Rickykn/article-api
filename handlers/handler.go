package handlers

import "github.com/Rickykn/article-api/services"

type Handler struct {
	articleService services.ArticleService
}

type HandlerConfig struct {
	ArticleService services.ArticleService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		articleService: c.ArticleService,
	}
}
