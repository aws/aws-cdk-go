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
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   stream := firehose.NewDeliveryStream(this, jsii.String("DeliveryStream"), &DeliveryStreamProps{
//   	Destination: firehose.NewS3Bucket(bucket),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Expression(jsii.String("rate(1 minute)")),
//   })
//   rule.AddTarget(targets.NewFirehoseDeliveryStream(stream))
//
type FirehoseDeliveryStream interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Firehose Stream as a result from a Event Bridge event.
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for FirehoseDeliveryStream
type jsiiProxy_FirehoseDeliveryStream struct {
	internal.Type__awseventsIRuleTarget
}

func NewFirehoseDeliveryStream(deliveryStream awskinesisfirehose.IDeliveryStream, props *FirehoseDeliveryStreamProps) FirehoseDeliveryStream {
	_init_.Initialize()

	if err := validateNewFirehoseDeliveryStreamParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehoseDeliveryStream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.FirehoseDeliveryStream",
		[]interface{}{deliveryStream, props},
		&j,
	)

	return &j
}

func NewFirehoseDeliveryStream_Override(f FirehoseDeliveryStream, deliveryStream awskinesisfirehose.IDeliveryStream, props *FirehoseDeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.FirehoseDeliveryStream",
		[]interface{}{deliveryStream, props},
		f,
	)
}

func (f *jsiiProxy_FirehoseDeliveryStream) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	if err := f.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

