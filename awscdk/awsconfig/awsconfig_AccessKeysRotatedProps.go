package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a AccessKeysRotated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inputParameters interface{}
//   var ruleScope ruleScope
//
//   accessKeysRotatedProps := &accessKeysRotatedProps{
//   	configRuleName: jsii.String("configRuleName"),
//   	description: jsii.String("description"),
//   	inputParameters: map[string]interface{}{
//   		"inputParametersKey": inputParameters,
//   	},
//   	maxAge: cdk.duration.minutes(jsii.Number(30)),
//   	maximumExecutionFrequency: awscdk.Aws_config.maximumExecutionFrequency_ONE_HOUR,
//   	ruleScope: ruleScope,
//   }
//
type AccessKeysRotatedProps struct {
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
	// The maximum number of days within which the access keys must be rotated.
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}

