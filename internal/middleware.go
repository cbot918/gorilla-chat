package internal

import (
	"fmt"
	"net/http"

	"gorilla-chat/internal/pkg/jwty"
	"gorilla-chat/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type RequireLoginRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequireLogin(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		var req RequireLoginRequest
		if err := c.BindJSON(&req); err != nil {
			// Handle error, maybe return a 400 Bad Request response
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authToken := c.GetHeader("Authorization")

		// Check if the Authorization header is empty
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		user := jwty.NewJwty().DecodeJwt(authToken)
		var u types.User
		err := db.Get(&u, "SELECT id,email,name FROM users WHERE id=?", user.Id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "bad token or user data"})
			return
		}
		fmt.Println("user: ")
		fmt.Println(user)
		fmt.Println("u: ")
		fmt.Println(u)
		// Process request
		c.Next()

	}
}
