package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"serv/logger"
	"serv/server"

	_ "github.com/go-sql-driver/MySQL"
)

// @title Geekbrains HW7 Server
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {

	conf, err := ReadConfig("./config.yaml")
	if err != nil {
		panic(fmt.Sprintf("can't read config: %s", err))
	}

	fmt.Println(conf)

	flagRootDir := flag.String("rootdir", conf.Logger.Rootdir, "root dir of the server")
	flagServAddr := flag.String("addr", conf.Logger.Addr, "server address")

	flag.Parse()

	lg, err := logger.ConfigureLogger(&conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("can't configure logger: %s", err))
	}

	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	_ = client.Connect(ctx)

	db := client.Database("geekbrains")
	serv := server.New(lg, *flagRootDir, ctx, db)

	go func() {
		err := serv.Start(*flagServAddr)
		if err != nil {
			lg.WithError(err).Fatal("can't run the server")
		}
	}()

	stopSig := make(chan os.Signal)
	signal.Notify(stopSig, os.Interrupt, os.Kill)
	<-stopSig

}
