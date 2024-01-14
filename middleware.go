package main

import (
	"fmt"
	"os"

	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
)
func denyAccess(w http.ResponseWriter){
	 writeJSON(w, http.StatusForbidden, APIError{Error: "Access has been denied"})
	
}

//id 11 
func jwtAuthFunc(handlerFunc http.HandlerFunc , store storage) http.HandlerFunc {
	//fmt.Println("Using JWT Middleware")

	return func(w http.ResponseWriter, r *http.Request) {
		// get token from http req header
		tokenString := r.Header.Get("x-jwt-token")
		
	
		token, err := validateJWT(tokenString)
		if err != nil {
			denyAccess(w)
			return
		}
		if(!token.Valid){
			
			denyAccess(w)
			return
		}
		claims:=token.Claims.(jwt.MapClaims)
		userId,err:=getIDFromRequest(r)
		if(err!=nil){
			
			return
		}
		account, err :=store.GetAccountByID(userId)
		if err != nil {
			denyAccess(w)
			return
		}
		//Yeah this sucks.
		if(account.Number !=int64(claims["AccountNumber"].(float64))){
			
			denyAccess(w)
			return
		}
		fmt.Println(claims)
		handlerFunc(w, r)
	}

}

// Reference for this validation func -https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac
func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("jwt_secret")

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		hmacSampleSecret := []byte(secret)
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

}
