package main

type plugin struct{}

var JavaScriptPlugin plugin

func (*plugin) Name() string {
	return "ICMP"
}

func (*plugin) Setup() error {
	return nil
}

func (*plugin) Teardown() error {
	return nil
}

func (*plugin) GetModules() map[string]interface{} {
	mods := map[string]interface{}{
		"icmp": New(),
	}
	return mods
}

func init() {
	JavaScriptPlugin = plugin{}
}
