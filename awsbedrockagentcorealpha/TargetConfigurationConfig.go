package awsbedrockagentcorealpha


// Configuration returned by binding a target configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   targetConfigurationConfig := &TargetConfigurationConfig{
//   	Bound: jsii.Boolean(false),
//   }
//
// Experimental.
type TargetConfigurationConfig struct {
	// Indicates that the configuration has been successfully bound.
	// Experimental.
	Bound *bool `field:"required" json:"bound" yaml:"bound"`
}

