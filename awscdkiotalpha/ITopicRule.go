package awscdkiotalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2/internal"
)

// Represents an AWS IoT Rule.
// Experimental.
type ITopicRule interface {
	awscdk.IResource
	// The value of the topic rule Amazon Resource Name (ARN), such as arn:aws:iot:us-east-2:123456789012:rule/rule_name.
	// Experimental.
	TopicRuleArn() *string
	// The name topic rule.
	// Experimental.
	TopicRuleName() *string
}

// The jsii proxy for ITopicRule
type jsiiProxy_ITopicRule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITopicRule) TopicRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicRule) TopicRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleName",
		&returns,
	)
	return returns
}

