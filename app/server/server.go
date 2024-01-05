package server

import (
	"context"

	"github.com/gin-gonic/gin"
	cartApp "github.com/o-ga09/tutorial-ec-backend/app/application/cart"
	orderApp "github.com/o-ga09/tutorial-ec-backend/app/application/order"
	productApp "github.com/o-ga09/tutorial-ec-backend/app/application/product"
	userApp "github.com/o-ga09/tutorial-ec-backend/app/application/user"
	"github.com/o-ga09/tutorial-ec-backend/app/config"
	orderDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/order"
	"github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql"
	queryservice "github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/query_service"
	"github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/repository"
	"github.com/o-ga09/tutorial-ec-backend/app/infrastructure/redis"
	redisRepo "github.com/o-ga09/tutorial-ec-backend/app/infrastructure/redis/repository"
	cartHandler "github.com/o-ga09/tutorial-ec-backend/app/presentation/cart"
	"github.com/o-ga09/tutorial-ec-backend/app/presentation/health"
	orderHandler "github.com/o-ga09/tutorial-ec-backend/app/presentation/order"
	productsHandler "github.com/o-ga09/tutorial-ec-backend/app/presentation/products"
	userHandler "github.com/o-ga09/tutorial-ec-backend/app/presentation/user"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
)

func NewServer() (*gin.Engine, error) {
	r := gin.New()
	cfg := config.GetConfig()
	if cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	// setting logger
	logger := middleware.New()
	httpLogger := middleware.RequestLogger(logger)

	//setting a CORS
	cors := middleware.CORS()

	r.Use(cors)
	r.Use(httpLogger)

	db := mysql.New(context.Background())
	redis := redis.NewRedisClient(context.Background())
	systemHandler := health.NewHandler()
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	fetchQueryService := queryservice.NewproductQueryService(db)
	orderRepo := repository.NewOrderRepository(db)
	cartRepo := redisRepo.NewCartRepository(redis)
	userHandler := userHandler.NewHandler(userApp.NewFindUserUsecase(userRepo),userApp.NewFindAllUsersUseCase(userRepo),userApp.NewSaveUserUsecase(userRepo),userApp.NewDeleteUserUsecase(userRepo))
	productHandler := productsHandler.NewHandler(productApp.NewSaveProductUsecase(productRepo),productApp.NewFetchProductUseCase(fetchQueryService))
	orderHandler := orderHandler.NewHandlre(orderApp.NewSaveOrderUseCase(orderDomain.NewOrderDomainService(orderRepo,productRepo),redisRepo.NewCartRepository(redis)))
	cartHandler := cartHandler.NewHandler(cartApp.NewAddCartUsecase(cartRepo,productRepo))

	v1 := r.Group("/v1")
	users := v1.Group("/users")
	order := v1.Group("/order")
	cart := v1.Group("/cart")
	products := v1.Group("/products")
	{
		v1.GET("/health", systemHandler.Health)
		users.GET("/:id",userHandler.GetUserById)
		users.GET("",userHandler.GetUsers)
		users.POST("",userHandler.CreateUser)
		users.DELETE("/:id",userHandler.DeleteUser)
		products.GET("",productHandler.GetProducts)
		products.POST("",productHandler.PostProducts)
		order.POST("",orderHandler.PostOrders)
		cart.POST("",cartHandler.PostCart)
	}


	return r, nil
}