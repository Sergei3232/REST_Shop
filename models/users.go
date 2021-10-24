package models

import (
	u "github.com/Sergei3232/REST_Shop/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Users struct {
	gorm.Model
	Account  string `json:"account"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Token    string `json:"token";sql:"-"`
}

func (users *Users) Validate() (map[string]interface{}, bool) {
	if len(users.Password) < 6 {
		return u.Message(400, "Password is required"), false
	}

	//Account должен быть уникальным
	temp := &Users{}
	//проверка на наличие ошибок и дубликатов аккаунтов в таблице юзеров
	err := GetDB().Table("users").Where("account = ?", users.Account).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(500, "Connection error. Please retry"), false
	}
	if temp.Account != "" {
		return u.Message(400, "Account already in use by another user."), false
	}

	return u.Message(200, "Requirement passed"), true
}

func (account *Users) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	if account.Type == "" {
		account.Type = User
	}
	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(500, "Failed to create account, connection error.")
	}

	//Создать новый токен JWT для новой зарегистрированной учётной записи
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //удалить пароль

	response := u.Message(201, "Account has been created")
	response["account"] = account
	return response
}

func Login(account, password string) map[string]interface{} {
	users := &Users{}
	err := GetDB().Table("users").Where("account = ?", account).First(users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(404, "Account not found")
		}
		return u.Message(500, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
		return u.Message(400, "Invalid login credentials. Please try again")
	}
	//Работает! Войти в систему
	users.Password = ""

	//Создать токен JWT
	tk := &Token{UserId: users.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	users.Token = tokenString // Сохраните токен в ответе

	resp := u.Message(200, "Logged In")
	resp["account"] = users
	return resp
}

func GetUser(account string) *Users {
	acc := &Users{}
	GetDB().Table("users").Where("account = ?", account).First(acc)
	if acc.Account == "" { //Пользователь не найден!
		return nil
	}

	acc.Password = ""
	return acc
}
