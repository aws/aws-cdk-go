package previewawsdevopsagentmixins


// Properties for CfnAgentSpacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentSpaceMixinProps := &CfnAgentSpaceMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html
//
type CfnAgentSpaceMixinProps struct {
	// The description of the AgentSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the AgentSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

