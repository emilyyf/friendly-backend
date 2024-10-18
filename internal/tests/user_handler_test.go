package tests

import (
	"bytes"
	"encoding/json"
	"friendly-backend/internal/db/entities"
	"friendly-backend/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	router.POST("/create", handlers.CreateUserHandler)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/profile", handlers.AuthMiddleware(), handlers.ProfileHandler)
	return router, db
}

func TestCreateUserHandler(t *testing.T) {
	router, db := setupRouter(t)

	userInput := entities.SignInInput{Email: "test@example.com", Password: "password123"}
	body, _ := json.Marshal(userInput)
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var user entities.User
	db.Where("email = ?", userInput.Email).First(&user)
	assert.NotEmpty(t, user.ID)
	assert.True(t, user.Verified)
}

func TestLoginHandler(t *testing.T) {
	router, db := setupRouter(t)
	defer db.Exec("DROP TABLE users")

	userInput := entities.SignInInput{Email: "test@example.com", Password: "password123"}
	hashedPassword, _ := handlers.HashPassword(userInput.Password)
	db.Create(&entities.User{Email: userInput.Email, Password: hashedPassword, Verified: true})

	body, _ := json.Marshal(userInput)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"])
}

func TestProfileHandler(t *testing.T) {
	router, db := setupRouter(t)
	defer db.Exec("DROP TABLE users")

	userInput := entities.SignInInput{Email: "test@example.com", Password: "password123"}
	hashedPassword, _ := handlers.HashPassword(userInput.Password)
	db.Create(&entities.User{Email: userInput.Email, Password: hashedPassword, Verified: true})

	loginBody, _ := json.Marshal(userInput)
	reqLogin, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
	wLogin := httptest.NewRecorder()
	router.ServeHTTP(wLogin, reqLogin)

	var loginResponse map[string]string
	json.Unmarshal(wLogin.Body.Bytes(), &loginResponse)
	token := loginResponse["token"]

	reqProfile, _ := http.NewRequest("GET", "/profile", nil)
	reqProfile.Header.Set("Authorization", token)
	wProfile := httptest.NewRecorder()
	router.ServeHTTP(wProfile, reqProfile)

	assert.Equal(t, http.StatusOK, wProfile.Code)
}

func TestAuthMiddleware(t *testing.T) {
	router, _ := setupRouter(t)

	req, _ := http.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	req.Header.Set("Authorization", "invalid-token")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateUserBadRequest(t *testing.T) {
	router, _ := setupRouter(t)

	body := []byte(`{"email": "test@example.com"}`)
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
