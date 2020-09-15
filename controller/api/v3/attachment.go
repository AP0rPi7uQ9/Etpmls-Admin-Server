package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AttachmentUploadImage(c *gin.Context)  {
	file, err := c.FormFile("file")
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Get"), nil, err)
		return
	}

	var a model.Attachment
	extension, err := a.AttachmentValidateImage(file)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	path, err := a.AttachmentUploadImage(c, file, extension)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Upload"), nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Upload"), gin.H{"path": path})
	return
}

func AttachmentDeleteImage(c *gin.Context)  {
	var j model.ApiAttachmentDeleteImage
	if err := c.ShouldBindJSON(&j); err != nil{
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}
	// Validate Path
	var a model.Attachment
	err = a.AttachmentValidatePath(j.Path)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	// Delete Image
	if err = a.AttachmentDeleteImage(j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Delete"), nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Delete"), nil)
	return
}