package controllers

import (
	"connect-custom-hash-srv/app/models/api"
	"connect-custom-hash-srv/app/models/responses"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func verifyCustomHashPassword(userPassword, storedHash string) bool {
	// TODO Implement your own hash verify method
	return false
}

func Process(c *fiber.Ctx) error {

	body := &api.PasswordVerification{}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	fmt.Println(string(body.Password), string(body.Algorithm))

	if verifyCustomHashPassword(body.Password, body.PasswordVerification.PasswordHash) {
		fmt.Println("Password verified successfully!")

		// Create bcrypt hash
		return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Success: true, Data: &api.PasswordVerifiedData{
			Verified: true,
		}})

	} else {
		fmt.Println("Password verification failed!")
		return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Success: true, Data: &api.PasswordVerifiedData{
			Verified: false,
		}})
	}
}

func Ping(c *fiber.Ctx) error {

	result := "ping"

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Success: true, Data: result})
}
