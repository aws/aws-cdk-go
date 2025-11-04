package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
)

// Glue schema registry configuration for a Lambda event source.
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
//   // Your Glue Schema Registry
//   glueRegistry := awscdk.NewCfnRegistry(this, jsii.String("Registry"), &CfnRegistryProps{
//   	Name: jsii.String("schema-registry"),
//   	Description: jsii.String("Schema registry for event source"),
//   })
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String(Topic),
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(1),
//   		MaximumPollers: jsii.Number(3),
//   	},
//   	SchemaRegistryConfig: awscdk.NewGlueSchemaRegistry(&GlueSchemaRegistryProps{
//   		SchemaRegistry: glueRegistry,
//   		EventRecordFormat: lambda.EventRecordFormat_JSON(),
//   		SchemaValidationConfigs: []KafkaSchemaValidationConfig{
//   			&KafkaSchemaValidationConfig{
//   				Attribute: lambda.KafkaSchemaValidationAttribute_KEY(),
//   			},
//   		},
//   	}),
//   }))
//
type GlueSchemaRegistry interface {
	awslambda.ISchemaRegistry
	// Returns a schema registry configuration.
	Bind(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.KafkaSchemaRegistryConfig
}

// The jsii proxy struct for GlueSchemaRegistry
type jsiiProxy_GlueSchemaRegistry struct {
	internal.Type__awslambdaISchemaRegistry
}

func NewGlueSchemaRegistry(props *GlueSchemaRegistryProps) GlueSchemaRegistry {
	_init_.Initialize()

	if err := validateNewGlueSchemaRegistryParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueSchemaRegistry{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.GlueSchemaRegistry",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewGlueSchemaRegistry_Override(g GlueSchemaRegistry, props *GlueSchemaRegistryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.GlueSchemaRegistry",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GlueSchemaRegistry) Bind(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.KafkaSchemaRegistryConfig {
	if err := g.validateBindParameters(target, targetHandler); err != nil {
		panic(err)
	}
	var returns *awslambda.KafkaSchemaRegistryConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{target, targetHandler},
		&returns,
	)

	return returns
}

