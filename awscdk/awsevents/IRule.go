package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
)

// Represents an EventBridge Rule.
type IRule interface {
	awscdk.IResource
	// The value of the event rule Amazon Resource Name (ARN), such as arn:aws:events:us-east-2:123456789012:rule/example.
	RuleArn() *string
	// The name event rule.
	RuleName() *string
}

// The jsii proxy for IRule
type jsiiProxy_IRule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRule) RuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

