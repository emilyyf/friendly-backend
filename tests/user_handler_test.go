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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	router.GET("/profile", handlers.AuthMiddleware(), handlers.ProfileHandler)

	return router, db, mock
}

func TestCreateUserHandler(t *testing.T) {
	router, db, mock := setupRouter(t)
	mockID := uuid.New()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"users\"").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), true, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "verified"}).
			AddRow(mockID.String(), "test@example.com", true))

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
	assert.True(t, response.Verified)

	var user *entities.User
	err = db.Where("email = ?", userInput.Email).First(&user).Error
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.True(t, user.Verified)
}

func TestLoginHandler(t *testing.T) {
	router, _, mock := setupRouter(t)

	userInput := entities.SignInInput{Email: "test@example.com", Password: "password123"}
	hashedPassword, _ := handlers.HashPassword(userInput.Password)
	mockID := uuid.New()
	mockUser := entities.User{ID: mockID, Email: userInput.Email, Password: hashedPassword, Verified: true}

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

// func TestProfileHandler(t *testing.T) {
//
// 	router, db := setupRouter(t)
// 	defer db.Exec("DROP TABLE users")
//
// 	userInput := entities.SignInInput{Email: "test@example.com", Password: "password123"}
// 	hashedPassword, _ := handlers.HashPassword(userInput.Password)
// 	db.Create(&entities.User{Email: userInput.Email, Password: hashedPassword, Verified: true})
//
// 	loginBody, _ := json.Marshal(userInput)
// 	reqLogin, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
// 	wLogin := httptest.NewRecorder()
// 	router.ServeHTTP(wLogin, reqLogin)
//
// 	var loginResponse map[string]string
// 	json.Unmarshal(wLogin.Body.Bytes(), &loginResponse)
// 	token := loginResponse["token"]
//
// 	reqProfile, _ := http.NewRequest("GET", "/profile", nil)
// 	reqProfile.Header.Set("Authorization", token)
// 	wProfile := httptest.NewRecorder()
// 	router.ServeHTTP(wProfile, reqProfile)
//
// 	assert.Equal(t, http.StatusOK, wProfile.Code)
// }
//
// func TestAuthMiddleware(t *testing.T) {
// 	router, _ := setupRouter(t)
//
// 	req, _ := http.NewRequest("GET", "/profile", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	assert.Equal(t, http.StatusUnauthorized, w.Code)
//
// 	req.Header.Set("Authorization", "invalid-token")
// 	w = httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	assert.Equal(t, http.StatusUnauthorized, w.Code)
// }
//
// func TestCreateUserBadRequest(t *testing.T) {
// 	router, _ := setupRouter(t)
//
// 	body := []byte(`{"email": "test@example.com"}`)
// 	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(body))
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }
