package main

import (
	"flag"

	"github.com/kevincobain2000/instachart/pkg"
)

var (
	port string
)

func main() {
	cliArgs()
	e := pkg.NewEcho()

	pkg.GracefulServerWithPid(e, port)
}

func cliArgs() {
	flag.StringVar(&port, "port", "3000", "port to serve")
	flag.Parse()
}
