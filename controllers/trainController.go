package controllers

import (
	"context"
	"net/http"
	"os"
	"trainTicketsGo/database"
	"trainTicketsGo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTrains(c *gin.Context) {
	// from the database package, we are using the (DB) object created by me,  to connect to the database
	// and then we are using the (Database) method to connect to the database named "trainTicketGo"
	// and then we are using the (Collection) method to connect to the collection named "trains"
	collection := database.DB.Database(os.Getenv("DB_NAME")).Collection("trains")
	cursor, err := collection.Find(context.TODO(), bson.M{}) // #TODO what is context.TODO() ? and bson.M{} ?

	if err != nil { // If error is not nil, it means there was an error while fetching the data
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching trains "}) // #TODO what is gin.H ?
		return                                                                           // Exit the function
	}

	var trains []models.Train                                   // Create a slice(vector) of Train struct data type
	if err := cursor.All(context.TODO(), &trains); err != nil { // This will iterate over all the documents in the cursor and decode them into the trains slice and return an error if there is any . Cursor is like a pointer to the result set of a query.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding  trains"})
		return
	}

	// If there are no errors, return the trains slice as a JSON response and status OK
	c.JSON(http.StatusOK, trains)

}

func AddTrain(c *gin.Context) { // 'c' is the context object and it is similar to the 'request' object in Express.js
	var train models.Train // variable train of type Train struct
	if err := c.ShouldBindJSON(&train); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.DB.Database("trainTicketGo").Collection("trains") // Collection is like a table in SQL databases
	_, err := collection.InsertOne(context.TODO(), train)                    // Insert the train object into the collection

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding train"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Train added successfully"}) // Return a success message if the train is added successfully

}
