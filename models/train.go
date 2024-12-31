package models

// models package contains the struct definitions for the data models used in the application
// Train struct defines the structure of a ticket object
import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Source      string             `bson:"source"`
	Destination string             `bson:"destination"`
	Schedule    string             `bson:"schedule"`
}
