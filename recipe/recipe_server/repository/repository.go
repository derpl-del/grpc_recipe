package repository

import "recipe_book/derpl-del/recipe/recipe_server/models"

type RecipeRepository interface {
	CreateRecipe(data interface{}) error
	QueryRecipe(data *models.RecipeModel) error
	UpdateRecipe(data *models.RecipeModel) error
	DeleteRecipe(data *models.RecipeModel) error
}
