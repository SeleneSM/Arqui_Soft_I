package controller

import (
	Dto "Arqui_Soft_I/dto"
	Service "Arqui_Soft_I/service"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving the file"})
		return
	}

	// Log para verificar el nombre del archivo recibido
	c.JSON(http.StatusOK, gin.H{"file_name": handler.Filename, "curso_id": c.PostForm("curso_id")})

	defer file.Close()

	courseIdStr := c.PostForm("curso_id")
	courseId, err := strconv.Atoi(courseIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el file"})
		return
	}

	fileDomain := Dto.ArchivoDto{Name: handler.Filename, Curso_id: courseId}

	newFile, apiErr := Service.CreateFile(fileDomain, fileContent)
	if apiErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    apiErr.Message(),
			"location": "service layer",
		})
		return
	}

	c.JSON(http.StatusOK, newFile)
}
