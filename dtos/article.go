package dtos

type ArticleInputDTO struct {
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

type ArticleInputUpdateDTO struct {
	Author string `json:"author" `
	Title  string `json:"title" `
	Body   string `json:"body" `
}
