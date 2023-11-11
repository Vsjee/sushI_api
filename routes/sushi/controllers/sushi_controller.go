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

var sushisCollection *mongo.Collection = configs.GetCollection(configs.DB, "all_sushi")
var validate = validator.New()

func GetAllSushi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var sushis []models.SushiElement
	defer cancel()

	results, err := sushisCollection.Find(ctx, bson.M{})

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleSushi models.SushiElement
		if err = results.Decode(&singleSushi); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.SushiResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		sushis = append(sushis, singleSushi)
	}

	return c.Status(http.StatusOK).JSON(
		sushis,
	)
}
