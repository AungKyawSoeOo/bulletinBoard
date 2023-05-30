package controller

import (
	"encoding/csv"
	"errors"
	"fmt"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/model"
	service "gin_test/bulletin_board/service/post"
	uservice "gin_test/bulletin_board/service/user"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
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

	var tagResponse []response.PostResponse
	if userID != 0 {
		currentUser := controller.userService.FindById(userID)
		if currentUser.Type == "1" {
			tagResponse = controller.tagsService.FindAll()
		} else {
			tagResponse = controller.tagsService.FindPostByUserId(userID)
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"tags":        tagResponse,
			"IsLoggedIn":  isLoggedIn,
			"CurrentUser": currentUser,
		})
		return
	}

	// If userID is 0 (no user logged in), retrieve all tags without currentUser
	tagResponse = controller.tagsService.FindAll()

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"tags":       tagResponse,
		"IsLoggedIn": isLoggedIn,
	})
}

// Create Form
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

// Upload csv form
func (controller *PostController) UploadForm(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}
	currentUser := controller.userService.FindById(userID)
	ctx.HTML(http.StatusOK, "upload.html", gin.H{
		"IsLoggedIn":  isLoggedIn,
		"CurrentUser": currentUser,
	})
}

// Update Form
func (controller *PostController) UpdateForm(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	currentUser := controller.userService.FindById(userID)

	// Retrieve the post ID from the URL parameter
	postID := ctx.Param("tagId")
	id, err := strconv.Atoi(postID)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/posts")
		return
	}

	post := controller.tagsService.FindById(id)

	if post.Id == 0 {
		ctx.Redirect(http.StatusFound, "/posts")
		return
	}

	if currentUser.Type != "1" {
		if userID != post.CreateUserId {
			ctx.Redirect(http.StatusFound, "/posts")
			return
		}
	}

	ctx.HTML(http.StatusOK, "update.html", gin.H{
		"Tag":         post,
		"IsLoggedIn":  isLoggedIn,
		"CurrentUser": currentUser,
	})

}

func (controller *PostController) UploadPosts(ctx *gin.Context, db *gorm.DB) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Error getting user Id: %s", err.Error()))
		return
	}

	currentUser := controller.userService.FindById(userID)

	// Retrieve the uploaded file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "upload.html", gin.H{
			"IsLoggedIn":  isLoggedIn,
			"CurrentUser": currentUser,
			"Errors": map[string]string{
				"Nofile": "Please choose a file.",
			},
		})
		return
	}

	// Check the file format
	if filepath.Ext(file.Filename) != ".csv" {
		ctx.HTML(http.StatusBadRequest, "upload.html", gin.H{
			"IsLoggedIn":  isLoggedIn,
			"CurrentUser": currentUser,
			"Errors": map[string]string{
				"FileFormat": "Invalid file format. Only CSV files are allowed.",
			},
		})
		return
	}

	// Open the uploaded file
	csvfile, err := file.Open()
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error opening file: %s", err.Error()))
		return
	}
	defer csvfile.Close()

	// Parse the CSV file
	reader := csv.NewReader(csvfile)
	records, err := reader.ReadAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error reading CSV: %s", err.Error()))
		return
	}

	// Save each row to the database
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < 3 || len(row) > 3 {
			ctx.HTML(http.StatusBadRequest, "upload.html", gin.H{
				"IsLoggedIn":  isLoggedIn,
				"CurrentUser": currentUser,
				"Errors": map[string]string{
					"RowCount": "Post upload csv must have 3 columns",
				},
			})
			return
		}
		status, err := strconv.Atoi(row[2])
		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error converting status: %s", err.Error()))
			return
		}

		post := model.Posts{
			Title:        row[0],
			Description:  row[1],
			Status:       status,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			CreateUserId: userID,
			UpdateUserId: userID,
			// Set the CreateUserId and UpdateUserId appropriately
		}

		// Save the post to the database using GORM
		if err := db.Create(&post).Error; err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error saving post: %s", err.Error()))
			return
		}
	}

	// Return a success message
	// ctx.String(http.StatusOK, "CSV file uploaded and processed successfully")
	ctx.Redirect(http.StatusFound, "/posts")
}

func (controller *PostController) DownloadPosts(ctx *gin.Context, db *gorm.DB) {
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		helper.ErrorPanic(err)
	}
	currentUser := controller.userService.FindById(userID)
	// Define the filter based on the user's role and ID
	var posts []model.Posts
	filter := db
	if currentUser.Type == "1" {
		// Admin can download all posts
		filter = filter.Find(&posts)
	} else {
		// User can only download their own posts
		filter = filter.Where("create_user_id = ?", userID).Find(&posts)
	}

	// Check for errors while retrieving the posts
	if err := filter.Error; err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error retrieving posts: %s", err.Error()))
		return
	}

	// Generate CSV content from the posts data
	csvData := [][]string{
		{"Id", "Title", "Description", "Status", "Created At", "Updated At", "Create User ID", "Update User ID"},
	}
	for _, post := range posts {
		csvData = append(csvData, []string{
			strconv.Itoa(post.Id),
			post.Title,
			post.Description,
			strconv.Itoa(post.Status),
			post.CreatedAt.Format("2006-01-02"),
			post.UpdatedAt.Format("2006-01-02"),
			strconv.Itoa(post.CreateUserId),
			strconv.Itoa(post.UpdateUserId),
		})
	}

	// Create a temporary file to write the CSV data
	tempFile, err := ioutil.TempFile("", "table_data_*.csv")
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error creating temporary file: %s", err.Error()))
		return
	}
	defer os.Remove(tempFile.Name())

	// Write the CSV data to the temporary file
	csvWriter := csv.NewWriter(tempFile)
	if err := csvWriter.WriteAll(csvData); err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error writing CSV data: %s", err.Error()))
		return
	}
	csvWriter.Flush()

	// Check for any errors during the CSV write
	if err := csvWriter.Error(); err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error writing CSV data: %s", err.Error()))
		return
	}

	// Read the contents of the temporary file
	fileContents, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error reading file contents: %s", err.Error()))
		return
	}

	// Set the appropriate headers for the response
	ctx.Header("Content-Disposition", "attachment; filename=table_data.csv")
	ctx.Data(http.StatusOK, "text/csv", fileContents)
}
