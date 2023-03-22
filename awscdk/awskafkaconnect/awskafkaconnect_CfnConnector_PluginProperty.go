package awskafkaconnect


// A plugin is an AWS resource that contains the code that defines your connector logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginProperty := &pluginProperty{
//   	customPlugin: &customPluginProperty{
//   		customPluginArn: jsii.String("customPluginArn"),
//   		revision: jsii.Number(123),
//   	},
//   }
//
type CfnConnector_PluginProperty struct {
	// Details about a custom plugin.
	CustomPlugin interface{} `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

