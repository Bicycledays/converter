package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"sartorius/converter/src/service"
)

type Document struct {
	Path         string `json:"path"`
	OutputFormat string `json:"outputFormat"`
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

	converter := service.NewConverter(doc.OutputFormat, doc.Path)
	convertedFile, err := converter.Convert()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "convert error", err.Error())
		return
	}

	newResultResponse(c, map[string]string{
		"file": convertedFile,
	})
}
