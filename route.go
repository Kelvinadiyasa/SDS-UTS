package route

import (
	"strconv"
	"uts/database"
	"uts/models"

	"github.com/gofiber/fiber/v2"
)

// Insert Data sesuai dengan Collection Postman
func InsertData(c *fiber.Ctx) error {
	user := models.User{
		Nama:     "Muhamad Kelvin Adiyasa",
		Email:    "Kelvin.sasli@gmail.com",
		Username: "Kelvinadiyasa11",
		Password: "Kelvinady11_",
	}

	err := database.CreateUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil di tambahkan",
	})
}

// Mengambil data untuk semua user
func GetAllData(c *fiber.Ctx) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

// Mengambil data dari id_user berdasarkan Parameter
func GetUserByid(c *fiber.Ctx) error {
	id_user := c.Params("id_user")

	userID, err := strconv.ParseUint(id_user, 10, 64)
	if err != nil {
		return err
	}

	user, err := database.GetUserByID(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// Menghapus data user berdasarkan id_user yang di tempatkan di parameter
func Delete(c *fiber.Ctx) error {
	id_user := c.Params("id_user")

	userID, err := strconv.ParseUint(id_user, 10, 64)
	if err != nil {
		return err
	}

	err = database.DeleteUserByID(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"Pesan": "Data telah di hapus",
	})
}

// Mengupdate data user berdasarkan id_user yang di tempatkan di parameter
func Update(c *fiber.Ctx) error {
	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	userID, err := strconv.ParseUint(id_user, 10, 64)
	if err != nil {
		return err
	}

	user, err := database.GetUserByID(uint(userID))
	if err != nil {
		return err
	}

	update := models.User{
		Nama:     data["Nama"],
		Email:    data["Email"],
		Username: data["Username"],
		Password: data["Password"],
	}

	user.Nama = update.Nama
	user.Email = update.Email
	user.Username = update.Username
	user.Password = update.Password

	err = database.UpdateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di Update",
	})
}
