package service

import (
	"fmt"
	"log"
	"os/exec"
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

func NewConverter(file string, outputDir string) *Converter {
	cmd := exec.Command(
		"libreoffice",
		"--headless",
		"--convert-to",
		"pdf:calc_pdf_Export",
		file,
		"--outdir",
		outputDir,
	)

	return &Converter{cmd}
}

func (c *Converter) Convert() error {
	fmt.Println(c.Args)
	output, err := c.CombinedOutput()

	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		fmt.Println(string(output))
	}

	return nil
}
