package delivery

import (
	"context"
	"recipe_book/derpl-del/recipe/recipe_server/models"
	"recipe_book/derpl-del/recipe/recipe_server/usecase"
	"recipe_book/derpl-del/recipe/recipepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	usecase usecase.RecipeUsecase
}

func NewRecipeServer(gserver *grpc.Server, usecase usecase.RecipeUsecase) {
	RecipeServer := &server{usecase}
	recipepb.RegisterRecipeServiceServer(gserver, RecipeServer)
	reflection.Register(gserver)
}

func (s *server) CreateRecipe(ctx context.Context, request *recipepb.CreateRecipeRequest) (*recipepb.CreateRecipeResponse, error) {
	data := models.RecipeModel{
		Id:          request.Recipe.Id,
		AuthorId:    request.Recipe.AuthorId,
		Title:       request.Recipe.Title,
		Ingredients: request.Recipe.Ingredients,
		Procedures:  request.Recipe.Procedures,
	}
	err := s.usecase.CreateRecipe(data)
	if err != nil {
		return nil, err
	}
	res := &recipepb.CreateRecipeResponse{
		Recipe: &recipepb.Recipe{
			Id:          data.Id,
			AuthorId:    data.AuthorId,
			Title:       data.AuthorId,
			Ingredients: data.Ingredients,
			Procedures:  data.Procedures,
		},
	}
	return res, nil
}

func (s *server) QueryRecipe(ctx context.Context, request *recipepb.QueryRecipeRequest) (*recipepb.QueryRecipeResponse, error) {
	if len(request.Id) == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot find data with empty id",
		)
	}
	data := models.RecipeModel{Id: request.Id}
	err := s.usecase.QueryRecipe(&data)
	if err != nil {
		return nil, err
	}
	res := &recipepb.QueryRecipeResponse{
		Recipe: &recipepb.Recipe{
			Id:          data.Id,
			AuthorId:    data.AuthorId,
			Title:       data.Title,
			Ingredients: data.Ingredients,
			Procedures:  data.Procedures,
		},
	}
	return res, nil
}

func (s *server) UpdateRecipe(ctx context.Context, request *recipepb.UpdateRecipeRequest) (*recipepb.UpdateRecipeResponse, error) {
	if len(request.Recipe.Id) == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot update data with empty id",
		)
	}
	data := models.RecipeModel{
		Id:          request.Recipe.Id,
		AuthorId:    request.Recipe.AuthorId,
		Title:       request.Recipe.Title,
		Ingredients: request.Recipe.Ingredients,
		Procedures:  request.Recipe.Procedures,
	}
	err := s.usecase.UpdateRecipe(&data)
	if err != nil {
		return nil, err
	}
	res := recipepb.UpdateRecipeResponse{
		Recipe: &recipepb.Recipe{
			Id:          data.Id,
			AuthorId:    data.AuthorId,
			Title:       data.Title,
			Ingredients: data.Ingredients,
			Procedures:  data.Procedures,
		},
	}
	return &res, nil
}

func (s *server) DeleteRecipe(ctx context.Context, request *recipepb.DeleteRecipeRequest) (*recipepb.DeleteRecipeResponse, error) {
	if len(request.Id) == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot delete data with empty id",
		)
	}
	data := models.RecipeModel{Id: request.Id}
	err := s.usecase.DeleteRecipe(&data)
	if err != nil {
		return nil, err
	}
	res := recipepb.DeleteRecipeResponse{Id: request.Id}
	return &res, nil
}
