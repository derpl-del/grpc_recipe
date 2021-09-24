package main

import (
	"fmt"
	"log"
	"net"
	"recipe_book/derpl-del/recipe/recipe_server/delivery"
	"recipe_book/derpl-del/recipe/recipe_server/models"
	"recipe_book/derpl-del/recipe/recipe_server/repository"
	"recipe_book/derpl-del/recipe/recipe_server/usecase"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Running")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	connection := "postgres://postgres:welcome1@localhost:5432/banana"
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		//	Logger: logger.Default.LogMode(logger.Silent)
	},
	)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.RecipeModel{})
	recipeRepo := repository.NewRecipeRepositoryPg(db)
	recipeUsecase := usecase.NewRecipeUsecase(recipeRepo)
	s := grpc.NewServer()
	delivery.NewRecipeServer(s, recipeUsecase)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("Morning")
}
