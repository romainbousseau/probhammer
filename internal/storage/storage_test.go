package storage

import (
	"net/http/httptest"
	"reflect"
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

func TestStorage_FindDatasheetByID(t *testing.T) {
	type args struct {
		ctx *gin.Context
		id  uint
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{
		{
			name: "Returns a datasheet",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: "Boyz",
		},
		{
			name: "Returns an error if datasheet does not exist",
			args: args{
				ctx: ctx,
				id:  4,
			},
			wantErr: "record not found",
		},
	}
	for _, tt := range tests {
		db, cleanup := utils.SetTestDB(t)
		s := NewStorage(db)

		t.Run(tt.name, func(t *testing.T) {
			createDatasheets(t, db)
			t.Cleanup(cleanup)

			got, err := s.FindDatasheetByID(tt.args.ctx, tt.args.id)
			if (err != nil) && (err.Error() != tt.wantErr) {
				t.Errorf("Storage.FindDatasheetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) && (!reflect.DeepEqual(got.Name, tt.want)) {
				t.Errorf("Storage.FindDatasheetByID() = %v, want %v", got.Name, tt.want)
			}
		})
	}
}

func TestStorage_CreateDatasheet(t *testing.T) {
	type args struct {
		ctx       *gin.Context
		datasheet *models.Datasheet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Create a new datasheet",
			args: args{
				ctx:       ctx,
				datasheet: &models.Datasheet{Name: "Orks"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		db, cleanup := utils.SetTestDB(t)
		s := NewStorage(db)

		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(cleanup)

			if err := s.CreateDatasheet(tt.args.ctx, tt.args.datasheet); (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateDatasheet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createDatasheets(t *testing.T, db *gorm.DB) {
	datasheets := []models.Datasheet{
		{Name: "Boyz"},
		{Name: "Nobz"},
		{Name: "Weirdboy"},
	}

	db.Create(&datasheets)
}
