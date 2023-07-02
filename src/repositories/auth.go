package repositories

import (
	"errors"
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

	var id uint64
	sql := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	row := repository.db.QueryRowx(sql, user.Username, hash)
	if row.Err() != nil {
		return "", row.Err()
	}
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return generateToken(id, 7*24*time.Hour)
}

func (repository AuthRepositoryImpl) Login(user models.User) (string, error) {
	userFromDB := models.User{}
	sql := "SELECT * FROM users WHERE username = $1"
	if err := repository.db.Get(&userFromDB, sql, user.Username); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		return "", errors.New("wrong password")
	}

	return generateToken(userFromDB.ID, 7*24*time.Hour)
}
