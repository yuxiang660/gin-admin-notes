package main

import (
	"log"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/welcome", Welcome)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func getToken(username string, expirationTime time.Time) (token string, err error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Please enter correct JSON format of credential!")
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Please enter correct credential!")
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	tokenString, err := getToken(creds.Username, expirationTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal error!")
		return
	}

	log.Println("Login and set cookie:")
	log.Println("Token:", tokenString)
	log.Println("Expires:", expirationTime)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Please login, and send with cookie!")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(c.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Please login with correct credential!")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Please login with correct credential!")
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
