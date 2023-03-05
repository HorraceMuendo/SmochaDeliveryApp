package handlers

import (
	customers "SmochaDeliveryApp/Customers"
	database "SmochaDeliveryApp/Database"

	"github.com/gofiber/fiber/v2"
)

func GetCustomer(c *fiber.Ctx) error {
	Cdb := database.Db
	var customerDetails []customers.CustomerDetails
	Cdb.Find(&customerDetails)
	return c.Status(200).JSON(customerDetails)
}

func GetCustomerById(c *fiber.Ctx) error {
	id := c.Params("id")
	Cdb := database.Db
	var customerDetail []customers.CustomerDetails
	match := Cdb.Find(&customerDetail, id)

	if match.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&customerDetail)

}
func CreateCustomer(c *fiber.Ctx) error {
	Cdb := database.Db
	customer := new(customers.CustomerDetails)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	Cdb.Create(customer)
	return c.Status(200).JSON(customer)

}

func UpdateCustomer(c *fiber.Ctx) error {
	Cdb := database.Db
	customer := new(customers.CustomerDetails)
	id := c.Params("id")
	if err := c.BodyParser(customer); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	Cdb.Where("id=?", id).Updates(&customer)
	return c.Status(200).JSON(customer)

}
func DeleteCustomer(c *fiber.Ctx) error {
	Cdb := database.Db
	var customer customers.CustomerDetails
	id := c.Params("id")
	delete := Cdb.Delete(&customer, id)

	if delete.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
