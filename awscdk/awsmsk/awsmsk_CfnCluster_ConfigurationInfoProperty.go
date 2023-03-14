package awsmsk


// Specifies the Amazon MSK configuration to use for the brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationInfoProperty := &configurationInfoProperty{
//   	arn: jsii.String("arn"),
//   	revision: jsii.Number(123),
//   }
//
type CfnCluster_ConfigurationInfoProperty struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, `arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1` .
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The revision of the Amazon MSK configuration to use.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

