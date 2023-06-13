package router

import (
	"clean-arch/adapter/api/controller"
	"clean-arch/adapter/presenter"
	"clean-arch/adapter/repository"
	"clean-arch/core/usecase"
	"clean-arch/infrastructure/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Port int64

	Server interface {
		Listen()
	}

	ginEngine struct {
		router     *gin.Engine
		port       Port
		ctxTimeout time.Duration
		db         repository.NoSQL
	}
)

func NewGinServer(port Port, db repository.NoSQL, timeout time.Duration) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		port:       port,
		db:         db,
		ctxTimeout: timeout,
	}
}

func (engine ginEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	engine.setAppHandlers(engine.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", engine.port),
		Handler:      engine.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.WithError(err).Fatalln("Server shutdown Failed")
	}
}

func (engine ginEngine) setAppHandlers(router *gin.Engine) {
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, nil) })
	router.GET("/v1/users/:userId", engine.handleGetUserById())
	router.POST("/v1/users", engine.handleCreateUser())
	router.PUT("/v1/users/:userId", engine.handleUpdateUser())
	router.DELETE("/v1/users/:userId", engine.handleDeleteUser())
}

func (engine ginEngine) handleGetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()

		usecase := usecase.NewGetUserUseCase(
			repository.NewUserRepositoryDynamoDB(engine.db),
			presenter.NewGetUserPresenter(),
			engine.ctxTimeout,
		)

		userController := controller.NewGetUserController(usecase)
		userController.Execute(ctx.Writer, ctx.Request)
	}
}

func (engine ginEngine) handleCreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usecase := usecase.NewCreateUserUseCase(
			repository.NewUserRepositoryDynamoDB(engine.db),
			presenter.NewCreateUserPresenter(),
			engine.ctxTimeout,
		)

		userController := controller.NewCreateUserController(usecase)
		userController.Execute(ctx.Writer, ctx.Request)
	}
}

func (engine ginEngine) handleUpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()

		usecase := usecase.NewUpdateUserUseCase(
			repository.NewUserRepositoryDynamoDB(engine.db),
			presenter.NewUpdateUserPresenter(),
			engine.ctxTimeout,
		)

		userController := controller.NewUpdateUserController(usecase)
		userController.Execute(ctx.Writer, ctx.Request)
	}
}

func (engine ginEngine) handleDeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()

		usecase := usecase.NewDeleteUserUseCase(
			repository.NewUserRepositoryDynamoDB(engine.db),
			presenter.NewDeleteUserPresenter(),
			engine.ctxTimeout,
		)

		userController := controller.NewDeleteUserController(usecase)
		userController.Execute(ctx.Writer, ctx.Request)
	}
}
