package awsbedrockagentcore


// Properties for defining a `CfnRuntimeEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuntimeEndpointProps := &CfnRuntimeEndpointProps{
//   	AgentRuntimeId: jsii.String("agentRuntimeId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AgentRuntimeVersion: jsii.String("agentRuntimeVersion"),
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtimeendpoint.html
//
type CfnRuntimeEndpointProps struct {
	// The agent runtime ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtimeendpoint.html#cfn-bedrockagentcore-runtimeendpoint-agentruntimeid
	//
	AgentRuntimeId *string `field:"required" json:"agentRuntimeId" yaml:"agentRuntimeId"`
	// The name of the AgentCore Runtime endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtimeendpoint.html#cfn-bedrockagentcore-runtimeendpoint-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtimeendpoint.html#cfn-bedrockagentcore-runtimeendpoint-agentruntimeversion
	//
	AgentRuntimeVersion *string `field:"optional" json:"agentRuntimeVersion" yaml:"agentRuntimeVersion"`
	// Contains information about an agent runtime endpoint.
	//
	// An agent runtime is the execution environment for a Amazon Bedrock Agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtimeendpoint.html#cfn-bedrockagentcore-runtimeendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags for the AgentCore Runtime endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtimeendpoint.html#cfn-bedrockagentcore-runtimeendpoint-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

