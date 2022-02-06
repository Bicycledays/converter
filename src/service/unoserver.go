package service

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"syscall"
	"time"
)

const (
	statusSleep = 0
	statusWait  = 1
	statusWork  = 2
)

type Unoserver struct {
	port   int
	status int
}

func (u *Unoserver) Status() int {
	return u.status
}

func (u *Unoserver) SetStatus(status int) {
	u.status = status
}

func (u *Unoserver) Port() int {
	return u.port
}

func (u *Unoserver) SetPort(port int) {
	u.port = port
}

func (u *Unoserver) Start(ch chan byte) {
	// todo CommandContext https://pkg.go.dev/os/exec#CommandContext
	cmd := exec.Command("unoserver", "--port", fmt.Sprintf("%d", u.port))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var isWorking bool
	//_, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	go func() {
		err := cmd.Start()
		var writer = cmd.String()
		log.Print(writer)
		//log.Print(out)
		if err != nil {
			log.Print(err.Error())
		}
		pgid, err := syscall.Getpgid(cmd.Process.Pid)
		isWorking = true

		log.Print("Server working")
		<-ch
		_ = syscall.Kill(-pgid, 15)
		log.Print("Finish kill")
		if err != nil {
			log.Print(err.Error())
		}
		ch <- 2
	}()
	time.Sleep(time.Second * 10)
	for isWorking == false {

	}
	log.Print("Change status")
	u.SetStatus(statusWait)
}

func (u *Unoserver) Convert(filepath, convertTo string) {
	defer u.SetStatus(statusWait)
	u.SetStatus(statusWork)

	var reg = regexp.MustCompile(`\.[^\.]+$`)
	convertedFile := reg.ReplaceAllString(filepath, "."+convertTo)

	cmd := exec.Command(
		"unoconvert",
		"--convert-to",
		convertTo,
		"--port",
		fmt.Sprintf("%d", u.port),
		filepath,
		convertedFile,
	)
	err := cmd.Start()
	log.Print("Start convert")
	log.Print(convertTo)
	log.Print(fmt.Sprintf("%d", u.port))
	log.Print(filepath)
	log.Print(convertedFile)
	if err != nil {
		log.Print(err.Error())
	}
	err = cmd.Wait()
	log.Print("End convert")
	if err != nil {
		log.Print(err.Error())
	}
}
