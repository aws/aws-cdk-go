package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Construction properties for a CustomRule.
//
// Example:
//   // Lambda function containing logic that evaluates compliance with the rule.
//   evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &FunctionProps{
//   	Code: lambda.AssetCode_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   })
//
//   // A custom rule that runs on configuration changes of EC2 instances
//   customRule := config.NewCustomRule(this, jsii.String("Custom"), &CustomRuleProps{
//   	ConfigurationChanges: jsii.Boolean(true),
//   	LambdaFunction: evalComplianceFn,
//   	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_INSTANCE()),
//   })
//
//   // A rule to detect stack drifts
//   driftRule := config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"))
//
//   // Topic to which compliance notification events will be published
//   complianceTopic := sns.NewTopic(this, jsii.String("ComplianceTopic"))
//
//   // Send notification on compliance change events
//   driftRule.onComplianceChange(jsii.String("ComplianceChange"), &OnEventOptions{
//   	Target: targets.NewSnsTopic(complianceTopic),
//   })
//
// Experimental.
type CustomRuleProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `field:"optional" json:"ruleScope" yaml:"ruleScope"`
	// The Lambda function to run.
	// Experimental.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Whether to run the rule on configuration changes.
	// Experimental.
	ConfigurationChanges *bool `field:"optional" json:"configurationChanges" yaml:"configurationChanges"`
	// Whether to run the rule on a fixed frequency.
	// Experimental.
	Periodic *bool `field:"optional" json:"periodic" yaml:"periodic"`
}

