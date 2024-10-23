package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gustazevedo/xlike_go/adapters/database"
    "github.com/gustazevedo/xlike_go/adapters/web"
    "github.com/gustazevedo/xlike_go/domain/service"
    "log"
)

func main() {
    db, err := database.NewPostgresConnection()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    userRepository := database.NewPostgresUserRepository(db)
    userService := service.NewUserService(userRepository)
    userHandler := web.NewUserHandler(userService)

    router := gin.Default()

    router.POST("/register", userHandler.RegisterUser)

    router.Run() // listen and serve on 0.0.0.0:8080
}
