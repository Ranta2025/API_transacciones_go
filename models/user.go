package models


type User struct{
	Username string
	Gmail string
	Saldo float64
}

type UserDB struct{
	User *User
	Id string
	Password string
}

type Login struct{
	Username string
	Password string
	Bloqueado string
}

type UserLogout struct{
	Username string
	Gmail string
	Password string
}