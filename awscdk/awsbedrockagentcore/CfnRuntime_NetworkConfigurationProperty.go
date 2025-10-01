package awsbedrockagentcore


// The network configuration for the agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	NetworkMode: jsii.String("networkMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-networkconfiguration.html
//
type CfnRuntime_NetworkConfigurationProperty struct {
	// The network mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-networkconfiguration.html#cfn-bedrockagentcore-runtime-networkconfiguration-networkmode
	//
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
}

