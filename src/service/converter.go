package service

import (
	"log"
	"os/exec"
)

type Converter struct {
	*exec.Cmd
}

func NewConverter(outputFormat, filePath string) *Converter {
	cmd := exec.Command("libreoffice7.3", "--headless", "--convert-to", outputFormat, filePath)
	return &Converter{cmd}
}

func (c *Converter) Convert() (convertedFile string, err error) {
	output, err := c.CombinedOutput()

	if err != nil {
		log.Println(err.Error())
		return "", err
	} else {
		log.Println("output", string(output))
	}

	convertedFile = c.Args[5] + c.Args[4]
	return convertedFile, nil
}
