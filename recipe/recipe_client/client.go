package main

import (
	"context"
	"fmt"
	"log"
	"recipe_book/derpl-del/recipe/recipepb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := recipepb.NewRecipeServiceClient(cc)
	req := &recipepb.CreateRecipeRequest{
		Recipe: &recipepb.Recipe{
			Id:          "2",
			AuthorId:    "2",
			Title:       "test hello",
			Ingredients: []string{"1.do", "2.go", "3.pus"},
			Procedures:  []string{"1.aaa", "2.bbb", "3.cccc"},
		},
	}
	res, err := c.CreateRecipe(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res)
}
