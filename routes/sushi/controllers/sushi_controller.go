package controllers

import (
	"context"
	"sushi-api/configs"
	"sushi-api/routes/sushi/models"
	"sushi-api/routes/sushi/responses"

	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var questionsCollection *mongo.Collection = configs.GetCollection(configs.DB, "all_sushi")
var validate = validator.New()

func GetAllSushi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var questions []models.SushiElement
	defer cancel()

	results, err := questionsCollection.Find(ctx, bson.M{})

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuestion models.SushiElement
		if err = results.Decode(&singleQuestion); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.SushiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		questions = append(questions, singleQuestion)
	}

	return c.Status(http.StatusOK).JSON(
		questions,
	)
}
