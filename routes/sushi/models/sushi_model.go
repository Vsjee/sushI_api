package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sushi []SushiElement

func UnmarshalSushi(data []byte) (Sushi, error) {
	var r Sushi
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Sushi) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SushiElement struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	SushiName   string             `json:"sushi_name,omitempty" validate:"required"`
	Price       int64              `json:"price,omitempty" validate:"required"`
	Discount    int64              `json:"discount,omitempty" validate:"required"`
	Image       string             `json:"image,omitempty" validate:"required"`
	Ingredients []Ingredient       `json:"ingredients,omitempty" validate:"required"`
	SushiOrigin string             `json:"sushi_origin,omitempty" validate:"required"`
	Rating      int64              `json:"rating,omitempty" validate:"required"`
}

type Ingredient struct {
	IngredientName string `json:"ingredient_name,omitempty" validate:"required"`
	Calories       int64  `json:"calories,omitempty" validate:"required"`
}
