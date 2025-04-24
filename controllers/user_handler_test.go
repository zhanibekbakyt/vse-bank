package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"vse-bank/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var router *gin.Engine
var db *gorm.DB

func setupRouter() {
	router = gin.Default()
	SetDB(db)

	router.GET("/users", GetUsers)
	router.POST("/users", AddUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)
	router.GET("/banks", GetBanks)
	router.POST("/banks", AddBank)
	router.PUT("/banks/:id", UpdateBank)
	router.DELETE("/banks/:id", DeleteBank)
}

func setupDatabase() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Bank{})
	return db
}

func TestGetUsers(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	user := models.User{Name: "testuser", Email: "testuser@example.com"}
	db.Create(&user)

	req, _ := http.NewRequest("GET", "/users", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var users []models.User
	json.NewDecoder(recorder.Body).Decode(&users)
	assert.Len(t, users, 1)
	assert.Equal(t, "testuser", users[0].Name)
}

func TestAddUser(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	user := models.User{Name: "newuser", Email: "newuser@example.com"}
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	var addedUser models.User
	json.NewDecoder(recorder.Body).Decode(&addedUser)
	assert.Equal(t, "newuser", addedUser.Name)
}

func TestUpdateUser(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	user := models.User{Name: "testuser", Email: "testuser@example.com"}
	db.Create(&user)

	updateData := models.User{Name: "updateduser", Email: "updateduser@example.com"}
	body, _ := json.Marshal(updateData)
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewReader(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var updatedUser models.User
	json.NewDecoder(recorder.Body).Decode(&updatedUser)
	assert.Equal(t, "updateduser", updatedUser.Name)
}

func TestDeleteUser(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	user := models.User{Name: "testuser", Email: "testuser@example.com"}
	db.Create(&user)

	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response map[string]string
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, "User deleted successfully", response["message"])
}

func TestGetBanks(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	bank := models.Bank{Name: "Bank A"}
	db.Create(&bank)

	req, _ := http.NewRequest("GET", "/banks", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var banks []models.Bank
	json.NewDecoder(recorder.Body).Decode(&banks)
	assert.Len(t, banks, 1)
	assert.Equal(t, "Bank A", banks[0].Name)
}

func TestAddBank(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	bank := models.Bank{Name: "Bank B"}
	body, _ := json.Marshal(bank)
	req, _ := http.NewRequest("POST", "/banks", bytes.NewReader(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var addedBank models.Bank
	json.NewDecoder(recorder.Body).Decode(&addedBank)
	assert.Equal(t, "Bank B", addedBank.Name)
}

func TestUpdateBank(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	bank := models.Bank{Name: "Bank C"}
	db.Create(&bank)

	updatedData := models.Bank{Name: "Updated Bank C"}
	body, _ := json.Marshal(updatedData)
	req, _ := http.NewRequest("PUT", "/banks/1", bytes.NewReader(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var updatedBank models.Bank
	json.NewDecoder(recorder.Body).Decode(&updatedBank)
	assert.Equal(t, "Updated Bank C", updatedBank.Name)
}

func TestDeleteBank(t *testing.T) {
	db = setupDatabase()
	setupRouter()

	bank := models.Bank{Name: "Bank D"}
	db.Create(&bank)

	req, _ := http.NewRequest("DELETE", "/banks/1", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]string
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, "Bank deleted successfully", response["message"])
}
