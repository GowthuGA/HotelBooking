package controllers

import (
    "github.com/gin-gonic/gin"
    "hotel-booking/models"
    "hotel-booking/services"
    "net/http"
)

type AdminController struct {
    Service *services.AdminService
}

func (ac *AdminController) LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Set content type to application/json
    // w.Header().Set("Content-Type", "application/json")

    // Parse the request body
    var loginRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
        return
    }

    // Call the repository's Authenticate method
    admin, err := ac.Service.Authenticate(loginRequest.Username, loginRequest.Password)
    if err != nil {
        http.Error(w, err.Message, err.Status)
        return
    }

    // Respond with the admin details or a success message
    response := struct {
        Message string         `json:"message"`
        Admin   *models.Admin  `json:"admin"`
    }{
        Message: "Login successful",
        Admin:   admin,
    }

    c.JSON(http.StatusOK, response)
}
