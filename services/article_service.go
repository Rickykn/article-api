package services

import (
	"github.com/Rickykn/article-api/dtos"
	help "github.com/Rickykn/article-api/helpers"
	r "github.com/Rickykn/article-api/repositories"
)

type ArticleService interface {
	GetAllArticle(query *help.Query) (*help.JsonResponse, error)
	CreateNewArticle(articleInput *dtos.ArticleInputDTO) (*help.JsonResponse, error)
	GetArticleById(id int) (*help.JsonResponse, error)
	DeleteArticle(id int) (*help.JsonResponse, error)
	UpdateArticle(id int, input *dtos.ArticleInputUpdateDTO) (*help.JsonResponse, error)
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

	if len(articles) == 0 {
		return help.HandlerError(404, "Data not found", nil), err
	}

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

func (a *articleService) GetArticleById(id int) (*help.JsonResponse, error) {
	article, err := a.articleRepository.FindOne(id)

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}
	return help.HandlerSuccess(200, "Get Article Success", article), nil

}

func (a *articleService) DeleteArticle(id int) (*help.JsonResponse, error) {
	row, err := a.articleRepository.Delete(id)

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}

	if row == 0 {
		return help.HandlerError(404, "Data not found for delete", nil), err
	}

	return help.HandlerSuccess(200, "Success Delete article", nil), nil
}

func (a *articleService) UpdateArticle(id int, input *dtos.ArticleInputUpdateDTO) (*help.JsonResponse, error) {
	newArticle, err := a.articleRepository.Update(id, input.Author, input.Title, input.Body)

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}

	return help.HandlerSuccess(200, "Success Update article", newArticle), nil
}
