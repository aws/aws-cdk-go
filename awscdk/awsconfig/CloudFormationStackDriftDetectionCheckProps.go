package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// Construction properties for a CloudFormationStackDriftDetectionCheck.
//
// Example:
//   // compliant if stack's status is 'IN_SYNC'
//   // non-compliant if the stack's drift status is 'DRIFTED'
//   // compliant if stack's status is 'IN_SYNC'
//   // non-compliant if the stack's drift status is 'DRIFTED'
//   config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"), &CloudFormationStackDriftDetectionCheckProps{
//   	OwnStackOnly: jsii.Boolean(true),
//   })
//
type CloudFormationStackDriftDetectionCheckProps struct {
	// A name for the AWS Config rule.
	// Default: - CloudFormation generated name.
	//
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The modes the AWS Config rule can be evaluated in.
	//
	// The valid values are distinct objects.
	// Default: - Detective evaluation mode only.
	//
	EvaluationModes EvaluationMode `field:"optional" json:"evaluationModes" yaml:"evaluationModes"`
	// Input parameter values that are passed to the AWS Config rule.
	// Default: - No input parameters.
	//
	InputParameters *map[string]interface{} `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Default: MaximumExecutionFrequency.TWENTY_FOUR_HOURS
	//
	MaximumExecutionFrequency MaximumExecutionFrequency `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Default: - evaluations for the rule are triggered when any resource in the recording group changes.
	//
	RuleScope RuleScope `field:"optional" json:"ruleScope" yaml:"ruleScope"`
	// Whether to check only the stack where this rule is deployed.
	// Default: false.
	//
	OwnStackOnly *bool `field:"optional" json:"ownStackOnly" yaml:"ownStackOnly"`
	// The IAM role to use for this rule.
	//
	// It must have permissions to detect drift
	// for AWS CloudFormation stacks. Ensure to attach `config.amazonaws.com` trusted
	// permissions and `ReadOnlyAccess` policy permissions. For specific policy permissions,
	// refer to https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html.
	// Default: - A role will be created.
	//
	Role interfacesawsiam.IRoleRef `field:"optional" json:"role" yaml:"role"`
}

