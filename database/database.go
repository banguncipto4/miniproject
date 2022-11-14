package database

import (
	"MINIPROJECT/models"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("connected")
	AutoMigrate()
}

func main() {
	InitDB()
}

func AutoMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Publisher{},
		&models.Rating{},
		&models.Game{},
	)
}

func InitTestDB() {
	var err error

	namaKoneksi := "root:@/miniproject?parseTime=true"

	DB, err = gorm.Open(postgres.Open(namaKoneksi), &gorm.Config{})

	if err != nil {
		fmt.Println("tidak terkoneksi")
	}

	fmt.Println("terkoneksi")

	AutoMigrate()
}

// SeedGame returns game sample data
func SeedGame() models.Game {
	// add publisher and rating data as well
	publisher := SeedPublisher()
	rating := SeedRating()

	var game models.Game = models.Game{
		Game_name:   "game_name",
		Game_type:   "game_type",
		Game_desc:   "game_desc",
		Game_access: "game_access",
		PublisherID: publisher.ID,
		RatingID:    rating.ID,
	}

	if err := DB.Create(&game).Error; err != nil {
		panic(err)
	}

	var createdgame models.Game

	DB.Last(&createdgame)

	return createdgame
}

func SeedPublisher() models.Publisher {

	var publisher models.Publisher = models.Publisher{
		Publisher_name: "publisher_name",
		Publisher_desc: "publisher_desc",
	}

	if err := DB.Create(&publisher).Error; err != nil {
		panic(err)
	}

	var createdpublisher models.Publisher

	DB.Last(&createdpublisher)

	return createdpublisher
}

func SeedRating() models.Rating {

	var rating models.Rating = models.Rating{
		// choose uint (1 - 5)
		Star: 1,
		// bad or fun or happy or unique
		Reaction: "bad",
	}

	if err := DB.Create(&rating).Error; err != nil {
		panic(err)
	}

	var createdrating models.Rating

	DB.Last(&createdrating)

	return createdrating
}

func CleanSeeders() {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	userResult := DB.Exec("DELETE FROM users")
	itemgameResult := DB.Exec("DELETE FROM games")
	itempublisherResult := DB.Exec("DELETE FROM publishers")
	itemratingResult := DB.Exec("DELETE FROM ratings")

	var isFailed bool = itemgameResult.Error != nil || userResult.Error != nil || itempublisherResult.Error != nil || itemratingResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}

func SeedUser() models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user models.User = models.User{
		Name:     "test",
		Email:    "testing@mail.com",
		Password: string(password),
	}

	if err := DB.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser models.User

	DB.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}
