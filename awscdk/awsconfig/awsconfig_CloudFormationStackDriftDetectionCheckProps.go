package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a CloudFormationStackDriftDetectionCheck.
//
// Example:
//   // compliant if stack's status is 'IN_SYNC'
//   // non-compliant if the stack's drift status is 'DRIFTED'
//   // compliant if stack's status is 'IN_SYNC'
//   // non-compliant if the stack's drift status is 'DRIFTED'
//   config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"), &cloudFormationStackDriftDetectionCheckProps{
//   	ownStackOnly: jsii.Boolean(true),
//   })
//
type CloudFormationStackDriftDetectionCheckProps struct {
	// A name for the AWS Config rule.
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	InputParameters *map[string]interface{} `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	MaximumExecutionFrequency MaximumExecutionFrequency `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	RuleScope RuleScope `field:"optional" json:"ruleScope" yaml:"ruleScope"`
	// Whether to check only the stack where this rule is deployed.
	OwnStackOnly *bool `field:"optional" json:"ownStackOnly" yaml:"ownStackOnly"`
	// The IAM role to use for this rule.
	//
	// It must have permissions to detect drift
	// for AWS CloudFormation stacks. Ensure to attach `config.amazonaws.com` trusted
	// permissions and `ReadOnlyAccess` policy permissions. For specific policy permissions,
	// refer to https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

