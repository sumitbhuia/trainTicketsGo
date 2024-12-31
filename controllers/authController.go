package controllers

import (
	"context"
	"net/http"
	"os"
	"time"
	"trainTicketsGo/database"
	"trainTicketsGo/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Register(c *gin.Context) {
	var user models.User

	// Validate input
	// ShouldBindJSON is a method that binds the request body to the user struct , meaning it will take the request body and map it to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password (use bcrypt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// save the hashed password in the user object and create the user
	user.Password = string(hashedPassword)

	// Define collection and insert user
	collection := database.DB.Database(os.Getenv("DB_NAME")).Collection("users")
	_, err = collection.InsertOne(context.TODO(), user) // In the collection named "users", insert the user object

	// If no error interting the user to the database in the collection named "users", return a success message
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login a user
func Login(c *gin.Context) {

	// Define the credentials struct
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the request body to the credentials struct , if there is an error, return a bad request error
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Access the "users" collection in the database
	collection := database.DB.Database(os.Getenv("DB_NAME")).Collection("users")

	// Retrieve the user with the provided email
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": credentials.Email}).Decode(&user) // Find the user with the email provided in the credentials

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the password provided in the credentials with the password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID.Hex(),                         // MongoDB ObjectID as a string
		"exp": time.Now().Add(24 * time.Hour).Unix(), // Expiry time
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// Return the token , if there is no error
	// For the token to be used in the frontend, it should be stored in the local storage or in a cookie
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
