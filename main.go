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

	pkg.GracefulServerWithPid(e, port)
}

func flags() {
	flag.StringVar(&port, "port", "3001", "port to serve")
	flag.StringVar(&baseURL, "baseURL", "/", "base url with slash")
	flag.Parse()
}
