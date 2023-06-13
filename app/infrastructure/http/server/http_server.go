package server

import (
	"clean-arch/adapter/repository"
	"clean-arch/infrastructure/database"
	"clean-arch/infrastructure/http/router"
	"clean-arch/infrastructure/logger"
	"strconv"
	"time"
)

type config struct {
	appName       string
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
	db            repository.NoSQL
	awsRegion     string
	awsEndpoint   string
}

func NewConfig() *config {
	return &config{}
}

func (config *config) Name(appName string) *config {
	config.appName = appName
	return config
}

func (config *config) ContextTimeout(timeout time.Duration) *config {
	config.ctxTimeout = timeout
	return config
}

func (config *config) AwsRegion(awsRegion string) *config {
	config.awsRegion = awsRegion
	return config
}

func (config *config) AwsEndpoint(awsEndpoint string) *config {
	config.awsEndpoint = awsEndpoint
	return config
}

func (config *config) WebServerPort(port string) *config {
	intPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		logger.Fatal(err)
	}
	config.webServerPort = router.Port(intPort)
	return config
}

func (config *config) Logger() *config {
	logger.NewZapLogger()
	logger.Infof("Successfully configured log")
	return config
}

func (config *config) WebServer() *config {
	server := router.NewGinServer(config.webServerPort, config.db, config.ctxTimeout)
	logger.Infof("Successfully configured router server")

	config.webServer = server
	return config
}

func (config *config) NewDB() *config {
	config.db = database.NewDynamoDBClient(config.awsRegion, config.awsEndpoint)
	logger.Infof("Successfully configured db")
	return config
}

func (config *config) Start() {
	config.webServer.Listen()
}
