package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Construction properties for a CustomRule.
//
// Example:
//   // Lambda function containing logic that evaluates compliance with the rule.
//   evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &functionProps{
//   	code: lambda.assetCode.fromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   })
//
//   // A custom rule that runs on configuration changes of EC2 instances
//   customRule := config.NewCustomRule(this, jsii.String("Custom"), &customRuleProps{
//   	configurationChanges: jsii.Boolean(true),
//   	lambdaFunction: evalComplianceFn,
//   	ruleScope: config.ruleScope.fromResource(config.resourceType_EC2_INSTANCE()),
//   })
//
type CustomRuleProps struct {
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
	// The Lambda function to run.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Whether to run the rule on configuration changes.
	ConfigurationChanges *bool `field:"optional" json:"configurationChanges" yaml:"configurationChanges"`
	// Whether to run the rule on a fixed frequency.
	Periodic *bool `field:"optional" json:"periodic" yaml:"periodic"`
}

