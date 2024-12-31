package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client // DB is the database object and is exported so that it can be used in other packages

func ConnectDB() {
	err := godotenv.Load() // Load the .env file
	if err != nil {
		log.Fatal("Error loading .env file", err)
		fmt.Println("Error loading .env file")
	}

	log.Println("env file loaded")
	/*
	   Q: What is clientOptions ?
	   A: clientOptions is a variable that stores the client options for connecting to the MongoDB database.

	   Q: What are the default values of clientOptions ?
	   A: The default values of clientOptions are the default values for the ClientOptions struct, which include the default URI for connecting to the MongoDB database.
	*/
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")) // #TODO : what is options.Client() and ApplyURI() ? and cna i store thi uRI in .env ?

	// Create a new client to later export it to other packages , to be used to connect to the database
	mongoclient, err := mongo.Connect(context.TODO(), clientOptions) // Connect to the MongoDB database
	/*
	   Q: What is the differnce between context.Background() and context.TODO() ?
	   A:  context.TODO() returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).
	       context.Background() returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

	   Q: When to use context.TODO() ?
	   A:  context.TODO() is used when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).

	   Q: When to use context.Background() ?
	   A:  context.Background() is used by the main function, initialization, and tests, and as the top-level Context for incoming requests.


	*/
	if err != nil {
		fmt.Println("Error connecting to MongoDB", err)
		return
	}

	// Ping the database to check if it is connected
	err = mongoclient.Ping(context.Background(), readpref.Primary())
	/*
	   Q: What is a primary node in a replica set ?
	   A:  A primary node in a replica set is the main node that receives all write operations and replicates the data to secondary nodes. The primary node is responsible for coordinating the replication process and handling failover events in the replica set.

	   Q: What is a replica set in MongoDB ?
	   A:  A replica set in MongoDB is a group of MongoDB servers that maintain the same data set. Replica sets provide redundancy and high availability, allowing the system to continue operating even if one or more nodes fail. Replica sets also support read scaling by distributing read operations across multiple nodes.

	   Q: What is readpref.Primary() ?
	   A: readpref.Primary() returns a read preference that prefers the primary node of a replica set. This mode is the default for most drivers, and will send all read operations to the current primary by default.

	   Q: Why is nil passed as the first argument to the Ping() method ?
	   A:  The first argument to the Ping() method is the context object. The context object is used to pass the request deadline, cancellation signals, and other request-scoped values across API boundaries and between processes. In this case, the context.TODO() function returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).

	   Q: What is the differnce in output if nil is used instead of readpref.Primary() ?
	   A:  If nil is used instead of readpref.Primary(), the output will be "Connected to MongoDB!" because the Ping() method will not check the read preference of the database.
	       If readpref.Primary() is used, the output will be "Connected to MongoDB!" because the Ping() method will check the read preference of the database and return an error if the primary node is not available.


	*/

	if err != nil {
		fmt.Println("MongoDb Ping failed : ", err)
		return
	}

	fmt.Println("Connected to MongoDB!")
	DB = mongoclient
}
