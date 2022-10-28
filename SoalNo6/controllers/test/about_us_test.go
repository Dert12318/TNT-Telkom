package test

import (
	"SoalNo6/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAboutUs(t *testing.T) {
	engine := Engine()
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/api/v1/about-us", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	engine.ServeHTTP(recorder, request)
	if recorder.Code != 200 {
		t.Fatalf("bad status code: %d", recorder.Code)
	}

	var response models.ResponseSuccess
	body := recorder.Body.String()
	if err != nil {
		t.Fatalf("reading response body: %v", err)
	}
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("parsing json response: %v", err)
	}
	if response.ResponseMessage != "success" {
		t.Fatalf("bad response message: %s", response.ResponseMessage)
	}

	marshal, err := json.Marshal(response.Data)
	if err != nil {
		return
	}

	var result models.Cart
	err = json.Unmarshal(marshal, &result)
	if err != nil {
		return
	}

}

func TestGetAboutUsNotFound(t *testing.T) {
	engine := Engine()
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/hello-about-us", nil)
	if err != nil {
		t.Fatalf("building request: %v", err)
	}
	engine.ServeHTTP(recorder, request)
	if recorder.Code != 404 {
		t.Fatalf("bad status code: %d", recorder.Code)
	}
}
