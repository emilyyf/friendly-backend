package tests

import (
	"bytes"
	"encoding/json"
	"friendly-backend/internal/db/entities"
	"friendly-backend/internal/handlers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func setupRouter(t *testing.T) (*gin.Engine, *gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm DB:  %v", err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	router.POST("/create", handlers.CreateUserHandler)
	router.POST("/login", handlers.LoginHandler)

	return router, db, mock
}

func TestCreateUserHandler(t *testing.T) {
	router, db, mock := setupRouter(t)
	mockID := uuid.New()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"users\"").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), false, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "verified"}).
			AddRow(mockID.String(), "test@example.com", false))

	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there are unfulfilled expectations: %s", err)
		}
	})

	userInput := entities.SignInInput{Email: "test@example.com", Password: "password123"}
	body, err := json.Marshal(userInput)
	if err != nil {
		t.Fatalf("failed to marshal user input: %v", err)
	}
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.User
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	assert.Equal(t, "test@example.com", response.Email)
	assert.False(t, response.Verified)

	var user *entities.User
	err = db.Where("email = ?", userInput.Email).First(&user).Error
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.False(t, user.Verified)
}

func TestLoginHandler(t *testing.T) {
	router, _, mock := setupRouter(t)

	userInput := entities.SignInInput{Email: "test@example.com", Password: "password123"}
	hashedPassword, _ := handlers.HashPassword(userInput.Password)
	mockID := uuid.New()
	mockUser := entities.User{ID: mockID, Email: userInput.Email, Password: hashedPassword, Verified: false}

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(userInput.Email, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "verified"}).
			AddRow(mockUser.ID.String(), mockUser.Email, mockUser.Password, mockUser.Verified))

	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there are unfulfilled expectations: %s", err)
		}
	})

	body, err := json.Marshal(userInput)
	if err != nil {
		t.Fatalf("failed to marshal user input: %v", err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	assert.NotEmpty(t, response["token"])
}

func TestCreateUserBadRequest(t *testing.T) {
	router, _, _ := setupRouter(t)

	body := []byte(`{"Email": "test@example.com"}`)
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthMiddleware_ValidToken(t *testing.T) {

	mockUserID := "12345"
	claims := jwt.StandardClaims{
		Subject: mockUserID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecret)

	router := gin.Default()
	router.Use(handlers.AuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"message": "Access granted", "userID": userID})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", tokenString)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, "Access granted", response["message"])
	assert.Equal(t, mockUserID, response["userID"])
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	router := gin.Default()
	router.Use(handlers.AuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You have acess"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalidToken")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, "Invalid token", response["error"])
}

func TestAuthMiddleware_NoTokenProvided(t *testing.T) {
	router := gin.Default()
	router.Use(handlers.AuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You have access"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, "No token provided", response["error"])
}

func TestProfileHandler_UserNotFound(t *testing.T) {
	router, _, mock := setupRouter(t)

	mockUserID := "nonexistent"
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1 LIMIT 1`).
		WithArgs(mockUserID).WillReturnError(gorm.ErrRecordNotFound)

	req, _ := http.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "Bearer validToken")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, "", response["error"])
}
