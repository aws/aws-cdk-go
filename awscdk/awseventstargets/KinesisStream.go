package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// Use a Kinesis Stream as a target for AWS CloudWatch event rules.
//
// Example:
//   // put to a Kinesis stream every time code is committed
//   // to a CodeCommit repository
//   repository.onCommit(jsii.String("onCommit"), &OnCommitOptions{
//   	Target: targets.NewKinesisStream(stream),
//   })
//
type KinesisStream interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Kinesis Stream as a result from a CloudWatch event.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisStream
type jsiiProxy_KinesisStream struct {
	internal.Type__awseventsIRuleTarget
}

func NewKinesisStream(stream awskinesis.IStream, props *KinesisStreamProps) KinesisStream {
	_init_.Initialize()

	if err := validateNewKinesisStreamParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisStream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

func NewKinesisStream_Override(k KinesisStream, stream awskinesis.IStream, props *KinesisStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisStream) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := k.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

