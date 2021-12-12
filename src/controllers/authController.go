package controllers

import (
	"promo/src/database"
	"promo/src/models"
	"promo/src/middleware"
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["confirm_password"] {
		return c.Status(400).JSON(fiber.Map{"message": "Passwords do not match"})
	}

	user := models.User{
		FirstName:  data["first_name"],
		LastName:   data["last_name"],
		Email:      data["email"],
		IsPromoter: false,
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	user := models.User{}
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "User does not exist"})
	}
	if err := user.CheckPassword(data["password"]); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid password"})
	}
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error while generating token"})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "Successfully logged in"})
}

func GetAuthUser(c *fiber.Ctx) error {
	id,_ := middleware.GetUserId(c)
	user := models.User{}
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie:= fiber.Cookie{
		Name:"jwt",
		Value:"",
		Expires: time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "Successfully logged out"})
}