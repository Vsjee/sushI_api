package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sushi []SushiElement

type SushiElement struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	SushiName   string             `bson:"sushi_name" json:"sushi_name,omitempty" validate:"required"`
	Price       int64              `bson:"price" json:"price,omitempty" validate:"required"`
	Discount    int64              `bson:"discount" json:"discount,omitempty" validate:"required"`
	Image       string             `bson:"image" json:"image,omitempty" validate:"required"`
	Ingredients []Ingredient       `json:"ingredients,omitempty" validate:"required"`
	SushiOrigin string             `bson:"sushi_origin" json:"sushi_origin,omitempty" validate:"required"`
	Rating      int64              `bson:"rating" json:"rating,omitempty" validate:"required"`
}

type Ingredient struct {
	IngredientName string `bson:"ingredient_name" json:"ingredient_name,omitempty" validate:"required"`
	Calories       int64  `bson:"calories" json:"calories,omitempty" validate:"required"`
}
