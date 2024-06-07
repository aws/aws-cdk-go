package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Construction properties for a CloudFormationStackNotificationCheck.
//
// Example:
//   // topics to which CloudFormation stacks may send event notifications
//   topic1 := sns.NewTopic(this, jsii.String("AllowedTopic1"))
//   topic2 := sns.NewTopic(this, jsii.String("AllowedTopic2"))
//
//   // non-compliant if CloudFormation stack does not send notifications to 'topic1' or 'topic2'
//   // non-compliant if CloudFormation stack does not send notifications to 'topic1' or 'topic2'
//   config.NewCloudFormationStackNotificationCheck(this, jsii.String("NotificationCheck"), &CloudFormationStackNotificationCheckProps{
//   	Topics: []iTopic{
//   		topic1,
//   		topic2,
//   	},
//   })
//
type CloudFormationStackNotificationCheckProps struct {
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
	// A list of allowed topics.
	//
	// At most 5 topics.
	// Default: - No topics.
	//
	Topics *[]awssns.ITopic `field:"optional" json:"topics" yaml:"topics"`
}

