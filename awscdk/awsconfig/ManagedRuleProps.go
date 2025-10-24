package awsconfig


// Construction properties for a ManagedRule.
//
// Example:
//   var fn Function
//   var samplePolicyText string
//
//
//   config.NewManagedRule(this, jsii.String("ManagedRule"), &ManagedRuleProps{
//   	Identifier: config.ManagedRuleIdentifiers_API_GW_XRAY_ENABLED(),
//   	EvaluationModes: config.EvaluationMode_DETECTIVE_AND_PROACTIVE(),
//   })
//
//   config.NewCustomRule(this, jsii.String("CustomRule"), &CustomRuleProps{
//   	LambdaFunction: fn,
//   	EvaluationModes: config.EvaluationMode_PROACTIVE(),
//   })
//
//   config.NewCustomPolicy(this, jsii.String("CustomPolicy"), &CustomPolicyProps{
//   	PolicyText: samplePolicyText,
//   	EvaluationModes: config.EvaluationMode_DETECTIVE(),
//   })
//
type ManagedRuleProps struct {
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
	// The identifier of the AWS managed rule.
	// See: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

