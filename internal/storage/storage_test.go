package storage

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/romainbousseau/probhammer/internal/models"
	"github.com/romainbousseau/probhammer/internal/utils"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/gorm"
)

var ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

// TODO: remove, this is a test function 
func TestStorage_FindDatasheets(t *testing.T) {
	t.Run("Returns all datasheets", func(t *testing.T) {
		db, cleanup := utils.SetTestDB(t)
		s := NewStorage(db)

		createDatasheets(t, db)
		t.Cleanup(cleanup)

		res, err := s.FindDatasheets(ctx)
		require.Nil(t, err)
		assert.Equal(t, len(res), 3)
	})
}

func createDatasheets(t *testing.T, db *gorm.DB) {
	datasheets := []models.Datasheet{
		{Name: "Boyz"},
		{Name: "Nobz"},
		{Name: "Weirdboy"},
	}

	db.Create(&datasheets)
}
