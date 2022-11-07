package repositories

import (
	"github.com/Rickykn/article-api/helpers"
	"github.com/Rickykn/article-api/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Find(query *helpers.Query) ([]*models.Articel, error)
	Create(author, title, body string) (*models.Articel, error)
}

type articleRepository struct {
	db *gorm.DB
}

type ARConfig struct {
	DB *gorm.DB
}

func NewArticleRepository(c *ARConfig) ArticleRepository {
	return &articleRepository{
		db: c.DB,
	}
}

func (a *articleRepository) Find(query *helpers.Query) ([]*models.Articel, error) {
	var articles []*models.Articel

	filterAuthor := struct {
		Author string
	}{
		Author: query.Author,
	}

	result := a.db.Order("created_at desc").
		Where("title ILIKE ?", "%"+query.Search+"%").Or("body ILIKE ?", "%"+query.Search+"%").
		Where(&filterAuthor).
		Find(&articles)

	return articles, result.Error
}

func (a *articleRepository) Create(author, title, body string) (*models.Articel, error) {
	newArticle := &models.Articel{
		Author: author,
		Title:  title,
		Body:   body,
	}
	result := a.db.Create(&newArticle)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return newArticle, result.Error
}
