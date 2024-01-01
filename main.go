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

var (
	port    string
	host    string
	baseURL string
)
var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version)
		return
	}
	flags()
	e := pkg.NewEcho(baseURL, publicDir)

	pkg.GracefulServerWithPid(e, host, port)
}

func flags() {
	flag.StringVar(&host, "host", "localhost", "host to serve")
	flag.StringVar(&port, "port", "3001", "port to serve")
	flag.StringVar(&baseURL, "baseURL", "/", "base url with slash")
	flag.Parse()
}
