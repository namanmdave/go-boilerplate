package http

import (
	"go-boilerplate/config"
	"go-boilerplate/server/http/handler"
	"go-boilerplate/store"
	db "go-boilerplate/store/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type HTTPServer struct {
	router         *gin.Engine
	store          *store.Store
	addr           string
	logger         *zap.SugaredLogger
	validator      *validator.Validate
	handlerFactory *handler.HandlerFactory
}

func InitHTTPServer() *HTTPServer {
	dbConfig := config.GetDBConfig()
	redisConfig := config.GetRedisConfig()

	postgresDB, err := store.InitPostgres(dbConfig)
	if err != nil {
		panic(err)
	}

	redisClient, err := store.InitRedis(redisConfig)
	if err != nil {
		panic(err)
	}

	querier := db.New(postgresDB)

	appStore := store.NewStoreBuilder().
		WithPostgres(postgresDB).
		WithRedis(redisClient).
		WithQuerier(querier).
		Build()

	serverConfig := config.GetServerConfig()

	logger, _ := zap.NewProduction()
	validator := validator.New()
	router := gin.Default()

	handlerFactory := handler.NewHandlerFactory(appStore, logger.Sugar(), validator)

	return &HTTPServer{
		router:         router,
		store:          appStore,
		logger:         logger.Sugar(),
		validator:      validator,
		addr:           serverConfig.HTTPAddr,
		handlerFactory: handlerFactory,
	}
}

func (s *HTTPServer) Start() error {
	s.registerRoutes()
	return s.router.Run(s.addr)
}

func (s *HTTPServer) registerRoutes() {
	s.router.Use(gin.Recovery())
	s.router.Use(RequestLoggingMiddleware())
	// s.router.Use(CORSMiddleware())

	s.router.OPTIONS("/*any", handler.OptionsHandler)

	// Public routes
	public := s.router.Group("/api")
	{
		public.GET("/health", handler.HealthCheckHandler)
	}

	// Protected routes
	protected := s.router.Group("/api")
	protected.Use(AuthMiddleware())
	{
		userHandler := s.handlerFactory.GetUserHandler()
		users := protected.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
		}

		groupHandler := s.handlerFactory.GetGroupHandler()
		groups := protected.Group("/groups")
		{
			groups.GET("/:group_id", groupHandler.UpgradeHandler)
		}

		messageHandler := s.handlerFactory.GetMessageHandler()
		messages := protected.Group("/messages")
		{
			messages.GET("/:group_id", messageHandler.FetchGroupMessages)
		}
	}
}
