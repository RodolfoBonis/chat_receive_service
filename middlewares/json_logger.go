package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"chat_receive_service/utils"
)

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		go verifyTimeLogFile()
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := utils.GetDurationInMillseconds(start)

		entry := log.Fields{
			"date":        time.Now().Format("2006-01-02 15:04:05"),
			"client_ip":   utils.GetClientIP(c),
			"duration":    fmt.Sprintf("%f", duration) + "ms",
			"method":      c.Request.Method,
			"path":        c.Request.RequestURI,
			"status":      c.Writer.Status(),
			"user_id":     utils.GetUserID(c),
			"referrer":    c.Request.Referer(),
			"request_id":  c.Writer.Header().Get("Request-Id"),
			"api_version": utils.GetApiVersion(),
		}

		if c.Writer.Status() >= 300 {

			date := time.Now().Format("2006-01-02_15:04:05")
			f, _ := os.Create("logs/" + utils.GetProgramName() + "-request-" + date + ".log")

			jsonString, _ := json.Marshal(entry)
			_, err := io.MultiWriter(f).Write(jsonString)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func verifyTimeLogFile() {
	dirEntrys, _ := os.ReadDir("logs")

	for _, dirEntry := range dirEntrys {
		info, _ := dirEntry.Info()
		if info.ModTime().Add((24 * time.Hour) * 30).Before(time.Now()) {
			err := os.Remove("logs/" + dirEntry.Name())
			if err != nil {
				log.Error(err)
			}
		}
	}
}
