package main

import (
	"buildpack/docker"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var f = flag.NewFlagSet("nana", flag.ContinueOnError)

const (
	usagePrefix = `Usage: nana COMMAND [OPTIONS]
COMMAND:
  version       Showing version of nana.
  help          Showing usage.
  build         Build docker image for service(s) and push to registry.
Examples:
  nana version
  nana build all
Options:
  -c, --config  Current version file path
`
	runVersion = "v1.0.0"
)

func main() {
	args := os.Args

	f.Usage = func() {
		_, _ = fmt.Fprint(f.Output(), usagePrefix)
		f.PrintDefaults()
		os.Exit(1)
	}
	f.PrintDefaults()
	if len(args) >= 2 {
		if args[1] == "version" {
			color.Cyan("Nana version " + runVersion)
			color.White(" Github: https://github.com/kildo162/nano")
			color.White(" Author: KhanhND(khanhnd162@gmail.com)")
			color.White(" Sponsors me a coffee: https://github.com/sponsors/kildo162")
			os.Exit(1)
			return
		} else if args[1] == "help" {
			f.Usage()
			f.PrintDefaults()
			os.Exit(1)
		} else if args[1] == "build" {
			docker.DockerBuild(args)
			os.Exit(1)
			return
		} else if args[1] == "clear" {
			docker.ClearLocal()
		}
	}
}
