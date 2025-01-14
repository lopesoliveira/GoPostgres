package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/lopesoliveira/GoPostgres/config"
	"github.com/lopesoliveira/GoPostgres/controller"
	_ "github.com/lopesoliveira/GoPostgres/docs"
	"github.com/lopesoliveira/GoPostgres/helper"
	"github.com/lopesoliveira/GoPostgres/model"
	"github.com/lopesoliveira/GoPostgres/repository"
	"github.com/lopesoliveira/GoPostgres/router"
	"github.com/lopesoliveira/GoPostgres/service"
	"net/http"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	/*log.Info().Msg("Started Server!")*/
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

/*
package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/lopesoliveira/GoPostgres/config"
	"github.com/lopesoliveira/GoPostgres/controller"
	"github.com/lopesoliveira/GoPostgres/helper"
	"github.com/lopesoliveira/GoPostgres/model"
	"github.com/lopesoliveira/GoPostgres/repository"
	"github.com/lopesoliveira/GoPostgres/router"
	"github.com/lopesoliveira/GoPostgres/service"
	"net/http"
	"time"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//Init Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	// Init Service
	tagService := service.NewTagsServiceImpl(tagRepository, validate)

	//Init Controller
	tagController := controller.NewTagController(tagService)

	//Router
	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
*/

/*
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lopesoliveira/GoPostgres/config"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {

	var dataBase *gorm.DB = config.DatabaseConnection()
	fmt.Println(dataBase.Name())

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
*/
