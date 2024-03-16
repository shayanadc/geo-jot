package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var Fields = graphql.Fields{
	"health": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "checked!", nil
		},
	},
}
var RootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: Fields}
var SchemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(RootQuery)}

func GetSchama() *graphql.Schema {
	schema, err := graphql.NewSchema(SchemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return &schema
}

func main() {
	fmt.Println(GenerateToken(1))
	http.Handle("/graphql", AuthMiddleware(handler.New(&handler.Config{
		Schema: GetSchama(),
	})))

	log.Println("geo-jot *** gql server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var secretKey = []byte("secretpassword")

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
