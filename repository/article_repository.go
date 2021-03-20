package repository

import (
	// "go-tech-blog/model"
	"go-tutorial/model"
)

//ArtitcleList...
func ArtitcleList() ([]*model.Article, error) {
	query := `SELECT * FROM articles;`

	var articles []*model.Article
	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}

	return articles, nil
}
