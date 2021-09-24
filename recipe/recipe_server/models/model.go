package models

import (
	"github.com/lib/pq"
)

type RecipeModel struct {
	Id          string         `json:"id"`
	AuthorId    string         `json:"author_id"`
	Title       string         `json:"tittle"`
	Ingredients pq.StringArray `json:"ingredients" gorm:"type:text[]"`
	Procedures  pq.StringArray `json:"Procedures" gorm:"type:text[]"`
}
