package main

import (
	"log"

	"github.com/kryast/Crud-1.git/database"
	"github.com/kryast/Crud-1.git/handlers"
	"github.com/kryast/Crud-1.git/models"
	"github.com/kryast/Crud-1.git/repositories"
	"github.com/kryast/Crud-1.git/routers"
	"github.com/kryast/Crud-1.git/services"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	db.AutoMigrate(&models.User{})

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := routers.SetupRouter(userHandler)
	r.Run(":8080")
}
