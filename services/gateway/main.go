package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/luizpbraga/microMP3/services/gateway/src/database/connector"
	"github.com/luizpbraga/microMP3/services/gateway/src/routes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main0() {
	// TODO: GridFS
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := connector.InitConnectionToMongo(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database("videos")
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := bucket.Drop(); err != nil {
			panic(err)
		}
	}()

	if file, err := os.Open("./file.txt"); err != nil {
		defer file.Close()

		uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "fist"}})
		objectID, err := bucket.UploadFromStream("file.txt", io.Reader(file),
			uploadOpts)
		if err != nil {
			panic(err)
		}
		log.Printf("New file uploaded with ID %s", objectID)
	}
}

func main() {
	routes.LoadRoutes()
	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
