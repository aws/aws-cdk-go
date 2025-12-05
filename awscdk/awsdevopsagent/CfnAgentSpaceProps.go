package awsdevopsagent


// Properties for defining a `CfnAgentSpace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentSpaceProps := &CfnAgentSpaceProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html
//
type CfnAgentSpaceProps struct {
	// The name of the AgentSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the AgentSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

