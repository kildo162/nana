package usage

const (
	UsageBuildPrefix = `Usage: nana build [COMMAND] [OPTIONS]
COMMAND:
  all 		 		 Build all services in versions.yaml file.
  [service name] 	 Build service name in versions.yaml file.
Examples:
  nana build all
  nana build module1
Options:
  -c, --config  Current version file path.
  -n, --next    Next version with auto increment for major, minor, patch. 
                Default is patch is not set.
  -h, --help    Show help for nana build command.
`
)
