package repository

import (
	"errors"
	"fmt"
	"recipe_book/derpl-del/recipe/recipe_server/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RecipeRepositoryPg struct {
	Conn *gorm.DB
}

func NewRecipeRepositoryPg(Conn *gorm.DB) RecipeRepository {
	return &RecipeRepositoryPg{Conn}
}

func (repository *RecipeRepositoryPg) CreateRecipe(data interface{}) error {
	err := repository.Conn.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *RecipeRepositoryPg) QueryRecipe(data *models.RecipeModel) error {
	err := repository.Conn.Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound); err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find data with specified ID: %v", err),
		)
	}
	return nil
}

func (repository *RecipeRepositoryPg) UpdateRecipe(data *models.RecipeModel) error {
	updates := *data
	repository.Conn.Find(&data)
	data.Id = updates.Id
	data.AuthorId = updates.AuthorId
	data.Ingredients = updates.Ingredients
	data.Procedures = updates.Procedures
	data.Title = updates.Title
	err := repository.Conn.Save(&data).Error
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}
	return nil
}

func (repository *RecipeRepositoryPg) DeleteRecipe(data *models.RecipeModel) error {
	err := repository.Conn.Delete(&data).Error
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}
	return nil
}
