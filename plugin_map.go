package l9explore

import (
	"github.com/LeakIX/l9format"
	// Import your plugins here
	"github.com/LeakIX/l9plugins"
)

var TcpPlugins []l9format.ServicePluginInterface
var WebPlugins []l9format.WebPluginInterface

func LoadL9ExplorePlugins() {
	TcpPlugins = append(TcpPlugins, l9plugins.GetTcpPlugins()...)
	WebPlugins = append(WebPlugins, l9plugins.GetWebPlugins()...)
	// Add your plugins here
}

