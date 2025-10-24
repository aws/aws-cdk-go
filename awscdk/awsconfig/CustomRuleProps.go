package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Construction properties for a CustomRule.
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
type CustomRuleProps struct {
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
	// The Lambda function to run.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Whether to run the rule on configuration changes.
	// Default: false.
	//
	ConfigurationChanges *bool `field:"optional" json:"configurationChanges" yaml:"configurationChanges"`
	// Whether to run the rule on a fixed frequency.
	// Default: false.
	//
	Periodic *bool `field:"optional" json:"periodic" yaml:"periodic"`
}

