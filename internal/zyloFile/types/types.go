package types

type ZyloConfig struct {
	Image      string
	Workdir    string
	EnvVars    map[string]string
	Ports      []string
	Commands   []string
	Copies     map[string]string
	Entrypoint string
}
