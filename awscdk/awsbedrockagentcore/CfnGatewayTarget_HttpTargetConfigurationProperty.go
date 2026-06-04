package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpTargetConfigurationProperty := &HttpTargetConfigurationProperty{
//   	AgentcoreRuntime: &RuntimeTargetConfigurationProperty{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		Qualifier: jsii.String("qualifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httptargetconfiguration.html
//
type CfnGatewayTarget_HttpTargetConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-httptargetconfiguration.html#cfn-bedrockagentcore-gatewaytarget-httptargetconfiguration-agentcoreruntime
	//
	AgentcoreRuntime interface{} `field:"required" json:"agentcoreRuntime" yaml:"agentcoreRuntime"`
}

