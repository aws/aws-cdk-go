package awsconfig


// Construction properties for a new rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inputParameters interface{}
//   var ruleScope ruleScope
//
//   ruleProps := &RuleProps{
//   	ConfigRuleName: jsii.String("configRuleName"),
//   	Description: jsii.String("description"),
//   	InputParameters: map[string]interface{}{
//   		"inputParametersKey": inputParameters,
//   	},
//   	MaximumExecutionFrequency: awscdk.Aws_config.MaximumExecutionFrequency_ONE_HOUR,
//   	RuleScope: ruleScope,
//   }
//
// Experimental.
type RuleProps struct {
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
}

