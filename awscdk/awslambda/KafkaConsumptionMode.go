package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How Lambda reads records from a Kafka topic.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction Function
//   var secret Secret
//
//
//   myFunction.AddEventSource(awscdk.NewSelfManagedKafkaEventSource(&SelfManagedKafkaEventSourceProps{
//   	BootstrapServers: []*string{
//   		jsii.String("kafka-broker-1:9092"),
//   		jsii.String("kafka-broker-2:9092"),
//   	},
//   	Topic: jsii.String("orders"),
//   	Secret: Secret,
//   	ConsumerGroupId: jsii.String("orders-queue"),
//   	ConsumptionMode: lambda.KafkaConsumptionMode_QUEUE(),
//   }))
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html
//
type KafkaConsumptionMode interface {
	// The value of the `ConsumptionMode` property.
	Mode() *string
}

// The jsii proxy struct for KafkaConsumptionMode
type jsiiProxy_KafkaConsumptionMode struct {
	_ byte // padding
}

func (j *jsiiProxy_KafkaConsumptionMode) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// A consumption mode that this library doesn't model.
func KafkaConsumptionMode_Of(mode *string) KafkaConsumptionMode {
	_init_.Initialize()

	if err := validateKafkaConsumptionMode_OfParameters(mode); err != nil {
		panic(err)
	}
	var returns KafkaConsumptionMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.KafkaConsumptionMode",
		"of",
		[]interface{}{mode},
		&returns,
	)

	return returns
}

func KafkaConsumptionMode_QUEUE() KafkaConsumptionMode {
	_init_.Initialize()
	var returns KafkaConsumptionMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaConsumptionMode",
		"QUEUE",
		&returns,
	)
	return returns
}

func KafkaConsumptionMode_STREAM() KafkaConsumptionMode {
	_init_.Initialize()
	var returns KafkaConsumptionMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaConsumptionMode",
		"STREAM",
		&returns,
	)
	return returns
}

