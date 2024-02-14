package controllers

import (
	"fmt"
	"gin_blogs/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BlogsIndex(c *gin.Context) {
	blogs := models.BlogsAll(c)
	switch getFormat(c) {
	case formatHTML:
		c.HTML(
			http.StatusOK,
			"blogs/index.tpl",
			gin.H{
				"blogs":      blogs,
				"page":       c.GetInt("page"),
				"pageSize":   c.GetInt("pageSize"),
				"totalPages": c.GetInt("totalPages"),
			},
		)
	case formatJSON:
		c.JSON(
			http.StatusOK,
			gin.H{
				"blogs":      blogs,
				"page":       c.GetInt("page"),
				"pageSize":   c.GetInt("pageSize"),
				"totalPages": c.GetInt("totalPages"),
			},
		)
	}

}

func BlogsShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	blog := models.BlogsFind(id)
	switch getFormat(c) {
	case formatHTML:
		c.HTML(
			http.StatusOK,
			"blogs/show.tpl",
			gin.H{
				"blog": blog,
			},
		)
	case formatJSON:
		c.JSON(
			http.StatusOK,
			gin.H{
				"blog": blog,
			},
		)
	}
}
