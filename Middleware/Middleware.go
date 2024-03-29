package middleware

import (
	database "SmochaDeliveryApp/Database"
	"SmochaDeliveryApp/model"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CustomerAuthBridge(c *fiber.Ctx) error {
	//get the cookie off the request body
	//decode and validate
	//check the expiration
	//find user with subject
	//attach to the req body
	//continue
	tokenString := c.Cookies("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check expiration
		expire := claims["expire"].(float64)
		if float64(time.Now().Unix()) > expire {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		// get customer
		var customer model.CustomerDetails
		database.Db.First(&customer, claims["subject"])
		if customer.ID == 0 {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "failed to get user",
			})
		}
		//attach to the request body
		c.Locals("customer", &customer)

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	c.Next()
	return c.SendStatus(fiber.StatusOK)
}

func RiderAuthBridge(c *fiber.Ctx) error {
	//get the cookie off the request body
	//decode and validate
	//check the expiration
	//find user with subject
	//attach to the req body
	//continue
	tokenString := c.Cookies("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check expiration
		expire := claims["expire"].(float64)
		if float64(time.Now().Unix()) > expire {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		// get customer
		var rider model.RiderDetails
		database.Db.First(&rider, claims["subject"])
		if rider.ID == 0 {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "failed to get user",
			})
		}
		//attach to the request body
		c.Locals("customer", &rider)

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	c.Next()
	return c.SendStatus(fiber.StatusOK)
}
