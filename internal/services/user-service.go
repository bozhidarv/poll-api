package services

import (
	"fmt"
	"time"

	"github.com/bozhidarv/poll-api/internal/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const JWT_SECRET = "asfbgakjl;gawobi;ioragjewgnVENBVRWOB"

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(user models.User) (string, error) {
	err, db := CheckDb()
	if err != nil {
		return "", err
	}

	hashedPass, err := hashPassword(*user.Password)
	if err != nil {
		return "", err
	}

	row := db.QueryRow(
		"INSERT INTO users (username, email, password, last_updated) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Username,
		user.Email,
		hashedPass,
		time.Now(),
	)

	var userId string
	err = row.Scan(&userId)
	if err != nil {
		apiError := &models.ApiError{
			Code:    500,
			Message: "Error while registering user.",
		}

		return "", apiError
	}

	return userId, nil
}

func CreateJwtToken(userId string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		Logger.Error().Msg(err.Error())
	}

	return tokenString
}

func checkTokenAlg(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return JWT_SECRET, nil
}

func ParseJwtToken(tokenString string) (string, float64) {
	token, err := jwt.Parse(tokenString, checkTokenAlg)
	if err != nil {
		Logger.Error().Msg(err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId := claims["userId"].(string)
		expire := claims["expire"].(float64)
		return userId, expire
	} else {
		Logger.Error().Msg(err.Error())
		return "", 0
	}
}
