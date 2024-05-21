package tests

import (
    "go-api/controllers"
    "go-api/models"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/books", controllers.GetBooks)
    r.GET("/books/:id", controllers.GetBook)
    r.POST("/books", controllers.CreateBook)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)
    return r
}

func TestGetBooks(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/books", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBook(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/books/1", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBook(t *testing.T) {
    router := setupRouter()

    newBook := `{"title": "Test Driven Development", "author": "Kent Beck"}`
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/books", strings.NewReader(newBook))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateBook(t *testing.T) {
    router := setupRouter()

    updatedBook := `{"title": "The Go Programming Language", "author": "Alan Donovan & Brian Kernighan"}`
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/books/1", strings.NewReader(updatedBook))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteBook(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/books/1", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
