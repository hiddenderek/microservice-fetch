package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"fetch-service/internal/models"
	"fetch-service/internal/service"
)

// @BasePath /api/

// ProcessReceipts godoc
// @Summary      Process a receipt
// @Tags         receipts
// @Accept       json
// @Produce      json
// @Param        receipt   body   models.Receipt  true "Receipt Data"
// @Failure      400
// @Failure      404
// @Failure      500
// @Router /receipts/process [post]
func processReceipts(c *gin.Context) {
	var req models.Receipt
	// validates incoming json body
	if err := c.ShouldBind(&req); err != nil {
		// TODO: enhance validation errors to be more readable
		c.JSON(http.StatusBadRequest, gin.H{"Validation Errors": fmt.Sprintf("%v", err)})
		return
	}
	defer func() {
		// handle any unforseen panics
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Unknown Error": r})
		}
	}()
	res := service.PostReceipt(req)
	c.JSON(http.StatusCreated, res)
}

// getPoints godoc
// @Summary      Calculate points for a receipt
// @Tags         receipts
// @Produce      json
// @Param        id   path      string  true  "Receipt Id"
// @Failure      400
// @Failure      404
// @Failure      500
// @Router 		 /receipts/{id}/points [get]
func getPoints(c *gin.Context) {
	defer func() {
		// handle any unforseen panics
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Unknown Error": r})
		}
	}()
	name := c.Param("id")
	var item, err, ok = service.GetPoints(name)
	if err != nil {
		// handle any returned errors
		c.JSON(http.StatusBadRequest, gin.H{"Malformed Data Error": err.Error()})
	} else if !ok {
		// handle 404 errors
		c.JSON(http.StatusNotFound, gin.H{"Not Found Error": "Could Not Find Points"})
	} else if ok {
		// return score if nothing is wrong
		c.JSON(http.StatusOK, item)
	}

}

// group up routes
func AddRoutes(route *gin.RouterGroup) {
	user := route.Group("/receipts")
	{
		user.POST("/process", processReceipts)
		user.GET("/:id/points", getPoints)
	}
}
