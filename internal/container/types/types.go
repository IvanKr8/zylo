package types

import "time"

type Container struct {
	ID        string
	Name      string
	Status    string
	Pid       int
	CreatedAt time.Time
}

type ContainerConfig struct {
	Name      string
	Image     string
	Command   []string
	Volumes   []string
	Network   string
	Resources ResourcesConfig
}

type ResourcesConfig struct {
	CPU    int
	Memory int
}
