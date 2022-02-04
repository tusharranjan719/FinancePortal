package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var randomKey = "dwqerthjui"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		woutBearer := r.Header.Get("Authorization")
		if !strings.Contains(woutBearer, "Bearer") {
			ctx := context.WithValue(r.Context(), "props", jwt.MapClaims{"user_name": ""})
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			fmt.Printf("authheader -> %s and len -> %d\n", authHeader, len(authHeader))
			if len(authHeader) != 2 || authHeader[0] == "null" {
				//fmt.Println("Malformed token")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Malformed Token"))
				log.Fatal("Malformed token")
			} else {
				jwtToken := authHeader[1]
				token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Signing process unexpected: %v", token.Header["alg"])
					}
					return []byte(randomKey), nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.WithValue(r.Context(), "props", claims)

					next.ServeHTTP(w, r.WithContext(ctx))

				} else {
					fmt.Println("token err -> ", err)

					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Token expired or you are unauthorized"))
				}
			}
		}

	})
}

func MakeToken(userId uint64, name string) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", randomKey)
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["user_name"] = name
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", errors.New("Create token is causing error")
	}
	fmt.Println("JWT Map is ---> ", atClaims)
	fmt.Println("Token is ---> ", token)
	return token, nil
}
