package main

import (
	"fmt"
	"os"

	coreControl "github.com/core-tools/hsu-core/go/control"
	coreDomain "github.com/core-tools/hsu-core/go/domain"
	coreLogging "github.com/core-tools/hsu-core/go/logging"
	"github.com/core-tools/hsu-echo-super-srv-go/internal/domain"
	"github.com/core-tools/hsu-echo-super-srv-go/internal/logging"
	echoControl "github.com/core-tools/hsu-echo/go/control"
	echoLogging "github.com/core-tools/hsu-echo/go/logging"

	flags "github.com/jessevdk/go-flags"
)

type flagOptions struct {
	Port int `long:"port" description:"port to listen on"`
}

func logPrefix(module string) string {
	return fmt.Sprintf("module: %s-server , ", module)
}

func main() {
	var opts flagOptions
	var argv []string = os.Args[1:]
	var parser = flags.NewParser(&opts, flags.HelpFlag)
	var err error
	_, err = parser.ParseArgs(argv)
	if err != nil {
		fmt.Printf("Command line flags parsing failed: %v", err)
		os.Exit(1)
	}

	logger := logging.NewSprintfLogger()

	logger.Infof("opts: %+v", opts)

	if opts.Port == 0 {
		fmt.Println("Port is required")
		os.Exit(1)
	}

	logger.Infof("Starting...")

	coreLogger := coreLogging.NewLogger(
		logPrefix("hsu-core"), coreLogging.LogFuncs{
			Debugf: logger.Debugf,
			Infof:  logger.Infof,
			Warnf:  logger.Warnf,
			Errorf: logger.Errorf,
		})
	echoLogger := echoLogging.NewLogger(
		logPrefix("hsu-echo"), echoLogging.LogFuncs{
			Debugf: logger.Debugf,
			Infof:  logger.Infof,
			Warnf:  logger.Warnf,
			Errorf: logger.Errorf,
		})

	coreServerOptions := coreControl.ServerOptions{
		Port: opts.Port,
	}
	coreServer, err := coreControl.NewServer(coreServerOptions, coreLogger)
	if err != nil {
		logger.Errorf("Failed to create local server: %v", err)
		return
	}

	coreServerHandler := coreDomain.NewDefaultHandler(coreLogger)
	echoServerHandler := domain.NewSuperHandler(echoLogger)

	echoControl.RegisterGRPCServerHandler(coreServer.GRPC(), echoServerHandler, echoLogger)
	coreControl.RegisterGRPCServerHandler(coreServer.GRPC(), coreServerHandler, coreLogger)

	coreServer.Run(nil)
}
