package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// An event destination.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//   var myConfigurationSet ConfigurationSet
//
//
//   bus := events.EventBus_FromEventBusName(this, jsii.String("EventBus"), jsii.String("default"))
//
//   myConfigurationSet.AddEventDestination(jsii.String("ToEventBus"), &ConfigurationSetEventDestinationOptions{
//   	Destination: ses.EventDestination_EventBus(bus),
//   })
//
type EventDestination interface {
	// Use Event Bus as event destination.
	// Default: - do not send events to Event bus.
	//
	Bus() awsevents.IEventBus
	// A list of CloudWatch dimensions upon which to categorize your emails.
	// Default: - do not send events to CloudWatch.
	//
	Dimensions() *[]*CloudWatchDimension
	// Use Firehose Delivery Stream.
	// Default: - do not send events to Firehose Delivery Stream.
	//
	Stream() *FirehoseDeliveryStreamDestination
	// A SNS topic to use as event destination.
	// Default: - do not send events to a SNS topic.
	//
	Topic() awssns.ITopic
}

// The jsii proxy struct for EventDestination
type jsiiProxy_EventDestination struct {
	_ byte // padding
}

func (j *jsiiProxy_EventDestination) Bus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"bus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventDestination) Dimensions() *[]*CloudWatchDimension {
	var returns *[]*CloudWatchDimension
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventDestination) Stream() *FirehoseDeliveryStreamDestination {
	var returns *FirehoseDeliveryStreamDestination
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventDestination) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


func NewEventDestination_Override(e EventDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.EventDestination",
		nil, // no parameters
		e,
	)
}

// Use CloudWatch dimensions as event destination.
func EventDestination_CloudWatchDimensions(dimensions *[]*CloudWatchDimension) EventDestination {
	_init_.Initialize()

	if err := validateEventDestination_CloudWatchDimensionsParameters(dimensions); err != nil {
		panic(err)
	}
	var returns EventDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.EventDestination",
		"cloudWatchDimensions",
		[]interface{}{dimensions},
		&returns,
	)

	return returns
}

// Use Event Bus as event destination.
func EventDestination_EventBus(eventBus awsevents.IEventBus) EventDestination {
	_init_.Initialize()

	if err := validateEventDestination_EventBusParameters(eventBus); err != nil {
		panic(err)
	}
	var returns EventDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.EventDestination",
		"eventBus",
		[]interface{}{eventBus},
		&returns,
	)

	return returns
}

// Use Firehose Delivery Stream as event destination.
func EventDestination_FirehoseDeliveryStream(stream *FirehoseDeliveryStreamDestination) EventDestination {
	_init_.Initialize()

	if err := validateEventDestination_FirehoseDeliveryStreamParameters(stream); err != nil {
		panic(err)
	}
	var returns EventDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.EventDestination",
		"firehoseDeliveryStream",
		[]interface{}{stream},
		&returns,
	)

	return returns
}

// Use a SNS topic as event destination.
func EventDestination_SnsTopic(topic awssns.ITopic) EventDestination {
	_init_.Initialize()

	if err := validateEventDestination_SnsTopicParameters(topic); err != nil {
		panic(err)
	}
	var returns EventDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.EventDestination",
		"snsTopic",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

