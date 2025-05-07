package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romainbousseau/probhammer/internal/models"
)

// Find Datasheets returns all datasheets from DB
func (s *Server) FindDatasheets(ctx *gin.Context) {
	datasheets, err := s.storage.FindDatasheets(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, datasheets)
	}
}

// FindDatasheetByID returns a datasheet by ID
func (s *Server) FindDatasheetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	datasheet, err := s.storage.FindDatasheetByID(ctx, uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, datasheet)
	}
}

// CreateDatasheet creates a new datasheet
func (s *Server) CreateDatasheet(ctx *gin.Context) {
	var datasheet models.Datasheet

	if err := ctx.BindJSON(&datasheet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.storage.CreateDatasheet(ctx, &datasheet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

// DeleteDatasheet deletes a datasheet
func (s *Server) DeleteDatasheet(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.storage.DeleteDatasheet(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
