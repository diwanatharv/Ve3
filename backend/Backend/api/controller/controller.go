package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log1 "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"log"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile("C:\\Users\\athar\\GolandProjects\\awesomeProject\\.env")
	err := viper.ReadInConfig()
	if err != nil {
		log1.Error(err.Error())
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log1.Error(err.Error())
		log.Fatalf("Invalid type assertion")
	}

	return value
}
func basicAuthValidator(username string, password string, c echo.Context) (bool, error) {
	viperenv1 := viperEnvVariable("USERNAME")
	viperenv2 := viperEnvVariable("PASSWORD")

	if username == viperenv1 && password == viperenv2 {
		return true, nil
	} else {
		return false, errors.New("password and username do not match")
	}
}
func Makeroutes(e *echo.Echo) {
	e.Use(middleware.BasicAuth(basicAuthValidator))
	e.POST("/task", createtask)
	e.GET("/task/:id", gettask)
	e.GET("/task", getalltask)
	e.PUT("/task/:id", updatetask)
	e.DELETE("/task/:id", deletetask)
}
