package authtoken

import (
	"net/http"
	"testing"
)

func TestBearerFromRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Set("Authorization", "Bearer TOKEN")
	result, err := FromRequest(req)

	if err != nil {
		t.Errorf("err = %v want nil", err)
	}

	if result != "TOKEN" {
		t.Errorf("FromRequest() = %v want %v", result, "TOKEN")
	}
}

func TestBasicFromRequestNotAllowed(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Set("Authorization", "Basic VE9LRU4=")
	result, err := FromRequest(req)
	expected := "Authorization requires Bearer scheme"
	if err.Error() != expected {
		t.Errorf("err = %v want nil", err)
	}

	if result != "" {
		t.Errorf("FromRequest() = %v want %v", result, "")
	}
}

func TestEmptyFromRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	expected := "Authorization header required"
	_, err := FromRequest(req)

	if err.Error() != expected {
		t.Errorf("err = %v want %v", err, expected)
	}
}
