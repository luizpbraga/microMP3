package main

import (
	"context"
	"log"
	"time"

	"github.com/luizpbraga/microMP3/services/gateway/src/database/connector"
)

func main() {
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

}
