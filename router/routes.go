package router

import (
	//"gin-inject/common/cache"
	//"gin-inject/common/datasource"
	//"gin-inject/common/rabbitmq"
	"gin-inject/common/datasource"
	"gin-inject/controller"
	"gin-inject/repository"
	"gin-inject/service"
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

func Configure(app *gin.Engine) {
	//controller declare
	var index controller.Index

	//inject declare
	db := datasource.Db{}
	//redis := cache.Redis{}
	//rabbit := rabbitmq.Mq{}

	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &index},
		&inject.Object{Value: &db},
		//&inject.Object{Value: &redis},
		//&inject.Object{Value: &rabbit},
		&inject.Object{Value: &repository.StartRepo{}},
		&inject.Object{Value: &service.StartService{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	//database connect
	err = db.Connect()
	if err != nil {
		log.Fatal("db fatal:", err)
	}
	// //redis server connect
	// err = redis.Connect()
	// if err != nil {
	// 	log.Fatal("redis fatal:", err)
	// }
	// // rabbitmq server connect
	// err = rabbit.Connect()
	// if err != nil {
	// 	log.Fatal("RabbitMQ fatal:", err)
	// }

	v1 := app.Group("/")
	{
		v1.GET("/get/:msg", index.GetName)
	}
}
