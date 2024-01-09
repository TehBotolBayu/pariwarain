package router

import (
    "github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/TehBotolBayu/pariwarainAPI/controllers"
    "github.com/TehBotolBayu/pariwarainAPI/models"
)

func route() {
    // Initialize GORM DB connection
    db, err := model.

    // Initialize Gin router
    router := gin.Default()

    // Initialize controllers
    propertiController := controllers.NewPropertiController(db)
    PenggunaController := controllers.NewPenggunaController(db)

    // Routes for Properti controller
    propertiRoutes := router.Group("/properties")
    {
        propertiRoutes.GET("/", propertiController.GetAllProperties)
        propertiRoutes.POST("/", propertiController.CreateProperty)
        propertiRoutes.PUT("/:id", propertiController.UpdateProperty)
        propertiRoutes.DELETE("/:id", propertiController.DeleteProperty)
    }

    // Routes for Pengguna controller
    PenggunaRoutes := router.Group("/Penggunas")
    {
        PenggunaRoutes.GET("/", PenggunaController.GetAllPengguna)
        PenggunaRoutes.POST("/", PenggunaController.CreatePengguna)
        PenggunaRoutes.PUT("/:id", PenggunaController.UpdatePengguna)
        PenggunaRoutes.DELETE("/:id", PenggunaController.DeletePengguna)
    }

    // Start the server
    router.Run(":8080")
}