package awswellarchitected


// Properties for CfnAgentGoalPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAgentGoalMixinProps := &CfnAgentGoalMixinProps{
//   	Description: jsii.String("description"),
//   	Pillars: []*string{
//   		jsii.String("pillars"),
//   	},
//   	ProfileArn: jsii.String("profileArn"),
//   	Title: jsii.String("title"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html
//
type CfnAgentGoalMixinProps struct {
	// A description of the Agent Goal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of Well-Architected pillars this goal targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-pillars
	//
	Pillars *[]*string `field:"optional" json:"pillars" yaml:"pillars"`
	// The Amazon Resource Name (ARN) of the parent Agent Profile that owns this goal.
	//
	// Pass `!Ref` of the parent AWS::WellArchitected::AgentProfile to flow its ARN here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-profilearn
	//
	ProfileArn *string `field:"optional" json:"profileArn" yaml:"profileArn"`
	// The title of the Agent Goal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
}

