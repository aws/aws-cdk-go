package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

// Customize the Amazon Data Firehose Stream Event Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnDeliveryStream CfnDeliveryStream
//   var ruleTargetInput RuleTargetInput
//
//   kinesisFirehoseStream := awscdk.Aws_events_targets.NewKinesisFirehoseStream(cfnDeliveryStream, &KinesisFirehoseStreamProps{
//   	Message: ruleTargetInput,
//   })
//
// Deprecated: Use FirehoseDeliveryStream.
type KinesisFirehoseStream interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Firehose Stream as a result from a Event Bridge event.
	// Deprecated: Use FirehoseDeliveryStream.
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisFirehoseStream
type jsiiProxy_KinesisFirehoseStream struct {
	internal.Type__awseventsIRuleTarget
}

// Deprecated: Use FirehoseDeliveryStream.
func NewKinesisFirehoseStream(stream awskinesisfirehose.CfnDeliveryStream, props *KinesisFirehoseStreamProps) KinesisFirehoseStream {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseStreamParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseStream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisFirehoseStream",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Deprecated: Use FirehoseDeliveryStream.
func NewKinesisFirehoseStream_Override(k KinesisFirehoseStream, stream awskinesisfirehose.CfnDeliveryStream, props *KinesisFirehoseStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisFirehoseStream",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisFirehoseStream) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	if err := k.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

