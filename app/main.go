package main

import (
	"clean-arch/infrastructure/http/server"
	"os"
	"time"
)

func main() {
	os.Setenv("APP_NAME", "clean-arch")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("AWS_ENDPOINT", "http://localhost:4566")
	os.Setenv("AWS_REGION", "sa-east-1")

	var app = server.NewConfig().
		Name(os.Getenv("APP_NAME")).
		WebServerPort(os.Getenv("SERVER_PORT")).
		ContextTimeout(10 * time.Second).
		AwsEndpoint(os.Getenv("AWS_ENDPOINT")).
		AwsRegion(os.Getenv("AWS_REGION")).
		Logger()

	app.NewDB().WebServer().Start()
}
