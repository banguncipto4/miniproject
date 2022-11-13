package main

import (
	"MINIPROJECT/database"
	"MINIPROJECT/models"
	"MINIPROJECT/routes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func newApp() *echo.Echo {
	database.InitTestDB()

	app := echo.New()

	routes.SetupRoute(app)

	return app
}
func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		database.CleanSeeders()
	}
}

func getJWTToken(t *testing.T) string {
	user := database.SeedUser()

	var userRequest *models.Register = &models.Register{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(newApp()).
		Post("/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

// Test All_Function Game

func TestGetgames_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/game").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetgame_NotFound(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/game/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreategame_Success(t *testing.T) {

	var gameRequest *models.InputGame = &models.InputGame{
		Game_name:   "game_name",
		Game_type:   "game_type",
		Game_desc:   "game_desc",
		Game_access: "game_access",
		PublisherID: 1,
		RatingID:    1,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/game").
		Header("Authorization", token).
		JSON(gameRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreategame_ValidationFailed(t *testing.T) {
	var gameRequest *models.InputGame = &models.InputGame{}

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Post("/game").
		Header("Authorization", token).
		JSON(gameRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdategame_Success(t *testing.T) {
	var game models.Game = database.SeedGame()

	var gameRequest *models.InputGame = &models.InputGame{
		Game_name:   "game_name",
		Game_type:   "game_type",
		Game_desc:   "game_desc",
		Game_access: "game_access",
		PublisherID: 1,
		RatingID:    1,
	}

	gameID := strconv.Itoa(int(game.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/game/"+gameID).
		Header("Authorization", token).
		JSON(gameRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdategame_ValidationFailed(t *testing.T) {
	var game models.Game = database.SeedGame()

	var gameRequest *models.InputGame = &models.InputGame{}

	gameID := strconv.Itoa(int(game.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Put("/game/"+gameID).
		Header("Authorization", token).
		JSON(gameRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeletegame_Success(t *testing.T) {
	var game models.Game = database.SeedGame()

	var token string = getJWTToken(t)

	gameID := strconv.Itoa(int(game.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Delete("/game/"+gameID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeletegame_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Delete("/game/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

// Test All_Function Publisher

func TestGetpublishers_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/publisher").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetpublisher_NotFound(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/publisher/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreatepublisher_Success(t *testing.T) {

	var publisherRequest *models.InputPublisher = &models.InputPublisher{
		Publisher_name: "publisher_name",
		Publisher_desc: "publisher_desc",
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/publisher").
		Header("Authorization", token).
		JSON(publisherRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreatepublisher_ValidationFailed(t *testing.T) {
	var publisherRequest *models.InputPublisher = &models.InputPublisher{}

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Post("/publisher").
		Header("Authorization", token).
		JSON(publisherRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdatepublisher_Success(t *testing.T) {
	var publisher models.Publisher = database.SeedPublisher()

	var publisherRequest *models.InputPublisher = &models.InputPublisher{
		Publisher_name: "publisher_name",
		Publisher_desc: "publisher_desc",
	}

	publisherID := strconv.Itoa(int(publisher.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/publisher/"+publisherID).
		Header("Authorization", token).
		JSON(publisherRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdatepublisher_ValidationFailed(t *testing.T) {
	var publisher models.Publisher = database.SeedPublisher()

	var publisherRequest *models.InputPublisher = &models.InputPublisher{}

	publisherID := strconv.Itoa(int(publisher.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Put("/publisher/"+publisherID).
		Header("Authorization", token).
		JSON(publisherRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeletepublisher_Success(t *testing.T) {
	var publisher models.Publisher = database.SeedPublisher()

	var token string = getJWTToken(t)

	publisherID := strconv.Itoa(int(publisher.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Delete("/publisher/"+publisherID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeletepublisher_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Delete("/publisher/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

// Test All_Function Rating

func TestGetratings_Success(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Get("/rating").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetrating_NotFound(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Get("/rating/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreaterating_Success(t *testing.T) {

	var ratingRequest *models.InputRating = &models.InputRating{
		Star:     1,
		Reaction: "bad",
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/rating").
		Header("Authorization", token).
		JSON(ratingRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestCreaterating_ValidationFailed(t *testing.T) {
	var ratingRequest *models.InputRating = &models.InputRating{}

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Post("/rating").
		Header("Authorization", token).
		JSON(ratingRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdaterating_Success(t *testing.T) {
	var rating models.Rating = database.SeedRating()

	var ratingRequest *models.InputGame = &models.InputGame{
		Game_name:   "game_name",
		Game_type:   "game_type",
		Game_desc:   "game_desc",
		Game_access: "game_access",
	}

	ratingID := strconv.Itoa(int(rating.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Put("/rating/"+ratingID).
		Header("Authorization", token).
		JSON(ratingRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdaterating_ValidationFailed(t *testing.T) {
	var rating models.Rating = database.SeedRating()

	var ratingRequest *models.InputRating = &models.InputRating{}

	ratingID := strconv.Itoa(int(rating.ID))

	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Put("/rating/"+ratingID).
		Header("Authorization", token).
		JSON(ratingRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeleterating_Success(t *testing.T) {
	var rating models.Rating = database.SeedRating()

	var token string = getJWTToken(t)

	ratingID := strconv.Itoa(int(rating.ID))

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Delete("/rating/"+ratingID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleterating_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		Handler(newApp()).
		Observe(cleanup).
		Delete("/rating/-1").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}
