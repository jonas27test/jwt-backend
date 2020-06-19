package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jonas27test/jwt-backend/cmd/db"
)

func Test_Verify(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	headerToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvbmFzQHRlc3QudCIsImV4cCI6IjIwMjEtMDYtMTlUMTA6Mzg6MzQrMDI6MDAiLCJpc3MiOiJ0YXNrZXIifQ.3X2_tPBIZ6UEDmE68paK3LWV0WDC14HSYNSW29heaJw"
	req, err := http.NewRequest(http.MethodGet, "/verify", strings.NewReader(""))
	req.Header.Set("Authorization", headerToken)
	tFatal(t, err)
	c := Controller{DB: db.DB{DB: db.Connection(dbURL)}}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.Verify)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
