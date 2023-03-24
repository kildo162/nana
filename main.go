package main

import (
	"buildpack/core"
	"buildpack/docker"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
)

var f = flag.NewFlagSet("nana", flag.ContinueOnError)

var usagePrefix = `Usage: nana COMMAND [OPTIONS]
COMMAND:
  version       Showing version of buildpack
  help          Showing usage
Examples:
  nana version
  nana build all
Options:
`

func main() {
	argsWithProg := os.Args

	f.Usage = func() {
		_, _ = fmt.Fprint(f.Output(), usagePrefix)
		f.PrintDefaults()
		os.Exit(1)
	}

	if len(argsWithProg) >= 2 {
		if argsWithProg[1] == "version" {
			_, _ = fmt.Fprintln(f.Output(), "Nana version: 1.0.0")
			f.PrintDefaults()
			os.Exit(1)
		} else if argsWithProg[1] == "help" {
			f.Usage()
			f.PrintDefaults()
			os.Exit(1)
		} else if argsWithProg[1] == "build" {

			color.Cyan("Nana Init")

			data, err := core.GetFileVersion()
			if err != nil {
				fmt.Printf("%v %v\n", " => Load configuration file:", color.RedString("Failed"))
				fmt.Println(" => Cannot read versions.yml - " + err.Error())
				os.Exit(1)
				return
			} else {
				fmt.Printf("%v %v\n", " => Load configuration file:", color.GreenString("Success"))
			}

			if len(argsWithProg) >= 3 {
				if argsWithProg[2] == "all" {
					color.Cyan("Build all service(s) in versions.yml")
					fmt.Println(" => Total service(s): " + strconv.Itoa(len(data.Modules)))

					for _, m := range data.Modules {
						buildOneModule(m, data)
					}

				} else {
					moduleName := argsWithProg[2]
					var m core.Module
					for _, v := range data.Modules {
						if v.Name == moduleName {
							m = v
						}
					}

					if m.Name == "" {
						color.Red(" => Cannot find config module; Please check your module name")
						os.Exit(1)
						return
					}

					buildOneModule(m, data)
				}
			}

			os.Exit(1)
			return
		}
	}
}

func buildOneModule(module core.Module, data *core.Data) {
	color.Cyan("Build service: " + module.Name + " (" + module.Name + ")")

	version, err := core.ParseVersion(module.Version)
	if err != nil {
		color.Red(" => Error parsing version: ", err)
		os.Exit(1)
		return
	}

	version.NextPatch()
	dockerImage := docker.BuildDockerImageName(&module, &data.Registry, version)
	dockerFilePath := docker.BuildDockerPath(&module)

	fmt.Println(" => Image name: " + dockerImage)

	start := time.Now()
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	if err := docker.Build(dockerFilePath, dockerImage); err != nil {
		s.Stop()
		fmt.Printf("%v %v\n", " => Build service:", color.RedString("Failed"))
		fmt.Println(" => Error - " + err.Error())
		os.Exit(1)
		return
	} else {
		s.Stop()
		fmt.Printf("%v %v - %.3fs\n", " => Build service:", color.GreenString("Success"), time.Since(start).Seconds())
	}

	s.Stop()

	if err := core.UpdateVersion(module.Name, version); err != nil {
		fmt.Printf("%v %v\n", " => Update version:", color.RedString("Failed"))
		fmt.Println(" => Error - " + err.Error())
		os.Exit(1)
		return
	}

	s.Start()
	start = time.Now()
	if err := docker.Push(dockerImage, &data.Registry); err != nil {
		s.Stop()
		fmt.Printf("%v %v\n", " => Push image:", color.RedString("Failed"))
		fmt.Println(" => Error - " + err.Error())
		os.Exit(1)
		return
	} else {
		s.Stop()
		fmt.Printf("%v %v - %.3fs\n", " => Push image:", color.GreenString("Success"), time.Since(start).Seconds())
	}
}
