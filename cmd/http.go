package cmd

import (
	"mini-online-store/internal/handlers"
	"mini-online-store/internal/repository"
	"mini-online-store/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouters(r *gin.Engine, db *gorm.DB) {
	// repo
	userRepo := repository.NewUserRepository(db)
	userTokenRepo := repository.NewUserTokenRepository(db)
	productRepo := repository.NewProductRepository(db)
	shoppingCartRepo := repository.NewShoppingCartRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// usecase
	userUsecase := usecase.NewUserUsecase(&userRepo, &userTokenRepo)
	productUsecase := usecase.NewProductUsecase(&productRepo)
	shoppingCartUsecase := usecase.NewShoppingCartUsecase(&shoppingCartRepo, &productRepo)
	transactionUsecase := usecase.NewTransactionUsecase(&transactionRepo, &productRepo)

	// handler
	olshopHandler := handlers.NewOlshopHandler(userUsecase, productUsecase, shoppingCartUsecase, transactionUsecase)
	olshopHandler.SetupHandler(r)
}
