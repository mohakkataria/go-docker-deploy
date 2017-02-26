package payloads

type Deploy struct {
	Name string `form:"name"`
	Image     string   `form:"image"`
	Ports     string   `form:"ports"`
	EnvironmentVars     string   `form:"environment_vars"`
	Command     string   `form:"command"`
}

