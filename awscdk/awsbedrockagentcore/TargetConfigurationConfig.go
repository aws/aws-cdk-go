package awsbedrockagentcore


// Configuration returned by binding a target configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetConfigurationConfig := &TargetConfigurationConfig{
//   	Bound: jsii.Boolean(false),
//   }
//
type TargetConfigurationConfig struct {
	// Indicates that the configuration has been successfully bound.
	Bound *bool `field:"required" json:"bound" yaml:"bound"`
}

