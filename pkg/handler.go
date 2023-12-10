package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Dao *Dao
}

func NewHandler(db *sqlx.DB) *Handler {

	return &Handler{
		Dao: NewDao(db),
	}
}

type AuthBeforeWSRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *Handler) AuthBeforeWSHandler(c *gin.Context) {
	var req AuthBeforeWSRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authToken := c.GetHeader("Authorization")

	u := NewJwty().DecodeJwt(authToken)
	fmt.Println(u)

	user, err := h.Dao.GetUserByID(u.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !(u.Id == user.ID && u.Email == user.Email && u.Name == user.Name) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("validate failed")})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth": "ok"})
}

// restful handler
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignupHandler(c *gin.Context) {
	var req SignupRequest
	if err := c.BindJSON(&req); err != nil {
		// Handle error, maybe return a 400 Bad Request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExists, err := h.Dao.userExists(req.Email)
	if err != nil {
		return
	}
	if userExists {
		err := fmt.Errorf("user exists")
		c.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	err = h.Dao.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "signup success"})

}

type SigninRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (h *Handler) SigninHandler(c *gin.Context) {
	var req SigninRequest
	if err := c.BindJSON(&req); err != nil {
		// Handle error, maybe return a 400 Bad Request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Dao.GetUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	token, err := NewJwty().FastJwt(user.ID, user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, SigninResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  GetNameFromEmail(user.Email),
		Token: token,
	})

}

// other handler
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "world"})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
