package awscdkcloudassemblyschema


// Query input for plugins.
//
// This alternate branch is necessary because it needs to be able to escape all type checking
// we do on on the cloud assembly -- we cannot know the properties that will be used a priori.
type PluginContextQuery struct {
	// The name of the plugin.
	PluginName *string `field:"required" json:"pluginName" yaml:"pluginName"`
}

