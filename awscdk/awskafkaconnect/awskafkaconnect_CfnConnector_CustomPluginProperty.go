package awskafkaconnect


// A plugin is an AWS resource that contains the code that defines a connector's logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPluginProperty := &customPluginProperty{
//   	customPluginArn: jsii.String("customPluginArn"),
//   	revision: jsii.Number(123),
//   }
//
type CfnConnector_CustomPluginProperty struct {
	// The Amazon Resource Name (ARN) of the custom plugin.
	CustomPluginArn *string `field:"required" json:"customPluginArn" yaml:"customPluginArn"`
	// The revision of the custom plugin.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

