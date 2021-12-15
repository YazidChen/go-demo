package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userJSON = "{\"id\":1,\"name\":\"Yazid\"}\n"

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &userController{}

	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "1", rec.Body.String())
	}
}

func TestShow(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("id", "1")
	q.Set("name", "Yazid")
	req := httptest.NewRequest(http.MethodGet, "/user/show?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &userController{}

	if assert.NoError(t, h.show(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
