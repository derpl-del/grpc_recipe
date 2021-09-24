package usecase

import (
	"recipe_book/derpl-del/recipe/recipe_server/models"
	"recipe_book/derpl-del/recipe/recipe_server/repository"
)

type RecipeUsecase interface {
	CreateRecipe(data interface{}) error
	QueryRecipe(data *models.RecipeModel) error
	UpdateRecipe(data *models.RecipeModel) error
	DeleteRecipe(data *models.RecipeModel) error
}

type recipeUsecase struct {
	recipeRepository repository.RecipeRepository
}

func NewRecipeUsecase(a repository.RecipeRepository) RecipeUsecase {
	return &recipeUsecase{a}
}

func (usecase *recipeUsecase) CreateRecipe(data interface{}) error {
	err := usecase.recipeRepository.CreateRecipe(data)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *recipeUsecase) QueryRecipe(data *models.RecipeModel) error {
	err := usecase.recipeRepository.QueryRecipe(data)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *recipeUsecase) UpdateRecipe(data *models.RecipeModel) error {
	err := usecase.recipeRepository.UpdateRecipe(data)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *recipeUsecase) DeleteRecipe(data *models.RecipeModel) error {
	err := usecase.recipeRepository.DeleteRecipe(data)
	if err != nil {
		return err
	}
	return nil
}
