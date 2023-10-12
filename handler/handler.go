package handler
import (
 "github.com/1Nelsonel/GO_REST_API/database"
 "github.com/1Nelsonel/GO_REST_API/model"
 "github.com/gofiber/fiber/v2"
 "github.com/google/uuid"
)

//Create a user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
   // Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
	 return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Something's wrong with your input", "data": err})
	}
   err = db.Create(&user).Error
	if err != nil {
	 return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Could not create user", "data": err})
	}
   // Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message":  "User has created", "data": user})
   }

// Get All Users from db
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User
   // find all users in the database
	db.Find(&users)
   // If no user found, return an error
	if len(users) == 0 {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
   // return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
   }

// GetSingleUser from db
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db
   // get id params
	id := c.Params("id")
   var user model.User
   // find single user in the database by id
	db.Find(&user, "id = ?", id)
   if user.ID == uuid.Nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
   return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
   }

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
   db := database.DB.Db
   var user model.User
   // get id params
	id := c.Params("id")
   // find single user in the database by id
	db.Find(&user, "id = ?", id)
   if user.ID == uuid.Nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
   var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
	 return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
   user.Username = updateUserData.Username
   user.Email = updateUserData.Email
   user.Password = updateUserData.Password
   // Save the Changes
	db.Save(&user)
   // Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
   }

// func UpdateUser(c *fiber.Ctx) error {
// 	// Create a struct for the updated data
// 	type updateUser struct {
// 		Username string `json:"username"`
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	// Get the user ID from the route parameters
// 	id := c.Params("id")

// 	// Find the user by ID
// 	var user model.User
// 	if err := database.DB.Db.Where("id = ?", id).First(&user).Error; err != nil {
// 		// User not found
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
// 	}

// 	// Parse the request body into an updateUser struct
// 	var updateUserData updateUser
// 	if err := c.BodyParser(&updateUserData); err != nil {
// 		// Bad request
// 		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Bad request", "data": err})
// 	}

// 	// Update user data
// 	user.Username = updateUserData.Username
// 	user.Email = updateUserData.Email
// 	// You should hash the new password before updating it, and handle password hashing appropriately.
// 	// For example, you can use a library like "bcrypt" to hash passwords.
// 	// updateUserData.Password should contain the new hashed password.
// 	user.Password = updateUserData.Password

// 	// Save the changes
// 	if err := database.DB.Db.Save(&user).Error; err != nil {
// 		// Database error
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Database error", "data": err})
// 	}

// 	// Return the updated user
// 	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})
// }


   // delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
   // get id params
	id := c.Params("id")
   // find single user in the database by id
	db.Find(&user, "id = ?", id)
   if user.ID == uuid.Nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
   }
   err := db.Delete(&user, "id = ?", id).Error
   if err != nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
   return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
   }