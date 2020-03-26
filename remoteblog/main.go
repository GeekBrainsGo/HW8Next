// Package remoteblog implements blog server with mongodb,
// swagger, tests, logging and configuration.
package main

/*
	Basics Go.
	Rishat Ishbulatov, dated Oct 09, 2019.
	1. Connect logging to your project.
	2. Select the parameters that should be moved to the configuration file
	(at your discretion). Write the code to download the configuration
	and use the data in the project.
	3. Lift up your server in VDS.
	It can be Google AppEngine, DigitalOcean, Amazon, etc. In particular, you
	can take VDS from any host (Fozzy, GoodDaddy, etc.).
	Google AppEngine will allow you to spin a good car for the first year for free.
	4. * Get a domain name from any registrar (Reg.ru or directly from your hoster).
	Bind it to the IP address of your VDS.
	The cost is different, some hosters give a domain name for
	free when buying VDS from them.
	AppEngine also gives a free domain in the * .appspot.com zone
*/

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"remoteblog/server"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Mongoblog server
// @version 1.0
// @description Blog server with swagger and tests for handlers.
// @contact.name Rishat Ishbulatov
// @contact.email progjb@gmail.com
// @license.name MIT
// @host localhost:8080
// @BasePath /

func main() { // 6
	flagConfigPath := flag.String("c", "./config.yaml", "yaml config file path")
	flag.Parse()

	conf, err := ReadConfig(*flagConfigPath)
	if err != nil {
		panic(fmt.Sprintf("can't read config: %s", err))
	}

	lg, err := ConfigureLogger(&conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("can't configure logger: %s", err))
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		lg.WithError(err).Fatal("can't get new client")
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		lg.WithError(err).Fatal("can't connect to db")
	}

	if err = client.Ping(ctx, nil); err != nil {
		lg.WithError(err).Fatal("can't ping to db")
	}

	db := client.Database(conf.Server.DBname)
	defer client.Disconnect(ctx)

	serv := server.New(ctx, lg, &conf.Server, db)
	serv.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	serv.Stop()
}
