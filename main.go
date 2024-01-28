package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/kevincobain2000/instachart/pkg"
)

//go:embed all:frontend/dist/*
var publicDir embed.FS

type Flags struct {
	host                 string
	port                 string
	baseUrl              string
	allowedRemoteDomains string
	pprofHost            string
	pprofPort            string
}

var f Flags
var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version)
		return
	}
	SetupFlags()
	e := pkg.NewEcho(f.baseUrl, publicDir)

	pkg.GracefulServerWithPid(e, f.host, f.port)
}

func SetupFlags() {
	flag.StringVar(&f.host, "host", "localhost", "host to serve")
	flag.StringVar(&f.port, "port", "3001", "port to serve")
	flag.StringVar(&f.baseUrl, "base-url", "/", "base url with slash")
	flag.StringVar(&f.allowedRemoteDomains, "remote-domains", "", "csv list of allowed domains for remote fetching")
	flag.StringVar(&f.pprofHost, "pprof-host", "", "pprof host")
	flag.StringVar(&f.pprofPort, "pprof-port", "", "pprof port")
	flag.Parse()

	if f.pprofHost != "" && os.Getenv("PPROF_HOST") == "" {
		err := os.Setenv("PPROF_HOST", f.pprofHost)
		if err != nil {
			pkg.Logger().Error(err)
		}
	}
	if f.pprofPort != "" && os.Getenv("PPROF_PORT") == "" {
		err := os.Setenv("PPROF_PORT", f.pprofPort)
		if err != nil {
			pkg.Logger().Error(err)
		}
	}
	if f.allowedRemoteDomains != "" && os.Getenv("ALLOWED_REMOTE_DOMAINS") == "" {
		err := os.Setenv("ALLOWED_REMOTE_DOMAINS", f.allowedRemoteDomains)
		if err != nil {
			pkg.Logger().Error(err)
		}
	}
}
