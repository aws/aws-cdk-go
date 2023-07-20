package awskafkaconnect


// A plugin is an AWS resource that contains the code that defines your connector logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginProperty := &PluginProperty{
//   	CustomPlugin: &CustomPluginProperty{
//   		CustomPluginArn: jsii.String("customPluginArn"),
//   		Revision: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-plugin.html
//
type CfnConnector_PluginProperty struct {
	// Details about a custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-plugin.html#cfn-kafkaconnect-connector-plugin-customplugin
	//
	CustomPlugin interface{} `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

