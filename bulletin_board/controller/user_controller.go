package controller

import (
	"fmt"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/data/response"
	"gin_test/bulletin_board/helper"
	service "gin_test/bulletin_board/service/user"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UsersController struct {
	userService service.UserService
}

func NewUsercontroller(service service.UserService) *UsersController {
	return &UsersController{
		userService: service,
	}
}

// func getLoggedIn(ctx *gin.Context) bool {
// 	isLoggedIn := false
// 	cookie, err := ctx.Request.Cookie("token")
// 	if err == nil && cookie.Value != "" {
// 		isLoggedIn = true
// 	}
// 	return isLoggedIn
// }

// Find All
func (controller *UsersController) GetUsers(ctx *gin.Context, userRole string) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	currentUser := controller.userService.FindById(userID)
	var users []response.UserResponse
	if currentUser.Type == "1" {
		users = controller.userService.FindAll()
	} else {
		users = controller.userService.FindUserById(userID)
	}

	// helper.ResponseHandler(ctx, http.StatusOK, "Get All Users Success.", users)
	ctx.HTML(http.StatusOK, "userList.html", gin.H{
		"users":       users,
		"IsLoggedIn":  isLoggedIn,
		"CurrentUser": currentUser,
		"userRole":    userRole,
	})
}

// Delete
func (controller *UsersController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.userService.Delete(id)
	ctx.Redirect(http.StatusFound, "/users")
}

// Update
func (controller *UsersController) Update(ctx *gin.Context) {
	userId := ctx.Param("userId")
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	utype := ctx.PostForm("type")
	phone := ctx.PostForm("phone")
	dob := ctx.PostForm("dob")
	address := ctx.PostForm("address")
	var dobTime *time.Time
	if dob != "" {
		parsedDOB, err := time.Parse("2006-01-02", dob)
		if err != nil {
			fmt.Print("Invalid date of birth")
		}
		dobTime = &parsedDOB
	}
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	if method := ctx.Request.Header.Get("X-HTTP-Method-Override"); method == "PUT" {
		ctx.Request.Method = "PUT"
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

		// Save the uploaded file to the desired location
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
	updateUserRequest := request.UpdateUserRequest{
		Id:            id,
		Username:      username,
		Email:         email,
		Type:          utype,
		Phone:         phone,
		Address:       address,
		UpdateUserId:  userID,
		Date_Of_Birth: dobTime,
		Profile_Photo: photoPath,
	}
	controller.userService.Update(updateUserRequest)
	ctx.Redirect(http.StatusFound, "/users")

}

// user create form
func (controller *UsersController) CreateUser(ctx *gin.Context) {
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	currentUser := controller.userService.FindById(userID)
	isLoggedIn := getIsLoggedIn(ctx)
	ctx.HTML(http.StatusOK, "usercreateform.html", gin.H{
		"IsLoggedIn":  isLoggedIn,
		"CurrentUser": currentUser,
	})
}

// user update form
func (controller *UsersController) UpdateForm(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/users")
		return
	}
	currentUser := controller.userService.FindById(userID)
	fmt.Print(currentUser)
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/users")
		return
	}
	fmt.Print(id)
	user := controller.userService.FindById(id)
	fmt.Print(user.Id, "UUU")
	if user.Id == 0 {
		ctx.Redirect(http.StatusFound, "/users")
		return
	}
	fmt.Print(user.Id, "idddd")
	isProfileUpdate := (userID == user.Id)

	if currentUser.Type != "1" {
		if userID != user.Created_User_ID {
			if userID != user.Id {
				ctx.Redirect(http.StatusFound, "/users")
				return
			}
		}
	}

	ctx.HTML(http.StatusOK, "userupdate.html", gin.H{
		"User":            user,
		"IsUpdate":        true,
		"IsLoggedIn":      isLoggedIn,
		"CurrentUser":     currentUser,
		"IsProfileUpdate": isProfileUpdate,
	})
}

// user update form
func (controller *UsersController) ProfileForm(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/users")
		return
	}
	currentUser := controller.userService.FindById(userID)
	ctx.HTML(http.StatusOK, "userprofile.html", gin.H{
		"CurrentUser": currentUser,
		"IsLoggedIn":  isLoggedIn,
	})
}
func checkPassword(currentPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currentPassword))
	return err == nil
}

// change password form
func (controller *UsersController) ChangePasswordForm(ctx *gin.Context) {
	// var hasErrors bool
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/users")
		return
	}
	currentUser := controller.userService.FindById(userID)
	isLoggedIn := getIsLoggedIn(ctx)

	ctx.HTML(http.StatusOK, "changepasswordform.html", gin.H{
		"CurrentUser": currentUser,
		"IsLoggedIn":  isLoggedIn,
	})

}

// // Validate new password and confirm password
// if newPassword != confirmPassword {
// 	// New password and confirm password do not match, display an error message
// 	hasErrors = true
// 	ctx.HTML(http.StatusBadRequest, "changepasswordform.html", gin.H{
// 		"Errors": map[string]string{
// 			"ConfirmPassword": "Passwords do not match.",
// 		},
// 	})
// 	return
// }

// if hasErrors {
// 	ctx.HTML(http.StatusOK, "changepasswordform.html", gin.H{
// 		"CurrentUser": currentUser,
// 		"IsLoggedIn":  isLoggedIn,
// 	})
// 	return
// }

func (controller *UsersController) UpdatePassword(ctx *gin.Context) {
	isLoggedIn := getIsLoggedIn(ctx)
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/users")
		return
	}
	currentUser := controller.userService.FindById(userID)
	// userId := ctx.Param("userId")
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	utype := ctx.PostForm("type")
	currentpassword := ctx.PostForm("cpassword")
	password := ctx.PostForm("password")
	newcofirmpassword := ctx.PostForm("ncpassword")
	phone := ctx.PostForm("phone")
	dob := ctx.PostForm("dob")
	address := ctx.PostForm("address")

	// Validate current password
	fmt.Print(currentUser, "full")
	fmt.Print(currentpassword, "cuu")
	fmt.Print(currentUser.Password, "duu")
	if !checkPassword(currentpassword, currentUser.Password) {

		// Current password does not match, display an error message
		ctx.HTML(http.StatusOK, "changepasswordform.html", gin.H{
			"CurrentUser": currentUser,
			"IsLoggedIn":  isLoggedIn,
			"Errors": map[string]string{
				"CurrentPassword": "Old password is wrong.",
			},
		})
		return
	}

	if password != newcofirmpassword {
		// New password and confirm password do not match, display an error message
		// hasErrors = true
		ctx.HTML(http.StatusBadRequest, "changepasswordform.html", gin.H{
			"CurrentUser": currentUser,
			"IsLoggedIn":  isLoggedIn,
			"Errors": map[string]string{
				"ConfirmPassword": "Passwords do not match.",
			},
		})
		return
	}

	var dobTime *time.Time
	if dob != "" {
		parsedDOB, err := time.Parse("2006-01-02", dob)
		if err != nil {
			fmt.Print("Invalid date of birth")
		}
		dobTime = &parsedDOB
	}
	// id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	if method := ctx.Request.Header.Get("X-HTTP-Method-Override"); method == "PUT" {
		ctx.Request.Method = "PUT"
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

		// Save the uploaded file to the desired location
		err := ctx.SaveUploadedFile(photoFile, photoPath)
		if err != nil {
			helper.ErrorPanic(err)
		}

		// Convert backslashes to forward slashes
		photoPath = filepath.ToSlash(photoPath)
	}

	updateUserRequest := request.UpdateUserRequest{
		Id:            currentUser.Id,
		Username:      username,
		Email:         email,
		Password:      password,
		Type:          utype,
		Phone:         phone,
		Address:       address,
		UpdateUserId:  userID,
		Date_Of_Birth: dobTime,
		Profile_Photo: photoPath,
	}
	controller.userService.Update(updateUserRequest)
	ctx.Redirect(http.StatusFound, "/users")

	err = controller.userService.UpdatePassword(updateUserRequest)
	if err != nil {
		// Handle error and display an error message
		return
	}

	// Password updated successfully, redirect or display a success message
}
