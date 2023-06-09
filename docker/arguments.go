package docker

import (
	"buildpack/core"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func DockerBuild(args []string) {

	versionFilePathShot := flag.String("c", "", "Current version file path")
	versionFilePathLong := flag.String("config", "", "Current version file path")

	fileConfigPath := ""
	if *versionFilePathShot != "" {
		fileConfigPath = *versionFilePathShot
	} else if *versionFilePathLong != "" {
		fileConfigPath = *versionFilePathLong
	} else {
		fileConfigPath = "versions.yml"
	}

	flag.Parse()

	color.Cyan("Nana Init")
	data, err := core.GetFileVersionWithPath(fileConfigPath)
	if err != nil {
		fmt.Printf("%v %v\n", " => Load configuration file:", color.RedString("Failed"))
		fmt.Println(" => Cannot read versions.yml - " + err.Error())
		os.Exit(1)
		return
	} else {
		fmt.Printf("%v %v\n", " => Load configuration file:", color.GreenString("Success"))
		fmt.Printf("%v %v\n", " => File:", color.GreenString(fileConfigPath))
	}

	if len(args) >= 3 {

		if args[2] == "all" {
			color.Cyan("Build all service(s) in versions.yml")
			fmt.Println(" => Total service(s): " + strconv.Itoa(len(data.Modules)))

			for _, m := range data.Modules {
				buildOneModule(m, data)
			}

		} else {
			moduleName := args[2]
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

func buildOneModule(module core.Module, data *core.Data) {
	color.Cyan("Build service: " + module.Name + " (" + module.Name + ")")

	version, err := core.ParseVersion(module.Version)
	if err != nil {
		color.Red(" => Error parsing version: ", err)
		os.Exit(1)
		return
	}

	log.Printf(" => Current version: %v", version)
	versionNextVersionShot := flag.String("n", "", "Current version file path")
	versionNextVersionLong := flag.String("next", "", "Current version file path")
	flag.Parse()

	versionNextMinor := ""
	if *versionNextVersionShot != "" {
		versionNextMinor = *versionNextVersionShot
	} else if *versionNextVersionLong != "" {
		versionNextMinor = *versionNextVersionLong
	}

	if versionNextMinor == "minor" {
		version.NextMinor()
	} else if versionNextMinor == "major" {
		version.NextMajor()
	} else {
		version.NextPatch()
	}

	dockerImage := BuildDockerImageName(&module, &data.Registry, version)
	dockerFilePath := BuildDockerPath(&module)

	fmt.Println(" => Image name: " + dockerImage)

	start := time.Now()
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	if err := Build(dockerFilePath, dockerImage); err != nil {
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
	if err := Push(dockerImage, &data.Registry); err != nil {
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
