package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of authentication protocol for your schema registry.
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
type KafkaSchemaRegistryAccessConfigType interface {
	// The key to use in `SchemaRegistryConfig.AccessConfig.Type` property in CloudFormation.
	Type() *string
}

// The jsii proxy struct for KafkaSchemaRegistryAccessConfigType
type jsiiProxy_KafkaSchemaRegistryAccessConfigType struct {
	_ byte // padding
}

func (j *jsiiProxy_KafkaSchemaRegistryAccessConfigType) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// A custom source access configuration property for schema registry.
func KafkaSchemaRegistryAccessConfigType_Of(name *string) KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()

	if err := validateKafkaSchemaRegistryAccessConfigType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns KafkaSchemaRegistryAccessConfigType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func KafkaSchemaRegistryAccessConfigType_BASIC_AUTH() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"BASIC_AUTH",
		&returns,
	)
	return returns
}

func KafkaSchemaRegistryAccessConfigType_CLIENT_CERTIFICATE_TLS_AUTH() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"CLIENT_CERTIFICATE_TLS_AUTH",
		&returns,
	)
	return returns
}

func KafkaSchemaRegistryAccessConfigType_SASL_SCRAM_256_AUTH() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"SASL_SCRAM_256_AUTH",
		&returns,
	)
	return returns
}

func KafkaSchemaRegistryAccessConfigType_SASL_SCRAM_512_AUTH() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"SASL_SCRAM_512_AUTH",
		&returns,
	)
	return returns
}

func KafkaSchemaRegistryAccessConfigType_SERVER_ROOT_CA_CERTIFICATE() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"SERVER_ROOT_CA_CERTIFICATE",
		&returns,
	)
	return returns
}

func KafkaSchemaRegistryAccessConfigType_VIRTUAL_HOST() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"VIRTUAL_HOST",
		&returns,
	)
	return returns
}

func KafkaSchemaRegistryAccessConfigType_VPC_SECURITY_GROUP() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"VPC_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func KafkaSchemaRegistryAccessConfigType_VPC_SUBNET() KafkaSchemaRegistryAccessConfigType {
	_init_.Initialize()
	var returns KafkaSchemaRegistryAccessConfigType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.KafkaSchemaRegistryAccessConfigType",
		"VPC_SUBNET",
		&returns,
	)
	return returns
}

