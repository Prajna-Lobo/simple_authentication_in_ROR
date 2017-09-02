package handler

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"Demo-auth/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"
	"Demo-auth/factory"
	"golang.org/x/crypto/bcrypt"
	"log"
)
var MySigningKey = []byte("secret")


func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user model.UserCredentials
	var databaseUsername string
	var databasePassword string

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	dbConnection, err := factory.SqlConnector()
	if(dbConnection != nil) {

		dbConnection.QueryRow("SELECT username, password FROM user WHERE username=?",
			user.Username).Scan(&databaseUsername, &databasePassword)

		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(user.Password))
		if err != nil {
			http.Redirect(w, r, "/login", 301)
			return
		}

		w.Write([]byte("Hello " + databaseUsername))
		GetToken(w)
	}else{
		log.Println("database connection failed")
	}
}

func GetToken(w http.ResponseWriter) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "testing"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(MySigningKey)

	tokenObject := &model.AccessToken{
		Token: tokenString,
	}
	payload, _ := json.Marshal(tokenObject)


	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
var SimpleHandler =  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`success`))
})
var JwtValidator = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})