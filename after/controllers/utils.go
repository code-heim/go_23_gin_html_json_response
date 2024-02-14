package controllers

import "github.com/gin-gonic/gin"

type respFormat int

const (
	formatHTML respFormat = iota
	formatJSON
)

var formatMap = map[string]respFormat{
	"json": formatJSON,
	"html": formatHTML,
}

func getFormat(c *gin.Context) respFormat {
	format := c.GetHeader("X-Original-Extension")
	if format == "" {
		format = "html"
	}
	return formatMap[format]
}
