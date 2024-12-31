package models

// models package contains the struct definitions for the data models used in the application
// Ticket struct defines the structure of a ticket object
import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"user_id"`
	TrainID primitive.ObjectID `bson:"train_id"`
	Seat    string             `bson:"seat"`
	Status  string             `bson:"status"`
}
