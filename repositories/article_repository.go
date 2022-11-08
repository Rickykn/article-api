package repositories

import (
	"github.com/Rickykn/article-api/helpers"
	"github.com/Rickykn/article-api/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Find(query *helpers.Query) ([]*models.Articel, error)
	Create(author, title, body string) (*models.Articel, error)
	FindOne(id int) (*models.Articel, error)
	Delete(id int) (int, error)
	Update(id int, author, title, body string) (*models.Articel, error)
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

	if query.Author != "" && query.Search != "" {
		result := a.db.Order("created_at desc").
			Where(&filterAuthor).
			Where("title ILIKE ?", "%"+query.Search+"%").Or("body ILIKE ?", "%"+query.Search+"%").
			Find(&articles)

		return articles, result.Error
	} else if query.Author != "" {

		result := a.db.Order("created_at desc").
			Where(&filterAuthor).
			Find(&articles)

		return articles, result.Error

	} else {
		result := a.db.Order("created_at desc").
			Where("title ILIKE ?", "%"+query.Search+"%").Or("body ILIKE ?", "%"+query.Search+"%").
			Find(&articles)
		return articles, result.Error
	}

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

func (a *articleRepository) FindOne(id int) (*models.Articel, error) {
	var article *models.Articel

	result := a.db.Where("id = ?", id).First(&article)

	return article, result.Error
}

func (a *articleRepository) Delete(id int) (int, error) {
	var article *models.Articel
	result := a.db.Where("id = ?", id).Delete(&article)

	return int(result.RowsAffected), result.Error
}

func (a *articleRepository) Update(id int, author, title, body string) (*models.Articel, error) {
	newArticle := models.Articel{
		Author: author,
		Title:  title,
		Body:   body,
	}

	result := a.db.Where("id=?", id).Updates(newArticle)

	return &newArticle, result.Error
}
