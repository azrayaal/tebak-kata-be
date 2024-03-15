package main

import (
	"context"
	"fmt"

	mongodb "example.com/web-service-gin/db"

	"example.com/web-service-gin/controllers/avatarControllers"
	"example.com/web-service-gin/controllers/diamondControllers"
	"example.com/web-service-gin/controllers/quizControllers"
	"example.com/web-service-gin/controllers/userControllers"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var collection *mongo.Collection

type Config struct {
	NGROKURL string `yaml:"ngrok_url"`
}

func main() {

	client, err := mongodb.ConnectToMongoDB()
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

	avatarCollection := client.Database("teka_apps").Collection("avatar")
	quizcCollection := client.Database("teka_apps").Collection("quiz")
	userCollection := client.Database("teka_apps").Collection("user")
	diamondCollection := client.Database("teka_apps").Collection("diamond")

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("avatarCollection", avatarCollection)
		c.Set("quizCollection", quizcCollection)
		c.Set("userCollection", userCollection)
		c.Set("diamondCollection", diamondCollection)
		c.Next()
	})

	router.GET("/api/v1/avatar", avatarControllers.GetAllAvatars)
	router.GET("/api/v1/avatar/:id", avatarControllers.GetDetailAvatar)

	router.GET("/api/v1/quiz", quizControllers.GetAllquizes)
	router.GET("/api/v1/quiz/:id", quizControllers.GetDetailQuiz)

	router.GET("/api/v1/users", userControllers.GetAllUsers)
	router.GET("/api/v1/users/:id", userControllers.GetDetailUser)

	router.GET("/api/v1/diamond", diamondControllers.GetAllDiamonds)
	router.GET("/api/v1/diamond/:id", diamondControllers.GetDetailDiamond)

	router.Run()

	// Construct ngrok URL with endpoint
	// ngrokURL := config.NGROKURL
	// if ngrokURL != "" {
	// 	ngrokURL += "/api/v1"
	// } else {
	// 	ngrokURL = "http://localhost:8080/api/v1" // Default to localhost
	// }

	// fmt.Println("Your ngrok endpoint:", ngrokURL)

	// router.Run(":8080") // Assuming your application runs on port 8080 locally
}

// func loadConfig(filename string) (Config, error) {
// 	var config Config
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return config, err
// 	}
// 	defer file.Close()

// 	decoder := yaml.NewDecoder(file)
// 	err = decoder.Decode(&config)
// 	return config, err
// }

// func run(ctx context.Context) error {
// 	listener, err := ngrok.Listen(ctx,
// 		config.HTTPEndpoint(),
// 		ngrok.WithAuthtokenFromEnv(),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("App URL", listener.URL())
// 	return http.Serve(listener, http.HandlerFunc(handler))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "<h1>Hello from ngrok-go!</h1>")
// }

// func run(ctx context.Context) error {
// 	// Listen using ngrok
// 	listener, err := ngrok.Listen(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	// Print the ngrok URL
// 	log.Println("App URL:", listener.URL())

// 	// Serve HTTP using the ngrok listener
// 	return http.Serve(listener, http.HandlerFunc(handler))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Respond with a simple message
// 	fmt.Fprintln(w, "<h1>Hello from ngrok-go!</h1>")
// }

// func main() {
// 	// Set up a context
// 	ctx := context.Background()

// 	// Run the server
// 	if err := run(ctx); err != nil {
// 		log.Fatal("Error:", err)
// 	}
// }
