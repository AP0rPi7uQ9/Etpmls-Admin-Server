package v2

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
		core.JsonError(c, http.StatusBadRequest, core.ERROR_AttachmentUploadImage_Get_file, core.ERROR_MESSAGE_AttachmentUploadImage_Get_file, nil, err)
		return
	}

	extension, err := model.AttachmentValidateImageV2(file)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_AttachmentUploadImage_Validate, core.ERROR_MESSAGE_AttachmentUploadImage_Validate, nil, err)
		return
	}

	path, err := model.AttachmentUploadImageV2(c, file, extension)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_AttachmentUploadImage, core.ERROR_MESSAGE_AttachmentUploadImage, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_UploadImage, core.SUCCESS_MESSAGE_UploadImage, gin.H{"path": path})
	return
}

func AttachmentDeleteImage(c *gin.Context)  {
	var j model.ApiAttachmentDeleteImageV2
	if err := c.ShouldBindJSON(&j); err != nil{
		core.JsonError(c, http.StatusBadRequest, core.ERROR_AttachmentDeleteImage_Bind, core.ERROR_MESSAGE_AttachmentDeleteImage_Bind, nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_AttachmentDeleteImage_Validate, core.ERROR_MESSAGE_AttachmentDeleteImage_Validate, nil, err)
		return
	}
	// Validate Path
	err = model.AttachmentValidatePathV2(j.Path)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_AttachmentDeleteImage_Validate_path, core.ERROR_MESSAGE_AttachmentDeleteImage_Validate_path, nil, err)
		return
	}

	// Delete Image
	if err = model.AttachmentDeleteImageV2(j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_AttachmentDeleteImage, core.ERROR_MESSAGE_AttachmentDeleteImage, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_UploadImage, core.SUCCESS_MESSAGE_UploadImage, nil)
	return
}