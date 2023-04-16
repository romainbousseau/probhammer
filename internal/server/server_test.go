// server is dedicated to build and run a server
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestServer_Calculate(t *testing.T) {
	type fields struct {
		storage Storage
		router  *gin.Engine
	}

	tests := []struct {
		name   string
		fields fields
		params []gin.Param
		code   int32
		body   string
	}{
		{
			name: "With correct parameters",
			fields: fields{
				storage: nil,
				router:  gin.Default(),
			},
			params: []gin.Param{
				{Key: "attacks", Value: "1"},
				{Key: "damage", Value: "1"},
				{Key: "hitSkill", Value: "3"},
				{Key: "strength", Value: "3"},
				{Key: "toughness", Value: "3"},
				{Key: "save", Value: "3"},
			},
			code: 200,
			body: `{"Hit":{"Success":1,"Failure":0,"Prob":0.6666666666666666},"Wound":{"Success":0,"Failure":1,"Prob":0.5},"Save":{"Success":0,"Failure":0,"Prob":0.6666666666666666},"ExpectedDamage":0}`,
		},
		{
			name: "With missing parameters",
			fields: fields{
				storage: nil,
				router:  gin.Default(),
			},
			params: []gin.Param{},
			code: 400,
			body: `{"error":"Key: 'Calculator.Attacks' Error:Field validation for 'Attacks' failed on the 'required' tag\nKey: 'Calculator.Damage' Error:Field validation for 'Damage' failed on the 'required' tag\nKey: 'Calculator.HitSkill' Error:Field validation for 'HitSkill' failed on the 'required' tag\nKey: 'Calculator.Strength' Error:Field validation for 'Strength' failed on the 'required' tag\nKey: 'Calculator.Toughness' Error:Field validation for 'Toughness' failed on the 'required' tag\nKey: 'Calculator.Save' Error:Field validation for 'Save' failed on the 'required' tag"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			buildTestRequest(ctx, tt.params, "/calculate")

			s := &Server{
				storage: tt.fields.storage,
				router:  tt.fields.router,
			}
			s.Calculate(ctx)

			assert.Equal(t, tt.code, int32(w.Code))
			assert.Equal(t, tt.body, w.Body.String())
		})
	}
}

// buildTestRequest builds a request needed for testing BindQuery
func buildTestRequest(ctx *gin.Context, params []gin.Param, path string) {
	ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)

	q := ctx.Request.URL.Query()
	for _, p := range params {
		q.Add(p.Key, p.Value)
	}

	ctx.Request.URL.RawQuery = q.Encode()
}
