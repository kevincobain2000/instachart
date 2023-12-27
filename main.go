package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kevincobain2000/instachart/pkg"
)

var (
	port string
)
var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version)
		return
	}
	flags()
	e := pkg.NewEcho()

	pkg.GracefulServerWithPid(e, port)
}

func flags() {
	flag.StringVar(&port, "port", "3001", "port to serve")
	flag.Parse()
}
