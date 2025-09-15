package awscleanrooms


// A reference to a ConfiguredTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configuredTableReference := &ConfiguredTableReference{
//   	ConfiguredTableArn: jsii.String("configuredTableArn"),
//   	ConfiguredTableIdentifier: jsii.String("configuredTableIdentifier"),
//   }
//
type ConfiguredTableReference struct {
	// The ARN of the ConfiguredTable resource.
	ConfiguredTableArn *string `field:"required" json:"configuredTableArn" yaml:"configuredTableArn"`
	// The ConfiguredTableIdentifier of the ConfiguredTable resource.
	ConfiguredTableIdentifier *string `field:"required" json:"configuredTableIdentifier" yaml:"configuredTableIdentifier"`
}

