package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specific schema validation configuration settings that tell Lambda the message attributes you want to validate and filter using your schema registry.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // Your MSK cluster arn
//   var clusterArn string
//
//   var myFunction Function
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
//   		SchemaValidationConfigs: []KafkaSchemaValidationConfig{
//   			&KafkaSchemaValidationConfig{
//   				Attribute: lambda.KafkaSchemaValidationAttribute_KEY(),
//   			},
//   		},
//   	}),
//   }))
//
type KafkaSchemaValidationAttribute interface {
	// The enum to use in `SchemaRegistryConfig.SchemaValidationConfigs.Attribute` property in CloudFormation.
	Value() *string
}

// The jsii proxy struct for KafkaSchemaValidationAttribute
type jsiiProxy_KafkaSchemaValidationAttribute struct {
	_ byte // padding
}

func (j *jsiiProxy_KafkaSchemaValidationAttribute) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A custom schema validation attribute property.
func KafkaSchemaValidationAttribute_Of(name *string) KafkaSchemaValidationAttribute {
	_init_.Initialize()

	if err := validateKafkaSchemaValidationAttribute_OfParameters(name); err != nil {
		panic(err)
	}
	var returns KafkaSchemaValidationAttribute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.KafkaSchemaValidationAttribute",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func KafkaSchemaValidationAttribute_KEY() KafkaSchemaValidationAttribute {
	_init_.Initialize()
	var returns KafkaSchemaValidationAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaValidationAttribute",
		"KEY",
		&returns,
	)
	return returns
}

func KafkaSchemaValidationAttribute_VALUE() KafkaSchemaValidationAttribute {
	_init_.Initialize()
	var returns KafkaSchemaValidationAttribute
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaValidationAttribute",
		"VALUE",
		&returns,
	)
	return returns
}

