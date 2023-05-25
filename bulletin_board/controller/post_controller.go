package controller

import (
	"fmt"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
	"gin_test/bulletin_board/helper"
	service "gin_test/bulletin_board/service/post"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostController struct {
	tagsService service.PostsService
}

func NewPostsController(service service.PostsService) *PostController {
	return &PostController{
		tagsService: service,
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

func (controller *PostController) Update(ctx *gin.Context) {
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

	uerr := controller.tagsService.Update(updateTagsRequest)
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

// findAll controller
func (controller *PostController) FindAll(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("token")
	if err != nil || cookie.Value == "" {
		fmt.Print("No token")
		return
	}
	tagResponse := controller.tagsService.FindAll()
	// userName:=controller.tagsService.FindById(1)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"tags": tagResponse,
	})
}

func (controller *PostController) CreateForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create.html", gin.H{})
}

func (controller *PostController) UpdateForm(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tag := controller.tagsService.FindById(id)
	ctx.HTML(http.StatusOK, "update.html", gin.H{
		"Tag": tag,
	})
}
