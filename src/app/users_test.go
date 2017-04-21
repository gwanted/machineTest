package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"conf"
	"mydb"
)

func InitTestDB() {
	mydb.InitDB(conf.App.DBAddress)
}

func TestLogin(t *testing.T) {
	InitTestDB()
	v := url.Values{}
	v.Add("account", "1")
	v.Add("pwd", "123456")

	req, err := http.NewRequest("POST", "/login", strings.NewReader(v.Encode()))
	if err != nil {
		t.Fatal(err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"code":0,"msg":"SUCCESS","data":{"id":"589d70e215cb6a1bb700e535","account":"1","pwd":"123456"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestRegister(t *testing.T) {
	InitTestDB()
	v := url.Values{}
	v.Add("account", "12")
	v.Add("pwd", "123456")

	req, err := http.NewRequest("POST", "/register", strings.NewReader(v.Encode()))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"code":0,"msg":"SUCCESS","data":null}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
