package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The format target function should recieve record in.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // Your MSK cluster arn
//   var clusterArn string
//
//   var myFunction function
//
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   secret := awscdk.NewSecret(this, jsii.String("Secret"), &SecretProps{
//   	SecretName: jsii.String("AmazonMSK_KafkaSecret"),
//   })
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String(Topic),
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(1),
//   		MaximumPollers: jsii.Number(3),
//   	},
//   	SchemaRegistryConfig: awscdk.NewConfluentSchemaRegistry(&ConfluentSchemaRegistryProps{
//   		SchemaRegistryUri: jsii.String("https://example.com"),
//   		EventRecordFormat: lambda.EventRecordFormat_JSON(),
//   		AuthenticationType: lambda.KafkaSchemaRegistryAccessConfigType_BASIC_AUTH(),
//   		Secret: secret,
//   		SchemaValidationConfigs: []kafkaSchemaValidationConfig{
//   			&kafkaSchemaValidationConfig{
//   				Attribute: lambda.KafkaSchemaValidationAttribute_KEY(),
//   			},
//   		},
//   	}),
//   }))
//
type EventRecordFormat interface {
	// The enum to use in `SchemaRegistryConfig.EventRecordFormat` property in CloudFormation.
	Value() *string
}

// The jsii proxy struct for EventRecordFormat
type jsiiProxy_EventRecordFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_EventRecordFormat) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A custom event record format.
func EventRecordFormat_Of(name *string) EventRecordFormat {
	_init_.Initialize()

	if err := validateEventRecordFormat_OfParameters(name); err != nil {
		panic(err)
	}
	var returns EventRecordFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EventRecordFormat",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func EventRecordFormat_JSON() EventRecordFormat {
	_init_.Initialize()
	var returns EventRecordFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.EventRecordFormat",
		"JSON",
		&returns,
	)
	return returns
}

func EventRecordFormat_SOURCE() EventRecordFormat {
	_init_.Initialize()
	var returns EventRecordFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.EventRecordFormat",
		"SOURCE",
		&returns,
	)
	return returns
}

