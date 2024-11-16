package security

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Jefschlarski/golang-modular/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a JWT token with the given userID and returns the token as a string.
//
// Parameters:
// - userID: uint64 - the ID of the user to be included in the token.
//
// Returns:
// - string: the generated JWT token as a string.
// - error: an error if the token generation fails.
func GenerateToken(userID uint64) (string, error) {

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	AUTHConfig := config.GetAuthConfig()

	tokenString, err := token.SignedString([]byte(AUTHConfig.Key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the given JWT token and returns the user ID if the token is valid.
func ValidateToken(ctx *gin.Context) error {
	tokenString := extractToken(ctx)
	if tokenString == "" {
		return fmt.Errorf("token is required to access this endpoint")
	}
	token, err := jwt.Parse(tokenString, verifyTokenSignKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// ExtractUserID extracts the user ID from the given JWT token.
func ExtractUserID(ctx *gin.Context) (uint64, error) {

	tokenString := extractToken(ctx)
	token, err := jwt.Parse(tokenString, verifyTokenSignKey)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%v", claims["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}
	return 0, fmt.Errorf("invalid token")
}

// extractToken extracts the token from the given request.
func extractToken(ctx *gin.Context) string {

	token := ctx.GetHeader("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// verifyTokenSignKey verifies the sign key from the token is the same as the expected sign key and returns it.
func verifyTokenSignKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.GetAuthConfig().Key), nil
}

// VerifyId verifies that the user ID in the request matches with the token user ID.
func VerifyId(id uint64, ctx *gin.Context) error {

	tokenId, err := ExtractUserID(ctx)
	if err != nil {
		return err
	}

	if tokenId != id {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
