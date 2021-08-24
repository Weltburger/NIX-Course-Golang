package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"task7/internal/controller"
	"task7/pkg/models"
	"time"
)

type Server struct {
	Router *echo.Echo
	Controller *controller.Controller
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:1323/callback",
		ClientID: "480864047918-rte9ogbch94e6mfu9c05h4g5s5v313ph.apps.googleusercontent.com",
		ClientSecret: "hTD6pKJ4hKIosY76y0POn64g",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	randomState = ""
)

func NewServer() *Server {
	server := &Server{
		Router: echo.New(),
		Controller: controller.New(),
	}

	server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())

	server.Router.GET("/", handleHome)
	server.Router.POST("/register", server.Controller.UserController().CreateUser)
	server.Router.POST("/login", server.Controller.UserController().LogIn)

	server.Router.GET("/oauth2/login", handleOAuth2Login)
	server.Router.GET("/callback", server.handleCallback)

	cookieGroup := server.Router.Group("/cookie")
	cookieGroup.Use(server.Controller.UserController().CheckAuthorization)
	cookieGroup.GET("/check_auth", cookiePage)

	cookieGroup.POST("/create/post", server.Controller.PostController().CreatePost)
	cookieGroup.GET("/get/post/:id", server.Controller.PostController().GetPost)
	cookieGroup.PUT("/update/post/:id", server.Controller.PostController().UpdatePost)
	cookieGroup.DELETE("/delete/post/:id", server.Controller.PostController().DeletePost)

	cookieGroup.POST("/create/comment", server.Controller.CommentController().CreateComment)
	cookieGroup.GET("/get/comment/:id", server.Controller.CommentController().GetComment)
	cookieGroup.PUT("/update/comment/:id", server.Controller.CommentController().UpdateComment)
	cookieGroup.DELETE("/delete/comment/:id", server.Controller.CommentController().DeleteComment)

	return server
}

func handleHome(c echo.Context) error {
	return c.HTML(http.StatusOK, `<html><body><a href="/oauth2/login">Google Log In </a></body></html>`)
}

func cookiePage(c echo.Context) error {
	return c.String(http.StatusOK, "good cookie")
}

func handleOAuth2Login(c echo.Context) error {
	randomState = controller.RandStringRunes(10)
	url := googleOauthConfig.AuthCodeURL(randomState)
	if err := c.Redirect(http.StatusTemporaryRedirect, url); err != nil {
		return err
	}
	return nil
}

func (s *Server) handleCallback(c echo.Context) error {
	if c.FormValue("state") != randomState {
		fmt.Println("State is not valid")
		if err := c.Redirect(http.StatusTemporaryRedirect, "/"); err != nil {
			return err
		}
	}

	token, err := googleOauthConfig.Exchange(context.Background(), c.FormValue("code"))
	if err != nil {
		fmt.Println("Could not get token")
		if err = c.Redirect(http.StatusTemporaryRedirect, "/"); err != nil {
			return err
		}
		return err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("Could not create get request")
		if err = c.Redirect(http.StatusTemporaryRedirect, "/"); err != nil {
			return err
		}
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not parse response")
		if err = c.Redirect(http.StatusTemporaryRedirect, "/"); err != nil {
			return err
		}
		return err
	}
	user := new(models.User)
	if err = json.Unmarshal(body, user); err != nil {
		return err
	}

	user.AuthID = user.ID
	checkUser := new(models.User)
	s.Controller.DB.GormDB.Where("auth_id = ?", user.AuthID).First(checkUser)

	if checkUser.ID != "" {
		data := user.ID + " " + user.SecretKey
		cookie := new(http.Cookie)
		cookie.Name = "user"
		cookie.Value = data
		cookie.Expires = time.Now().Add(2 * time.Minute)
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "you were logged in!",
		})
	}

	user.ID, err = s.Controller.UserController().GetActualID()
	user.Password = controller.RandStringRunes(8)
	user.SecretKey = controller.RandStringRunes(10)

	if err = s.Controller.UserController().CreateUserOAuth2(c, user); err != nil {
		fmt.Println("Could not create user")
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "you were registered!",
	})
}
