package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
)

// Use a Kinesis Stream as a target for AWS CloudWatch event rules.
//
// Example:
//   // put to a Kinesis stream every time code is committed
//   // to a CodeCommit repository
//   repository.onCommit(jsii.String("onCommit"), &onCommitOptions{
//   	target: targets.NewKinesisStream(stream),
//   })
//
// Experimental.
type KinesisStream interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Kinesis Stream as a result from a CloudWatch event.
	// Experimental.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisStream
type jsiiProxy_KinesisStream struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewKinesisStream(stream awskinesis.IStream, props *KinesisStreamProps) KinesisStream {
	_init_.Initialize()

	j := jsiiProxy_KinesisStream{}

	_jsii_.Create(
		"monocdk.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisStream_Override(k KinesisStream, stream awskinesis.IStream, props *KinesisStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisStream) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

