package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Customize the Amazon Data Firehose Stream Event Target V2 to support L2 Amazon Data Firehose Delivery Stream instead of L1 Cfn Firehose Delivery Stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryStream iDeliveryStream
//   var ruleTargetInput ruleTargetInput
//
//   kinesisFirehoseStreamV2 := awscdk.Aws_events_targets.NewKinesisFirehoseStreamV2(deliveryStream, &KinesisFirehoseStreamProps{
//   	Message: ruleTargetInput,
//   })
//
type KinesisFirehoseStreamV2 interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Firehose Stream as a result from a Event Bridge event.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisFirehoseStreamV2
type jsiiProxy_KinesisFirehoseStreamV2 struct {
	internal.Type__awseventsIRuleTarget
}

func NewKinesisFirehoseStreamV2(stream IDeliveryStream, props *KinesisFirehoseStreamProps) KinesisFirehoseStreamV2 {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseStreamV2Parameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseStreamV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisFirehoseStreamV2",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

func NewKinesisFirehoseStreamV2_Override(k KinesisFirehoseStreamV2, stream IDeliveryStream, props *KinesisFirehoseStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisFirehoseStreamV2",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisFirehoseStreamV2) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
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

