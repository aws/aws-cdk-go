package awswellarchitected


// Properties for defining a `CfnAgentProfilePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentProfilePolicyProps := &CfnAgentProfilePolicyProps{
//   	Policy: jsii.String("policy"),
//   	ProfileArn: jsii.String("profileArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofilepolicy.html
//
type CfnAgentProfilePolicyProps struct {
	// The IAM resource-based policy document as a JSON string.
	//
	// Allow-only statements, no negation condition operators, maximum 51200 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofilepolicy.html#cfn-wellarchitected-agentprofilepolicy-policy
	//
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The Amazon Resource Name (ARN) of the parent Agent Profile the policy is attached to.
	//
	// Pass `!GetAtt` (or `!Ref`) of the parent AWS::WellArchitected::AgentProfile to flow its ARN here. Also serves as the primary identifier of this policy attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofilepolicy.html#cfn-wellarchitected-agentprofilepolicy-profilearn
	//
	ProfileArn *string `field:"required" json:"profileArn" yaml:"profileArn"`
}

