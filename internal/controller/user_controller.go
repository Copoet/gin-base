package controllers

import (
	"gin-base/internal/model"
	"github.com/gin-gonic/gin"
	"log"
)

type UserController struct{}

func (u *UserController) Create(c *gin.Context) {
	var user model.User
	log.Print(user)
}
