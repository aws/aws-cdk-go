package awsmsk


// Specifies the configuration to use for the brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationInfoProperty := &ConfigurationInfoProperty{
//   	Arn: jsii.String("arn"),
//   	Revision: jsii.Number(123),
//   }
//
type CfnCluster_ConfigurationInfoProperty struct {
	// ARN of the configuration to use.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The revision of the configuration to use.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

