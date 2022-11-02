
package main


import(
	routes "github.com/lordscoba/golang-auth/pkg/routes"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-1"})
	})

	router.GET("/api-2", func (c *gin.Context){
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}









// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func main() {
// 	ctx := context.TODO()
//   // uri := "mongodb+srv://lordscoba.biwotjy.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=<path_to_certificate>"
//   uri := "mongodb+srv://lordscoba:blockchain@lordscoba.biwotjy.mongodb.net/?retryWrites=true&w=majority"
//   // uri := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.0"
//   serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
//   clientOptions := options.Client().
//       ApplyURI(uri).
//       SetServerAPIOptions(serverAPIOptions)

// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil { log.Fatal(err) }
// 	defer client.Disconnect(ctx)

// 	collection := client.Database("testDB").Collection("testCol")
// 	docCount, err := collection.CountDocuments(ctx, bson.D{})
// 	if err != nil { log.Fatal(err) }
// 	fmt.Println(docCount)
// 	fmt.Println("Successful")
// }
