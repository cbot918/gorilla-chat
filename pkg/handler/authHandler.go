package handler

import (
	"fmt"
	"gorilla-chat/pkg/jwty"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	u := jwty.NewJwty().DecodeJwt(authToken)
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
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(req)

	emailExists, err := h.Dao.EmailExists(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if emailExists {
		err := fmt.Errorf("email exists")
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	nameExists, err := h.Dao.NameExists(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if nameExists {
		err := fmt.Errorf("name exists")
		c.JSON(http.StatusBadRequest, errorResponse(err))
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
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := h.Dao.GetUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	token, err := jwty.NewJwty().FastJwt(user.ID, user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	c.JSON(200, SigninResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	})

}
