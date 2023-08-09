package handler

import (
	"NDerDB/server/storage"
	"github.com/gin-gonic/gin"
	"io"
)

func SaveStringHandler(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(500)
		return
	}

	fileName, err := storage.SaveString(data)
	if err != nil {
		c.Status(500)
		return
	}

	c.String(200, fileName)
}

func GetStringHandler(c *gin.Context) {
	fileName, ok := c.GetQuery("id")
	if fileName == "" || !ok {
		c.Status(400)

		return
	}

	data, err := storage.GetString(fileName)
	if err != nil {
		c.Status(500)
		return
	}

	c.String(200, string(data))
}
