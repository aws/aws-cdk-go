package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
)

// Confluent schema registry configuration for a Lambda event source.
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
type ConfluentSchemaRegistry interface {
	awslambda.ISchemaRegistry
	// Returns a schema registry configuration.
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.KafkaSchemaRegistryConfig
}

// The jsii proxy struct for ConfluentSchemaRegistry
type jsiiProxy_ConfluentSchemaRegistry struct {
	internal.Type__awslambdaISchemaRegistry
}

func NewConfluentSchemaRegistry(props *ConfluentSchemaRegistryProps) ConfluentSchemaRegistry {
	_init_.Initialize()

	if err := validateNewConfluentSchemaRegistryParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfluentSchemaRegistry{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ConfluentSchemaRegistry",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewConfluentSchemaRegistry_Override(c ConfluentSchemaRegistry, props *ConfluentSchemaRegistryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ConfluentSchemaRegistry",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_ConfluentSchemaRegistry) Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.KafkaSchemaRegistryConfig {
	if err := c.validateBindParameters(_target, targetHandler); err != nil {
		panic(err)
	}
	var returns *awslambda.KafkaSchemaRegistryConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_target, targetHandler},
		&returns,
	)

	return returns
}

