package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func TestCheckLoginRedirectUsesPanelBasePath(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name             string
		requestPath      string
		expectedLocation string
	}{
		{name: "custom base path", requestPath: "/app/api/load", expectedLocation: "/app/login"},
		{name: "root base path", requestPath: "/api/load", expectedLocation: "/login"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			engine := gin.New()
			engine.Use(sessions.Sessions("test", cookie.NewStore([]byte("test-secret"))))
			engine.GET(test.requestPath, checkLogin, func(c *gin.Context) {
				c.Status(http.StatusNoContent)
			})

			request := httptest.NewRequest(http.MethodGet, test.requestPath, nil)
			engine.ServeHTTP(recorder, request)

			if recorder.Code != http.StatusTemporaryRedirect {
				t.Fatalf("status = %d, want %d", recorder.Code, http.StatusTemporaryRedirect)
			}
			if location := recorder.Header().Get("Location"); location != test.expectedLocation {
				t.Fatalf("location = %q, want %q", location, test.expectedLocation)
			}
		})
	}
}

func TestCheckLoginXHRReturnsJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	engine := gin.New()
	engine.Use(sessions.Sessions("test", cookie.NewStore([]byte("test-secret"))))
	engine.GET("/app/api/load", checkLogin, func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	request := httptest.NewRequest(http.MethodGet, "/app/api/load", nil)
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	engine.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	if location := recorder.Header().Get("Location"); location != "" {
		t.Fatalf("unexpected redirect to %q", location)
	}
	if body := recorder.Body.String(); body != `{"success":false,"msg":"Invalid login","obj":null}` {
		t.Fatalf("body = %q", body)
	}
}
