// server is dedicated to build and run a server
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romainbousseau/probhammer/internal/calculator"
	"github.com/romainbousseau/probhammer/internal/models"
)

// Server implements a server object
type Server struct {
	storage Storage
	router  *gin.Engine
}

type Storage interface {
	FindDatasheets(ctx *gin.Context) ([]*models.Datasheet, error)
	CreateDatasheet(ctx *gin.Context, datasheet *models.Datasheet) error
	FindDatasheetByID(ctx *gin.Context, id uint) (*models.Datasheet, error)
	DeleteDatasheet(ctx *gin.Context, id uint) error
}

// NewServer builds a new server
func NewServer(storage Storage, router *gin.Engine) Server {
	return Server{storage: storage, router: router}
}

// SetRoutesAndRun set the API routes and runs the router
func (s *Server) SetRoutesAndRun() error {

	s.router.GET("/", s.Ping)

	s.router.GET("/datasheets", s.FindDatasheets)
	s.router.GET("/datasheets/:id", s.FindDatasheetByID)
	s.router.POST("/datasheets", s.CreateDatasheet)
	s.router.DELETE("/datasheets/:id", s.DeleteDatasheet)

	s.router.GET("/calculate", s.Calculate)

	err := s.router.Run()
	if err != nil {
		return err
	}

	return nil
}

// Ping checks health of the app
func (s *Server) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "It is fine 🔥"})
}

// Calculate an attack based on the query parameters
func (s *Server) Calculate(ctx *gin.Context) {
	var c calculator.Calculator

	if err := ctx.BindQuery(&c); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Calculate()

	ctx.JSON(http.StatusOK, c.Results)
}
