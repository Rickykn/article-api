package services

import (
	"github.com/Rickykn/article-api/dtos"
	help "github.com/Rickykn/article-api/helpers"
	r "github.com/Rickykn/article-api/repositories"
)

type ArticleService interface {
	GetAllArticle(query *help.Query) (*help.JsonResponse, error)
	CreateNewArticle(articleInput *dtos.ArticleInputDTO) (*help.JsonResponse, error)
}

type articleService struct {
	articleRepository r.ArticleRepository
}

type ASConfig struct {
	ArticleRepository r.ArticleRepository
}

func NewArticleService(c *ASConfig) ArticleService {
	return &articleService{
		articleRepository: c.ArticleRepository,
	}
}

func (a *articleService) GetAllArticle(query *help.Query) (*help.JsonResponse, error) {
	articles, err := a.articleRepository.Find(query)

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}

	return help.HandlerSuccess(200, "Get All Article Success", articles), nil
}

func (a *articleService) CreateNewArticle(articleInput *dtos.ArticleInputDTO) (*help.JsonResponse, error) {

	newArticle, err := a.articleRepository.Create(articleInput.Author, articleInput.Title, articleInput.Body)

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}

	return help.HandlerSuccess(201, "Create new article success", newArticle), nil
}
