package api

import (
	playlistController "education/api/playlist_controller"
	uploadController "education/api/upload_controller"
	userController "education/api/user_controller"

	"github.com/gin-gonic/gin"
)

func RoutesPool(router *gin.Engine) {

	users := router.Group("/users")
	{
		users.GET("", userController.GetUsers)
		users.POST("", userController.CreateUsers)
		users.PUT("/:id", userController.EditUser)
		users.GET("/:id", userController.GetUser)
		users.DELETE("/:id", userController.DeleteUser)
	}

	uploads := router.Group("/uploads")
	{
		uploads.GET("", uploadController.GetUploads)
		uploads.POST("", uploadController.CreateUploads)
		uploads.PUT("/:id", uploadController.EditUpload)
		uploads.GET("/:id", uploadController.GetUpload)
		uploads.DELETE("/:id", uploadController.DeleteUpload)
	}

	playlists := router.Group("/playlists")
	{
		playlists.GET("", playlistController.GetPlaylists)
		playlists.POST("", playlistController.CreatePlaylist)
		playlists.PUT("/:id", playlistController.EditPlaylist)
		playlists.GET("/:id", playlistController.GetPlaylist)
		playlists.DELETE("/:id", playlistController.DeletePlaylist)
	}
}
