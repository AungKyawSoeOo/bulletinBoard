package controller

import (
	"errors"
	"fmt"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/helper"
	service "gin_test/bulletin_board/service/auth"
	uservice "gin_test/bulletin_board/service/user"
	"gin_test/bulletin_board/utils"
	"os"

	"path/filepath"

	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type AuthController struct {
	AuthService service.Authservice
	UserService uservice.UserService
}

func NewAuthController(service service.Authservice, uservice uservice.UserService) *AuthController {
	return &AuthController{
		AuthService: service,
		UserService: uservice,
	}
}

func getCurrentUseID(ctx *gin.Context) (int, error) {
	cookie, err := ctx.Request.Cookie("token")
	if err != nil && err != http.ErrNoCookie {
		return 0, err
	}

	if cookie == nil {
		// Handle the case when the "token" cookie is not present
		// Return a default value for the user ID
		return 0, nil
	}

	tokenString := cookie.Value
	tokenSecret := os.Getenv("TOKEN_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userIDFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid user ID in token claims")
	}

	userID := int(userIDFloat)
	return userID, nil
}

// Register Controller
func (controller *AuthController) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	cpassword := ctx.PostForm("confirmPassword")
	phone := ctx.PostForm("phone")
	address := ctx.PostForm("address")
	dob := ctx.PostForm("dob")
	utype := ctx.PostForm("type")
	var userType string
	if utype == "1" {
		userType = "1"
	} else {
		userType = "0"
	}

	var dobTime *time.Time
	if dob != "" {
		parsedDOB, err := time.Parse("2006-01-02", dob)
		if err != nil {
			fmt.Print("Invalid date of birth")
		}
		dobTime = &parsedDOB
	}

	if cpassword != "" && password != cpassword {
		// Check the value of the "source" field
		source := ctx.PostForm("source")
		if source == "register" {
			// Redirect to the register form with the error
			ctx.Set("ConfirmPasswordError", "Passwords do not match.")
			ctx.HTML(http.StatusBadRequest, "register.html", gin.H{
				"Errors": map[string]string{
					"ConfirmPassword": "Passwords do not match.",
				},
			})
			return
		} else if source == "usercreateform" {
			// Redirect to the user create form with the error
			ctx.Set("ConfirmPasswordError", "Passwords do not match.")
			ctx.HTML(http.StatusBadRequest, "usercreateform.html", gin.H{
				"Errors": map[string]string{
					"ConfirmPassword": "Passwords do not match.",
				},
			})
			return
		}
	}

	// Check if email already exists
	existingUser := controller.AuthService.FindByEmail(email)
	if existingUser.Id != 0 {
		// Check the value of the "source" field
		source := ctx.PostForm("source")
		if source == "register" {
			// Redirect to the register form with the error
			ctx.Set("EmailExistsError", "Email already exists.")
			ctx.HTML(http.StatusBadRequest, "register.html", gin.H{
				"Errors": map[string]string{
					"EmailExists": "Email already exists.",
				},
			})
			return
		} else if source == "usercreateform" {
			userID, err := getCurrentUserID(ctx)
			if err != nil {
				ctx.Redirect(http.StatusFound, "/users")
				return
			}
			fmt.Print(userID)
			currentUser := controller.UserService.FindById(userID)
			// fmt.Print(currentUser)
			isLoggedIn := getIsLoggedIn(ctx)
			// Redirect to the user create form with the error
			ctx.Set("EmailExistsError", "Email already exists.")
			ctx.HTML(http.StatusBadRequest, "usercreateform.html", gin.H{
				"IsLoggedIn":  isLoggedIn,
				"CurrentUser": currentUser,
				"Errors": map[string]string{
					"EmailExists": "Email already exists.",
				},
			})
			return
		}
	}

	photoFile, err := ctx.FormFile("photo")
	if err != nil && err != http.ErrMissingFile {
		helper.ErrorPanic(err)
	}

	var photoPath string
	if photoFile != nil {
		// Generate a unique file name for the photo
		photoFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), photoFile.Filename)
		photoPath = filepath.Join("static", "images", photoFileName)

		err := ctx.SaveUploadedFile(photoFile, photoPath)
		if err != nil {
			helper.ErrorPanic(err)
		}

		// Convert backslashes to forward slashes
		photoPath = filepath.ToSlash(photoPath)
	}
	userID, err := getCurrentUseID(ctx)
	if err != nil {
		helper.ErrorPanic(err)
	}
	createUserRequest := request.CreateUserRequest{
		Username:        username,
		Email:           email,
		Password:        password,
		Phone:           phone,
		Address:         address,
		Date_Of_Birth:   dobTime,
		Type:            userType,
		Profile_Photo:   photoPath,
		Created_User_ID: userID,
	}
	aerr := controller.AuthService.Register(createUserRequest)
	if aerr != nil {
		if validationErr, ok := aerr.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)
			for _, fieldErr := range validationErr {
				fieldName := fieldErr.Field()
				errorMessage := ""
				switch fieldErr.Tag() {
				case "required":
					errorMessage = fieldName + " field is required"
				case "min":
					errorMessage = fieldName + " must be at least " + fieldErr.Param() + " characters long"
				case "max":
					errorMessage = fieldName + " must not exceed " + fieldErr.Param() + " characters"
				default:
					errorMessage = "Field validation failed"
				}
				errorMessages[fieldName] = errorMessage
				fmt.Println(errorMessages[fieldName])
			}
			source := ctx.PostForm("source")
			if source == "register" {
				ctx.HTML(http.StatusBadRequest, "register.html", gin.H{
					"Errors": errorMessages,
				})
				return
			} else if source == "usercreateform" {
				ctx.HTML(http.StatusBadRequest, "usercreateform.html", gin.H{
					"Errors": errorMessages,
				})
				return
			}
		}
	} else {
		source := ctx.PostForm("source")
		if source == "register" {
			ctx.Redirect(http.StatusFound, "/login")
		} else if source == "usercreateform" {
			ctx.Redirect(http.StatusFound, "/users")
		}
	}
}

// Login Controller
func (controller *AuthController) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	loginRequest := request.LoginRequest{
		Email:    email,
		Password: password,
	}
	token, err_token := controller.AuthService.Login(loginRequest)
	if err_token != nil {
		ctx.Set("logFail", "Invalid email or password.")
		ctx.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Errors": map[string]string{
				"logFail": "Invalid email or password.",
			},
		})
		return
		// helper.ResponseHandler(ctx, http.StatusBadRequest, "Invalid email or password.", nil)
		// return
	}

	// resp := response.LoginResponse{
	// 	TokenType: "Bearer",
	// 	Token:     token,
	// }
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}
	ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	ctx.Redirect(http.StatusFound, "/posts")
}

// Register Form
func (controller *AuthController) RegisterForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{})
}

// Login Form
func (controller *AuthController) LoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

// Logout

func (controller *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}

// Forget Password form
func (controller *AuthController) ForgetPasswordForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "forgetpassword.html", gin.H{})
}

// Forget Password
func (controller *AuthController) ForgetPassword(ctx *gin.Context) {
	tokenExpireInStr := os.Getenv("TOKEN_EXPIRED_IN")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	tokenDuration, err := time.ParseDuration(tokenExpireInStr)
	if err != nil {
		helper.ErrorPanic(err)
	}
	email := ctx.PostForm("email")
	existingUser := controller.AuthService.FindByEmail(email)
	if existingUser.Id == 0 {
		ctx.HTML(http.StatusOK, "forgetpassword.html", gin.H{
			"Errors": map[string]string{
				"NoEmail": "Email not exist.",
			},
		})
		return
	}
	// Generate a password reset token
	resetToken, err := utils.GenerateToken(tokenDuration, existingUser.Id, tokenSecret)
	if err != nil {
		// Handle the error appropriately
		fmt.Println("Failed to generate password reset token:", err)
		// Render an error message to the user
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"ErrorMessage": "Failed to generate password reset token.",
		})
		return
	}

	// Build the password reset URL
	resetURL := "http://localhost:8080/password_reset/" + resetToken + "/edit"

	// Compose the email
	subject := "Password Reset"
	body := "Hello,\n\nYou have requested to reset your password. Please click the link below to proceed:\n\n" + resetURL

	smtpServer := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	// Format the email
	message := "From: " + smtpUsername + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Send the email
	merr := smtp.SendMail(smtpServer+":"+smtpPort, smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer), smtpUsername, []string{email}, []byte(message))
	if merr != nil {
		// Handle the error appropriately
		fmt.Println("Failed to send email:", err)
		// Render an error message to the user
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"ErrorMessage": "Failed to send email.",
		})
		return
	}
	ctx.Redirect(http.StatusFound, "/forgetpassword")
}

func GeneratePasswordResetToken(ttl time.Duration, userId int, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256) // Use HMAC signing method

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = userId
	claim["exp"] = now.Add(ttl).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()
	claim["purpose"] = "password_reset" // Additional claim for password reset

	tokenString, err := token.SignedString([]byte(secretJWTKey)) // Convert string key to []byte

	if err != nil {
		return "", fmt.Errorf("generating password reset token failed: %w", err)
	}

	return tokenString, nil
}
