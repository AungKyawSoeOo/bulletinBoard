package main

import (
	"gin_test/bulletin_board/initializers"
	"gin_test/bulletin_board/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&model.User{}, &model.Posts{})
}
