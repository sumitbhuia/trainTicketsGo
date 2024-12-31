package models

// models package contains the struct definitions for the data models used in the application
// User struct defines the structure of a ticket object
import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` //#TODO : what is bson:"_id,omitempty"? Answer :
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
