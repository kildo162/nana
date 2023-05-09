package usage

const (
	UsagePrefix = `Usage: nana COMMAND [OPTIONS]
COMMAND:
  init          Create config file in home directory of user.	
  version       Showing version of nana.
  help          Showing usage.
  build         Build docker image for service(s) and push to registry.
Examples:
  nana version
  nana build all
Options:
  -c, --config  Current version file path
`
	RunVersion = "v1.0.1"
)
