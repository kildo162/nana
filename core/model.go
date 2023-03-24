package core

type Data struct {
	Modules  []Module `yaml:"modules"`
	Registry Registry `yaml:"registry"`
}

type Module struct {
	Name     string `yaml:"name,omitempty"`
	Version  string `yaml:"version,omitempty"`
	Image    string `yaml:"image,omitempty"`
	Tag      string `yaml:"tag,omitempty"`
	Build    bool   `yaml:"build,omitempty"`
	Registry string `yaml:"registry,omitempty"`
	Path     string `yaml:"path,omitempty"`
	Look     bool   `yaml:"look,omitempty"`
}

type Registry struct {
	Name     string `yaml:"name,omitempty"`
	Endpoint string `yaml:"endpoint,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Access   bool   `yaml:"access,omitempty"`
}
