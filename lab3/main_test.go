package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler_ValidCredentials(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader("username=testuser&password=testpass"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	loginHandler(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", status)
	}
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader("username=wronguser&password=wrongpass"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	loginHandler(w, req)

	if !strings.Contains(w.Body.String(), "Invalid credentials") {
		t.Errorf("Expected 'Invalid credentials', got %v", w.Body.String())
	}
}
func TestRegisterHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/register", strings.NewReader("username=newuser&password=newpass"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	registerHandler(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", status)
	}
}
func TestNotesHandler_Authenticated(t *testing.T) {
	req := httptest.NewRequest("GET", "/notes", nil)
	w := httptest.NewRecorder()

	session, _ := store.Get(req, "session-name")
	session.Values["authenticated"] = true
	session.Values["username"] = "testuser"
	session.Save(req, w)

	notesHandler(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", status)
	}
}
func TestLogoutHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/logout", nil)
	w := httptest.NewRecorder()

	session, _ := store.Get(req, "session-name")
	session.Values["authenticated"] = true
	session.Save(req, w)

	logoutHandler(w, req)

	if status := w.Code; status != http.StatusSeeOther {
		t.Errorf("Expected status code 303, got %v", status)
	}
}
