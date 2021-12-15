package model

type PluginsConfig struct {
	Plugins Plugins
}

type Plugins struct {
	InTree  []In
	OutTree []Out
}
type In struct {
	Name string
}
type Out struct {
	Name string
	Url  string
}
