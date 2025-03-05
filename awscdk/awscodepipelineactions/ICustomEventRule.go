package awscodepipelineactions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Represents a custom event rule in AWS CodePipeline Actions.
//
// This interface defines the structure for specifying a custom event rule
// in the AWS CodePipeline Actions module. The event rule is defined by an
// event pattern and a target.
// See: https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_events_targets-readme.html
//
type ICustomEventRule interface {
	// Description.
	Description() *string
	// event pattern when this rule should be triggered.
	EventPattern() *awsevents.EventPattern
	// Rulename.
	RuleName() *string
	// Target e.g. Lambda when event pattern is fulfilled.
	Target() awsevents.IRuleTarget
}

// The jsii proxy for ICustomEventRule
type jsiiProxy_ICustomEventRule struct {
	_ byte // padding
}

func (j *jsiiProxy_ICustomEventRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomEventRule) EventPattern() *awsevents.EventPattern {
	var returns *awsevents.EventPattern
	_jsii_.Get(
		j,
		"eventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomEventRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomEventRule) Target() awsevents.IRuleTarget {
	var returns awsevents.IRuleTarget
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

