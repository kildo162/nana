package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"html/template"
	"io"
	"log"
)

func Build(dockerfilePath string, imageName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("NewClientWithOpts: ", err.Error())
		return err
	}

	tarFile, err := archive.TarWithOptions(dockerfilePath, &archive.TarOptions{})
	if err != nil {
		return err
	}

	//fmt.Println("tarFile: ", tarFile)
	options := types.ImageBuildOptions{
		Dockerfile:  "Dockerfile",
		Tags:        []string{imageName},
		NoCache:     true,
		Remove:      true,
		ForceRemove: true,
	}

	response, err := cli.ImageBuild(ctx, tarFile, options)
	if err != nil {
		log.Println("ImageBuild: ", err.Error())
		return err
	}
	defer response.Body.Close()

	// Print the output
	str, err := DisplayDockerLog(response.Body)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			//searching all dangling image then remove them
			//it may is not safe in case there are many build-processes are running in parallel
		}
		return fmtError(err, str)
	}

	return nil
}

func fmtError(err error, msg string) error {
	type ErrTemp struct {
		Error  string
		Detail string
	}
	var ErrorDetail = `{{.Error}} {{.Detail}}`
	t := template.Must(template.New("error").Parse(ErrorDetail))
	var buf bytes.Buffer
	defer buf.Reset()
	e := t.Execute(&buf, ErrTemp{
		Error:  err.Error(),
		Detail: msg,
	})
	if e != nil {
		return err
	}
	return fmt.Errorf(buf.String())
}

func DisplayDockerLog(in io.Reader) (string, error) {
	var buf bytes.Buffer
	defer buf.Reset()
	var dec = json.NewDecoder(in)
	for {
		var jm jsonmessage.JSONMessage
		if err := dec.Decode(&jm); err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		if jm.Error != nil {
			return "", fmt.Errorf(jm.Error.Message)
		}
		if jm.Stream == "" {
			continue
		}
		buf.WriteString(jm.Stream)
	}
	return buf.String(), nil
}
