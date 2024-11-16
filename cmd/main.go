package main

import (
	"fmt"

	"github.com/Jefschlarski/golang-modular/modules"
	"github.com/Jefschlarski/golang-modular/modules/task"
	"github.com/Jefschlarski/golang-modular/modules/user"
	"github.com/Jefschlarski/golang-modular/pkg/config"
	"github.com/Jefschlarski/golang-modular/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {

	err := config.Load()
	if err != nil {
		panic(err)
	}

	dbconfig := config.GetDbConfig()

	// Initialize gin server engine
	server := gin.Default()

	// Initialize database
	db, err := db.ConnectDB(dbconfig)
	if err != nil {
		panic(err)
	}

	// Modules
	modules := []modules.ModuleInterface{
		user.NewUserModule(),
		task.NewTaskModule(),
	}

	v1 := server.Group("/api/v1")

	// Initialize modules
	for _, module := range modules {
		module.Init(v1, db)
	}

	// Ping route
	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	apiConfig := config.GetApiConfig()

	server.Run(fmt.Sprintf("%s:%d", apiConfig.Url, apiConfig.Port))
}
