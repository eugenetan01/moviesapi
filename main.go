package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	Title string `bson:"title"`
	Plot  string `bson:"plot"`
	Imdb  Imdb   `bson:"imdb"`
}

type Imdb struct {
	Rating interface{} `bson:"rating"`
	Votes  interface{} `bson:"votes"`
}

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {

	router := gin.Default()

	// Connect to MongoDB
	mongoURI := os.Getenv("MONGODB_EXT")
	print(mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	collection := client.Database("sample_mflix").Collection("movies")

	router.GET("/query", func(c *gin.Context) {
		title := c.Query("title")
		var result Movie
		err := collection.FindOne(context.Background(), bson.M{"title": title}).Decode(&result)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, result)
	})

	router.GET("/search", func(c *gin.Context) {
		keyword := c.Query("keyword")

		pipeline := mongo.Pipeline{
			bson.D{
				{"$search", bson.D{
					{"text", bson.D{
						{"query", keyword},
						{"path", bson.D{{"wildcard", "*"}}},
						{"fuzzy", bson.D{}},
					}},
				}},
			},
		}

		cursor, err := collection.Aggregate(context.Background(), pipeline)
		if err != nil {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
		}
		if err != nil {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
			return
		}
		var results []Movie
		if err = cursor.All(context.Background(), &results); err != nil {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(200, results)
	})

	router.Run(":8080")
}
