package handlers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"rabietf.me/go-assignment/models"
	"rabietf.me/go-assignment/services"
)

type Login struct {
	Email    string
	Password string
}

// Util function to hash password before storing it, uses bcrypt.
// Returns "", error if something went wrong.
// Returngs hashed password, nil if everything goes normally.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Panic(err)
		return "", err
	}

	return string(bytes), nil
}

// Util function to verify that the password provided by user corresponds to the one in database.
// Returns false, error if the password is not correct.
// Returns true, nil if it matches.
func verifyPassword(userPassword, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	if err != nil {
		return err
	}

	return nil
}

// Util function to validate email using Regular expressions.
// Returns true if the email is in a valid email format.
func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

// POST request at /users, creates a new user account but does NOT authentificate him.
// 201 if successful.
// 500 if internal server during processing.
// 400 if user doesn't respect correct format.
func SignUp(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please send in JSON: name, email and password"})
		return
	}

	if !isEmailValid(newUser.Email) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter a correct email."})
		return
	}

	if len(newUser.Password) < 8 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Password too short, please user a password longer than 8 characters."})
		return
	}

	userExits, err := newUser.Find(newUser.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if userExits {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User already exists."})
		return
	}

	hashedPassword, err := hashPassword(newUser.Password)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	newUser.Password = hashedPassword

	id, err := newUser.Save()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"userId": id, "message": "You can now login!"})
	return

}

// POST request at /login, authentificates the user using JWT. takes email and password
// 200 if successful.
// 400 if request body incorrect
// 401 if wrong credentials.
func SignIn(c *gin.Context) {
	var newUser models.User
	var login Login
	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please send in JSON: email and password"})
		return
	}

	userExists, err := newUser.Find(login.Email)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !userExists {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User doesn't not exist."})
		return
	}

	err = verifyPassword(newUser.Password, login.Password)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Wrong password."})
		return
	}

	tokenString, err := services.CreateToken(newUser)

	c.Header("Authorization", "Bearer "+tokenString)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfuly connected! Welcome " + newUser.Name + "!"})

}
