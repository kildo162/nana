package docker

import (
	"buildpack/core"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Push(dockerImage string, registry *core.Registry) error {
	var authConfig = types.AuthConfig{
		Username:      registry.Username,
		Password:      registry.Password,
		ServerAddress: registry.Endpoint,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("NewClientWithOpts: ", err.Error())
		return err
	}
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}
	response, err := cli.ImagePush(ctx, dockerImage, opts)
	if err != nil {
		return err
	}

	str, err := DisplayDockerLog(response)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			//searching all dangling image then remove them
			//it may is not safe in case there are many build-processes are running in parallel
		}
		return fmtError(err, str)
	}

	return nil
}
