package awsconfig


// Construction properties for a ManagedRule.
//
// Example:
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   config.NewManagedRule(this, jsii.String("AccessKeysRotated"), &ManagedRuleProps{
//   	Identifier: config.ManagedRuleIdentifiers_ACCESS_KEYS_ROTATED(),
//   	InputParameters: map[string]interface{}{
//   		"maxAccessKeyAge": jsii.Number(60),
//   	},
//
//   	// default is 24 hours
//   	MaximumExecutionFrequency: config.MaximumExecutionFrequency_TWELVE_HOURS,
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

