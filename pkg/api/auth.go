package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Password struct {
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Claims struct {
	Hash string `json:"pwd_hash"`
	jwt.RegisteredClaims
}

const TODO_PASSWORD = "TODO_PASSWORD"

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// смотрим наличие пароля
		pass := os.Getenv(TODO_PASSWORD)
		if len(pass) > 0 {
			var jwtCoockie string // JWT-токен из куки
			// получаем куку
			cookie, err := r.Cookie("token")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			jwtCoockie = cookie.Value
			// здесь код для валидации и проверки JWT-токена

			claims := Claims{}
			token, err := jwt.ParseWithClaims(jwtCoockie, &claims,
				func(token *jwt.Token) (interface{}, error) {
					// желательно проверять используемый метод
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(pass), nil
				})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !token.Valid {
				// возвращаем ошибку авторизации 401
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(*Claims); ok {
				currentHash := sha256.Sum256([]byte(pass))
				currentHashStr := hex.EncodeToString(currentHash[:])

				if claims.Hash != currentHashStr {
					http.Error(w, "Authentication required", http.StatusUnauthorized)
					return
				}
			} else {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var password Password
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonError(w, err)
		return
	}

	defer r.Body.Close()

	passwordEnv := os.Getenv(TODO_PASSWORD)
	if passwordEnv != "" {
		var token Token
		if passwordEnv != password.Password {
			w.WriteHeader(http.StatusUnauthorized)
			writeJsonError(w, errors.New("Неверный пароль"))
			return
		}

		hash := sha256.Sum256([]byte(passwordEnv))
		hashStr := hex.EncodeToString(hash[:])

		claims := &Claims{
			Hash: hashStr,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		token.Token, err = jwtToken.SignedString([]byte(passwordEnv))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJsonError(w, err)
			return
		}

		writeJson(w, token)
	}
}
