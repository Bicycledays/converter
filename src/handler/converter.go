package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"sartorius/converter/src/service"
)

type Document struct {
	File      string `json:"file"`
	OutputDir string `json:"outputDir"`
}

func (h *Handler) convert(c *gin.Context) {
	var doc Document

	js, err := c.GetRawData()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "parsing body request error", err.Error())
		return
	}

	err = json.Unmarshal(js, &doc)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "parsing json params error", err.Error())
		return
	}

	converter := service.NewConverter(doc.File, doc.OutputDir)
	err = converter.Convert()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "convert error", err.Error())
		return
	}

	newResponse(c)
}
