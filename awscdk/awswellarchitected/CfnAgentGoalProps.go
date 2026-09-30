package awswellarchitected


// Properties for defining a `CfnAgentGoal`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentGoalProps := &CfnAgentGoalProps{
//   	Pillars: []*string{
//   		jsii.String("pillars"),
//   	},
//   	ProfileArn: jsii.String("profileArn"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html
//
type CfnAgentGoalProps struct {
	// The list of Well-Architected pillars this goal targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-pillars
	//
	Pillars *[]*string `field:"required" json:"pillars" yaml:"pillars"`
	// The Amazon Resource Name (ARN) of the parent Agent Profile that owns this goal.
	//
	// Pass `!Ref` of the parent AWS::WellArchitected::AgentProfile to flow its ARN here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-profilearn
	//
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
	// The title of the Agent Goal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-title
	//
	Title *string `field:"required" json:"title" yaml:"title"`
	// A description of the Agent Goal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentgoal.html#cfn-wellarchitected-agentgoal-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

