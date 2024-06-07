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
//   var evaluationMode evaluationMode
//   var inputParameters interface{}
//   var ruleScope ruleScope
//
//   accessKeysRotatedProps := &AccessKeysRotatedProps{
//   	ConfigRuleName: jsii.String("configRuleName"),
//   	Description: jsii.String("description"),
//   	EvaluationModes: evaluationMode,
//   	InputParameters: map[string]interface{}{
//   		"inputParametersKey": inputParameters,
//   	},
//   	MaxAge: cdk.Duration_Minutes(jsii.Number(30)),
//   	MaximumExecutionFrequency: awscdk.Aws_config.MaximumExecutionFrequency_ONE_HOUR,
//   	RuleScope: ruleScope,
//   }
//
type AccessKeysRotatedProps struct {
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
	// The maximum number of days within which the access keys must be rotated.
	// Default: Duration.days(90)
	//
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}

