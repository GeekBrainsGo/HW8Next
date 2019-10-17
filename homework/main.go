package main

// @title HW7
// @version 0.1
// @description This is a HW blog

// @contact.name Dmitrii Fadeev

import (
	"context"
	"serv/models"
	"serv/server"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const staticDir = "www/static"

func main() {
	lg := NewLogger()
	var conf Config
	if err := conf.ReadConfig("config.yaml"); err != nil {
		lg.Panic("Can't read config: ", err)
	}

	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.Server.MongoDB))
	if err != nil {
		lg.Panic("Can't connect to DB", err)
	} else {
		lg.Info("Connection to DB successful")
	}

	_ = client.Connect(ctx)
	db := client.Database("blogs")

	blog := &models.Blog{
		Title:    "test",
		Contents: "test",
	}

	_, err = blog.Insert(ctx, db)
	if err != nil {
		lg.Fatal(err)
	}

	srv := server.New(lg, db, staticDir)
	srv.Start(conf.Server.Addr)
}
