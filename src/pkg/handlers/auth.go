package handlers

import (
	"fmt"
	"mainapp/pkg/auth"
	"mainapp/pkg/models"
	auth_models "mainapp/pkg/models/auth"
	"mainapp/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func AuthMiddleware(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid cookie"})
			return
		}
	
		// 2) Parse & validate the token
		token, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
			return
		}
	
		// 3) Inspect claims, e.g. username
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
	
		// 4) Store username in context and proceed
		c.Set("username", claims["username"])
		c.Next()
}

// func AuthMiddleware(c *gin.Context) {
// 	// 1) Grap the authorisation header
// 	authHeader := c.GetHeader("Authorisation")

// 	if authHeader == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
// 		return
// 	}

// 	// 2) Expect "Bearer <token>"
// 	parts := strings.SplitN(authHeader, " ", 2)
// 	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
// 		return 
// 	}
// 	tokenString := parts[1]

// 	// 3) Parse & validate the token
// 	token, err := auth.ValidateToken(tokenString)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
// 		return
// 	}

// 	// 4) Inspect claims, e.g. username
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
// 		return
// 	}

// 	c.Set("username", claims["username"])
	
// 	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area"})
// }

func RegisterHandler(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := services.SaveUser(input.Username, input.Password)
	
	if err != nil {
		// user already exists
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) 
		return
	}

	tokenString, err := auth.CreateToken(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.SetCookie(
		"jwt",
		tokenString,
		3600*24,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusCreated, gin.H{"token": tokenString})
}

func LoginHandler(c *gin.Context) {
	var userInput auth_models.Login
	
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.FetchUser(userInput.Username)
	
	if err != nil {
        // 2a) user not in DB
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        } else {
            // 2b) some other DB error
            c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
        }
        return
    }

	if bcrypt.CompareHashAndPassword(
        []byte(user.Password),
        []byte(userInput.Password),
    ) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
	
	tokenString, err := auth.CreateToken(userInput.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erorr": "could not generate token"})
		return
	}

	c.SetCookie(
		"jwt",
		tokenString,
		3600*24,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func LogoutHandler(c *gin.Context) {
    c.SetCookie(
        "jwt",   // name
        "",      // value
        -1,      // maxAge<0 means delete now
        "/",     // path — sent to all URLs on this host
        "",      // domain — empty means “current host” (e.g. localhost)
        false,   // secure — allow over HTTP in dev
        true,    // httpOnly — not accessible via JavaScript
    )
	
	if err := services.LogoutUser(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "logout failed"})
        return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

