package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/common/app"
	"github.com/rolexkdev/emvn-music-library-server/common/e"
	"github.com/rolexkdev/emvn-music-library-server/common/utils"
	"github.com/rolexkdev/emvn-music-library-server/dto"
)

// UploadFile godoc
//
//	@Summary		Upload files
//	@Description	upload files
//	@Tags			upload
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			files			formData	file	true	"Files to upload"
//	@Success		201				{object}	app.Response
//	@Failure		400				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/uploads [post]
func UploadFile(c *gin.Context) {
	appG := app.Gin{C: c}
	form, err := c.MultipartForm()
	if err != nil {
		appG.Response400(e.INVALID_PARAMS, "Invalid content-type, only accept form-data")
		return
	}

	fileHeaders := form.File["files"]
	if len(fileHeaders) == 0 {
		appG.Response400(e.INVALID_PARAMS, "File upload empty")
		return
	}

	uploadedFiles := make([]string, 0)

	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			appG.Response500(e.ERROR, "Error opening file")
			return
		}
		defer file.Close()

		// Save the file to the local file system
		filePath := filepath.Join(utils.UploadDir, fileHeader.Filename)
		if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
			appG.Response500(e.ERROR, "Error saving file")
			return
		}

		// Construct the URL to access the file
		// Hard code =((
		url := fmt.Sprintf("http://%s/api/v1/uploads/%s", c.Request.Host, fileHeader.Filename)
		uploadedFiles = append(uploadedFiles, url)
	}

	appG.Response201(dto.FileUploadResponse{
		FileURLs: uploadedFiles,
	})
}

// RetrieveFile godoc
//
//	@Summary		Retrieve file
//	@Description	Get a file by its filename
//	@Tags			upload
//	@Produce		plain
//	@Param			filename	path	string	true	"Filename"
//	@Success		200				{object}	app.Response
//	@Failure		404				{object}	app.Response
//	@Router			/uploads/{filename} [get]
func RetrieveFile(c *gin.Context) {
	appG := app.Gin{C: c}
	filename := c.Param("filename")
	filePath := filepath.Join(utils.UploadDir, filename)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		appG.Response404(e.NOTFOUND, gin.H{"message": "File not found"})
		return
	}

	// Set the content type based on the file extension
	contentType := utils.GetFileContentType(filename)
	c.Writer.Header().Set("Content-Type", contentType)

	// Serve the file
	c.File(filePath)
}
