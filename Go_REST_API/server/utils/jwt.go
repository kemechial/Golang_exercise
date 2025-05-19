package utils
import (
	 "github.com/golang-jwt/jwt/v5"
	 "time"
	 "errors"
)


const secretKey = "some secret value"
	


func GenerateToken(email string, userID int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}


func ValidateToken(tokenString string) error {	
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return errors.New("unexpected signing method"), nil
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("invalid token")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid claims")
	}

	//email, ok := claims["email"].(string)
	//userID, ok := claims["userId"].(float64)

	return nil

}