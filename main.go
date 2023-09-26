package main

import (
	"context"
	"fmt"
	"log"
	"quoation-backend/controllers"
	"quoation-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
var (
	server          *gin.Engine
	quoteservice    services.QuoteService
	quotecontroller controllers.QuoteController
	ctx             context.Context
	quotecollection *mongo.Collection
	mongoclient     *mongo.Client
	err             error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://kirangowdan34:r8zrgL71zLi35Pym@cluster0.c6y20gu.mongodb.net/")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")
	quotecollection = mongoclient.Database("quotesdb").Collection("quotes")
	quoteservice = services.NewQuoteService(quotecollection,ctx)
	quotecontroller = controllers.New(quoteservice)
	


}
// v1/quote/create
func main(){
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000","*"} // Allow all origins
	config.AllowMethods = []string{"GET", "POST", "PATCH","PUT","DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization","Access-Control-Allow-Origin"}
	server = gin.Default()
	server.Use(cors.New(config))

	
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	quotecontroller.RegisterQuoteRoutes(basepath)
	log.Fatal(server.Run(":9090"))
	
    
}