package docker

import (
	"buildpack/core"
	"log"
	"os"
	"regexp"
	"strings"
)

func BuildDockerImageName(module *core.Module, registry *core.Registry, version core.Version) string {
	endpoint := ""
	if len(module.Registry) > 0 {
		endpoint = module.Registry
	} else {
		endpoint = registry.Endpoint
	}

	compile := regexp.MustCompile(`^https?://`)
	r := compile.ReplaceAllString(endpoint, "")
	if !strings.HasSuffix(r, "/") {
		r = r + "/"
	}

	if len(module.Tag) > 0 {
		return r + module.Image + ":" + version.String() + "-" + module.Tag
	}
	return r + module.Image + ":" + version.String()
}

func BuildDockerPath(module *core.Module) string {
	if len(module.Path) > 0 {
		return getPath() + module.Path
	} else {
		return getPath() + "/" + module.Name
	}
}

func getPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
