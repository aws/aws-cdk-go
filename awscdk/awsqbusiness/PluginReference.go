package awsqbusiness


// A reference to a Plugin resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginReference := &PluginReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	PluginArn: jsii.String("pluginArn"),
//   	PluginId: jsii.String("pluginId"),
//   }
//
type PluginReference struct {
	// The ApplicationId of the Plugin resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ARN of the Plugin resource.
	PluginArn *string `field:"required" json:"pluginArn" yaml:"pluginArn"`
	// The PluginId of the Plugin resource.
	PluginId *string `field:"required" json:"pluginId" yaml:"pluginId"`
}

