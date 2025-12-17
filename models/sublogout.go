package models

import (

	"github.com/gin-gonic/gin"
)

type Sublog struct {
	Srouter *gin.RouterGroup
	Depends *Dependencias
}

