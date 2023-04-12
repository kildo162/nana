# nana

I make small tool for developer when we work on localdev or other develop environment. Base on support for me when we work local.
nana support build/push docker image and manager version of docker image. Version is increase by 1 when we build docker image.

For use nana, we need install nana on local machine. We can install nana by download binary file or build from source code.
Please read user guide for more detail and how to use nana.

## User guide

### Install nana

#### Download binary file

Download binary file from [release](https://github.com/kildo162/nana/releases) page.

#### Build from source code

You can build nana from source code. We need install go and git on local machine.

```bash
git clone
cd nana
chmod +x build.sh
./build.sh
```

### Commands Available

#### Build

Build docker image and push to docker registry.

Nana will read config in file `versions.yaml` in current directory.
If you want change config file, you can use option `-c` or `--config` to set config file.

For build all docker image in config file, we can use command:

```bash
nana build all
```

For build a docker image, we can use command:

```bash
nana build <docker-image-name>
```

#### Version

Nana show current version.

```bash
nana version
```

#### Help

Nana show Usage.

```bash
nana help
```

#### Clear

Nana clear all images, containers, networks, volumes not use.

```bash
nana clear
```

### Config file

Config file is `versions.yaml` in current directory. You can view on example folder.

```yaml
  modules:
    - name: module1
      version: 1.0.10
      image: example
      path: /module1
    - name: module2
      version: 1.0.4
      image: example2
      tag: latest
      registry: kildo162/example/x/
      path: /module2
  registry:
    name: Github Container Registry
    endpoint: https://ghcr.io
    username: docker
    password: docker

```
