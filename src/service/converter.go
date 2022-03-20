package service

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	varDir = "/app/var"
	pdfDir = varDir + "/upload/pdf"
)

/**
--convert-to pdf:writer_pdf_Export
--convert-to pdf:calc_pdf_Export
--convert-to pdf:draw_pdf_Export
--convert-to pdf:impress_pdf_Export
--convert-to pdf:writer_web_pdf_Export
*/

type Converter struct {
	*exec.Cmd
}

func NewConverter(file string) *Converter {
	path := varDir + file

	cmd := exec.Command(
		"libreoffice",
		"--headless",
		"--convert-to",
		"pdf:calc_pdf_Export",
		path,
		"--outdir",
		pdfDir,
	)

	return &Converter{cmd}
}

func (c *Converter) Convert() (convertedFile string, err error) {
	fmt.Println(c.Args)
	output, err := c.CombinedOutput()

	if err != nil {
		log.Println(err.Error())
		return "", err
	} else {
		fmt.Println(string(output))
	}

	convertedFile = c.Args[4] + ".pdf"
	return convertedFile, nil
}
