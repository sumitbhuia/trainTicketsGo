package controllers

import (
	"context"
	"net/http"
	"os"
	"trainTicketsGo/database"
	"trainTicketsGo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BookTicket(c *gin.Context) {
	// Input must be in the form of a JSON object with the following fields
	var requestData struct {
		UserID  string `json:"user_id"`
		TrainID string `json:"train_id"`
		Seat    string `json:"seat"`
	}

	// Extract the data from the request body and map it to the ticket struct
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Convert UserID and TrainID to primitive.ObjectID
	userID, err := primitive.ObjectIDFromHex(requestData.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
		return
	}

	// Convert UserID and TrainID to primitive.ObjectID
	trainID, err := primitive.ObjectIDFromHex(requestData.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TrainID"})
		return
	}

	// Create a ticket object
	ticket := models.Ticket{
		ID:      primitive.NewObjectID(), // Generate a new ObjectID for the ticket
		UserID:  userID,
		TrainID: trainID,
		Seat:    requestData.Seat,
		Status:  "Booked", // Default status
	}

	collection := database.DB.Database(os.Getenv("DB_NAME")).Collection("tickets")
	_, err = collection.InsertOne(context.TODO(), ticket)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error booking ticket"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ticket booked successfully"})

}

func CancelTicket(c *gin.Context) {

	var body struct {
		TicketID string `json:"ticket_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketID, err := primitive.ObjectIDFromHex(body.TicketID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	collection := database.DB.Database(os.Getenv("DB_NAME")).Collection("tickets")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": ticketID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error cancelling ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket cancelled successfully"})
}
