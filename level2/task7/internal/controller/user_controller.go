package controller

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"task7/pkg/models"
	"time"
)

type UserController struct {
	controller *Controller
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (userController *UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, user)

	checkUser := new(models.User)
	userController.controller.DB.GormDB.Where("email = ?", user.Email).First(checkUser)

	if checkUser.ID != "" {
		return c.String(http.StatusOK, "this user already exist!")
	}

	user.ID, _ = userController.GetActualID()
	user.SecretKey = RandStringRunes(10)
	userController.controller.DB.GormDB.Create(user)

	data := user.ID + " " + user.SecretKey
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = data
	cookie.Expires = time.Now().Add(2 * time.Minute)
	c.SetCookie(cookie)

	return c.String(http.StatusCreated, "user has been created")
}

func (userController *UserController) LogIn(c echo.Context) error {
	user := new(models.User)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, user)

	checkUser := new(models.User)
	userController.controller.DB.GormDB.Where("email = ?", user.Email).First(checkUser)

	if checkUser.ID == "" {
		return c.String(http.StatusOK, "wrong email!")
	}

	if checkUser.Password != user.Password {
		return c.String(http.StatusOK, "wrong password!")
	}

	data := checkUser.ID + " " + checkUser.SecretKey
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = data
	cookie.Expires = time.Now().Add(2 * time.Minute)
	c.SetCookie(cookie)

	return c.String(http.StatusCreated, "success log in")
}

func (userController *UserController) CreateUserOAuth2(c echo.Context, user *models.User) error {
	userController.controller.DB.GormDB.Create(user)

	data := user.ID + " " + user.SecretKey
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = data
	cookie.Expires = time.Now().Add(2 * time.Minute)
	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, user)
}

func (userController *UserController) CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("user")
		if err != nil {
			return err
		}
		vals := strings.Split(cookie.Value, " ")
		user := &models.User{
			ID:        vals[0],
			SecretKey: vals[1],
		}

		dbUser := new(models.User)
		_ = userController.controller.DB.GormDB.First(dbUser, user.ID)

		if dbUser.SecretKey != user.SecretKey {
			return errors.New("secret_key is not valid")
		}

		return next(c)
	}
}

func (userController *UserController) GetActualID() (string, error) {
	var count int64
	userController.controller.DB.GormDB.Table("users").Select("count(distinct(id))").Count(&count)
 	count++

	return strconv.FormatInt(count, 10), nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
