package main

import (
	"buildpack/core"
	"buildpack/docker"
	"buildpack/usage"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var f = flag.NewFlagSet("nana", flag.ContinueOnError)

func main() {
	args := os.Args

	f.Usage = func() {
		_, _ = fmt.Fprint(f.Output(), usage.UsagePrefix)
		f.PrintDefaults()
		os.Exit(1)
	}
	f.PrintDefaults()
	if len(args) >= 2 {
		if args[1] == "version" {
			color.Cyan("Nana version " + usage.RunVersion)
			color.White(" Github: https://github.com/kildo162/nano")
			color.White(" Author: KhanhND(khanhnd162@gmail.com)")
			color.White(" Sponsors me a coffee: https://github.com/sponsors/kildo162")
			os.Exit(1)
			return
		} else if args[1] == "help" {
			f.Usage()
			f.PrintDefaults()
			os.Exit(1)
		} else if args[1] == "init" {
			core.CreateNanaFolder()
			core.CreateNanaYaml()
			color.Cyan("Nana initial successs")
			return
		} else if args[1] == "build" {
			docker.DockerBuild(args)
			os.Exit(1)
			return
		} else if args[1] == "clear" {
			docker.ClearLocal()
			return
		}
	}
}
