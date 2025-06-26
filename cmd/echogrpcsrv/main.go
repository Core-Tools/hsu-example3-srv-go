package main

import (
	"github.com/core-tools/hsu-echo-super-srv-go/internal/domain"
	echoControl "github.com/core-tools/hsu-echo/go/control"
)

func main() {
	echoControl.MainEcho(domain.NewSuperHandler)
}
