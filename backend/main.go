package main

import (
	"context"
	"log"

	"dedalusrpg.com/ent"
	"github.com/facebook/ent/examples/start/ent/migrate"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=dedalus password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
