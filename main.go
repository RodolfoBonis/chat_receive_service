package main

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"chat_receive_service/domain/usecases"
	"chat_receive_service/middlewares"
	"chat_receive_service/utils"
)

func main() {
	utils.LoadEnvVars()
	utils.UseJSONLogFormat()

	r := gin.New()

	r.Use(middlewares.JSONLogMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())

	listenMessageUc := usecases.ListenMessagesUC{}

	listenMessageUc.ListenMessages()

	port := utils.GetEnv("PORT", "8080")
	err := r.Run(":" + port)

	if err != nil {
		date := time.Now().Format("2006-01-02|15:04:05")
		f, _ := os.Create("logs/" + utils.GetProgramName() + "-server-" + date + ".log")

		_, err := io.MultiWriter(f).Write([]byte(err.Error()))

		if err != nil {
			log.Error(err)
		}
	}
}
