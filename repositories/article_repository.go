package repositories

import (
	"github.com/Rickykn/article-api/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Find()
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

func (a *articleRepository) Find() {

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
