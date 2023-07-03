package repositories

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/SergeyCherepiuk/todo-app/src/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	SignUp(models.User) (string, error)
	Login(models.User) (string, error)
}

type AuthRepositoryImpl struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

type JwtCustomClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func generateToken(userId uint64, expiresIn time.Duration) (string, error) {
	claims := JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func (repository AuthRepositoryImpl) SignUp(user models.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", err
	}

	query := `INSERT INTO users (username, password) VALUES (:username, :password) RETURNING id`
	namedParams := map[string]any{
		"username": user.Username,
		"password": hash,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return "", fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var id uint64
	if err := stmt.Get(&id, namedParams); err != nil {
		return "", fmt.Errorf("failed to execute query: %w", err)
	}

	return generateToken(id, 7*24*time.Hour)
}

func (repository AuthRepositoryImpl) Login(user models.User) (string, error) {
	query := `SELECT * FROM users WHERE username = :username`
	namedParams := map[string]any{
		"username": user.Username,
	}

	stmt, err := repository.db.PrepareNamed(query)
	if err != nil {
		return "", fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	userFromDB := models.User{}
	if err := stmt.Get(&userFromDB, namedParams); err != nil {
		return "", fmt.Errorf("failed to execute query: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		return "", errors.New("wrong password")
	}

	return generateToken(userFromDB.ID, 7*24*time.Hour)
}
