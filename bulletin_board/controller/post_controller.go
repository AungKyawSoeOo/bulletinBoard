package controller

import (
	"errors"
	"fmt"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
	"gin_test/bulletin_board/helper"
	service "gin_test/bulletin_board/service/post"
	uservice "gin_test/bulletin_board/service/user"
	"os"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type PostController struct {
	tagsService service.PostsService
	userService uservice.UserService
}

func NewPostsController(service service.PostsService, uservice uservice.UserService) *PostController {
	return &PostController{
		tagsService: service,
		userService: uservice,
	}
}

// create controller
func (controller *PostController) Create(ctx *gin.Context, userId int) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	createTagsRequest := request.CreatePostsRequest{
		Title:       title,
		Description: description,
		UserId:      userId,
	}
	fmt.Print(userId)
	fmt.Print(createTagsRequest)
	// ctx.HTML(http.StatusOK, "createConfirm.html", gin.H{})
	err := controller.tagsService.Create(createTagsRequest, userId)
	if err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
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
			}

			ctx.HTML(http.StatusBadRequest, "create.html", gin.H{
				"Errors": errorMessages,
			})
			return
		}
	} else {
		ctx.Redirect(http.StatusFound, "/posts")

	}
}

// update controller

func (controller *PostController) Update(ctx *gin.Context, userId int) {
	tagId := ctx.Param("tagId")
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	statusValue := ctx.PostForm("status")
	fmt.Print(statusValue, "staVa")
	fmt.Println(tagId)
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	if method := ctx.Request.Header.Get("X-HTTP-Method-Override"); method == "PUT" {
		ctx.Request.Method = "PUT"
	}
	updateTagsRequest := request.UpdatePostsRequest{
		Id:          id,
		Title:       title,
		Description: description,
	}
	if statusValue == "on" {
		status := 1
		updateTagsRequest.Status = &status
	} else {
		status := 0
		updateTagsRequest.Status = &status
	}
	fmt.Println(updateTagsRequest)
	// if err := ctx.ShouldBind(&updateTagsRequest); err != nil {
	// 	ctx.HTML(http.StatusBadRequest, "update.html", gin.H{
	// 		"Tag":    updateTagsRequest,
	// 		"Errors": err.Error(),
	// 	})
	// 	return
	// }

	uerr := controller.tagsService.Update(updateTagsRequest, userId)
	if uerr != nil {
		if validationErr, ok := uerr.(validator.ValidationErrors); ok {
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
			}

			ctx.HTML(http.StatusBadRequest, "update.html", gin.H{
				"Errors": errorMessages,
			})
			return
		}
	} else {
		ctx.Redirect(http.StatusFound, "/posts")
	}
}

// delete controller
func (controller *PostController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	// Check for method override header
	controller.tagsService.Delete(id)
	ctx.Redirect(http.StatusFound, "/posts")
}

// findById controller
func (controller *PostController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagsResponse := controller.tagsService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Login  check
func getIsLoggedIn(ctx *gin.Context) bool {
	isLoggedIn := false
	cookie, err := ctx.Request.Cookie("token")
	if err == nil && cookie.Value != "" {
		isLoggedIn = true
	}
	return isLoggedIn
}

func getCurrentUserID(ctx *gin.Context) (int, error) {
	cookie, err := ctx.Request.Cookie("token")
	if err != nil && err != http.ErrNoCookie {
		return 0, err
	}

	if cookie == nil {
		// Handle the case when the "token" cookie is not present
		// Return a default value for the user ID
		// ctx.Redirect(http.StatusFound, "/")
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

// findAll controller
func (controller *PostController) FindAll(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	fmt.Print(userID, err)
	if err != nil && userID == 0 {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	if userID != 0 {
		currentUser := controller.userService.FindById(userID)
		tagResponse := controller.tagsService.FindAll()
		// userName:=controller.tagsService.FindById(1)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"tags":        tagResponse,
			"IsLoggedIn":  isLoggedIn,
			"CurrentUser": currentUser,
		})
		return
	}
	tagResponse := controller.tagsService.FindAll()
	// userName:=controller.tagsService.FindById(1)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"tags":       tagResponse,
		"IsLoggedIn": isLoggedIn,
	})

	fmt.Print(isLoggedIn)

}

func (controller *PostController) CreateForm(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	currentUser := controller.userService.FindById(userID)

	ctx.HTML(http.StatusOK, "create.html", gin.H{
		"IsLoggedIn":  isLoggedIn,
		"CurrentUser": currentUser,
	})
}

func (controller *PostController) UpdateForm(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	currentUser := controller.userService.FindById(userID)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tag := controller.tagsService.FindById(id)
	ctx.HTML(http.StatusOK, "update.html", gin.H{
		"Tag":         tag,
		"IsLoggedIn":  isLoggedIn,
		"CurrentUser": currentUser,
	})
}
