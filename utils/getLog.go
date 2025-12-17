package utils

import "proyecto_transacciones/models"

func GetLogout(user *models.App) *models.Sublog {
	sub := models.Sublog{Srouter: user.Router.Group("/logout"), Depends: &user.Depend}
	return &sub
}

func GetLogin(user *models.App) *models.Sublog{
	sub := models.Sublog{Srouter: user.Router.Group("/login"),Depends: &user.Depend}
	return &sub
}

func GetTransfer(user *models.App) *models.Sublog{
	sub := models.Sublog{Srouter: user.Router.Group("/transfer"),Depends: &user.Depend}
	return &sub
}

