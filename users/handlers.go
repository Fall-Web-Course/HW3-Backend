package users

import (
	"time"

	jwt "github.com/Fall-Web-Course/HW3/authorization"
	cache "github.com/Fall-Web-Course/HW3/cache"
	"github.com/Fall-Web-Course/HW3/db"
	"github.com/gin-gonic/gin"

	"errors"
	"strings"

	"math/rand"
	"net/http"
	"strconv"
)

var userList = []User{
	{Username: "admin", Password: "p@$$word", IsAdmin: true},
}

func isUserValid(username, password string) (bool, User) {
	var users []User
	db.GetDb().Find(&users)
	for _, u := range users {
		if u.Username == username && u.Password == password {
			// TODO: change
			return true, u
		}
	}
	return false, User{}
}

// Register a new user with the given username and password
// NOTE: For this demo, we
func registerNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := User{Username: username, Password: password, IsAdmin: false}
	userList = append(userList, u)

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}

func Login(c *gin.Context) {
	// Obtain the POSTed username and password values
	var user User
	c.BindJSON(&user)

	var sameSiteCookie http.SameSite
	// Check if the username/password combination is valid
	if valid, found_user := isUserValid(user.Username, user.Password); valid {
		// If the username/password is valid set the token in a cookie
		token := generateSessionToken()
		jwt_token, err := jwt.GenerateToken((uint64(found_user.ID)), found_user.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message":      "Login Failed",
				"ErrorMessage": "Can not generate token."})
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.SetSameSite(sameSiteCookie)
		c.Set("is_logged_in", true)

		cache.SetKeyWithDeadline(
			"access"+strconv.FormatUint(uint64(found_user.ID), 10),
			jwt_token.AccessToken,
			time.Unix(jwt_token.AccessExpire, 0))

		cache.SetKeyWithDeadline(
			"refresh"+strconv.FormatUint(uint64(found_user.ID), 10),
			jwt_token.RefreshToken,
			time.Unix(jwt_token.RefreshExpire, 0))

		c.JSON(http.StatusOK, gin.H{
			"title":         "Successful Login",
			"access_token":  jwt_token.AccessToken,
			"refresh_token": jwt_token.RefreshToken,
		})

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":      "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func logout(c *gin.Context) {

	var sameSiteCookie http.SameSite

	// Clear the cookie
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.SetSameSite(sameSiteCookie)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func Register(c *gin.Context) {
	// Obtain the POSTed username and password values from json
	var new_user User
	c.BindJSON(&new_user)

	var sameSiteCookie http.SameSite

	if registered_user, err := registerNewUser(new_user.Username, new_user.Password); err == nil {
		out := InsertToDb(new_user)
		if out.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message":      "Registration Failed",
				"ErrorMessage": "Duplicate username"})
			return
		}
		// If the user is created, set the token in a cookie and log the user in
		token := generateSessionToken()
		jwt_token, err := jwt.GenerateToken((uint64(registered_user.ID)), registered_user.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message":      "Login Failed",
				"ErrorMessage": "Can not generate token."})
		}

		c.SetCookie("token", token, 3600, "/", "", false, true)
		c.SetSameSite(sameSiteCookie)
		c.Set("is_logged_in", true)

		cache.SetKeyWithDeadline(
			"access"+strconv.FormatUint(uint64(registered_user.ID), 10),
			jwt_token.AccessToken,
			time.Unix(jwt_token.AccessExpire, 0))

		cache.SetKeyWithDeadline(
			"refresh"+strconv.FormatUint(uint64(registered_user.ID), 10),
			jwt_token.RefreshToken,
			time.Unix(jwt_token.RefreshExpire, 0))

		c.JSON(http.StatusOK, gin.H{
			"title":         "Successful registration and Login",
			"access_token":  jwt_token.AccessToken,
			"refresh_token": jwt_token.RefreshToken,
		})

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":      "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
