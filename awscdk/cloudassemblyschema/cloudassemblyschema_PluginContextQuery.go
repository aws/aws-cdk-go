package cloudassemblyschema


// Query input for plugins.
//
// This alternate branch is necessary because it needs to be able to escape all type checking
// we do on on the cloud assembly -- we cannot know the properties that will be used a priori.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginContextQuery := &pluginContextQuery{
//   	pluginName: jsii.String("pluginName"),
//   }
//
type PluginContextQuery struct {
	// The name of the plugin.
	PluginName *string `field:"required" json:"pluginName" yaml:"pluginName"`
}

