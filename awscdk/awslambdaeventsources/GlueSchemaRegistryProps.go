package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for glue schema registry configuration.
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
//   		SchemaValidationConfigs: []kafkaSchemaValidationConfig{
//   			&kafkaSchemaValidationConfig{
//   				Attribute: lambda.KafkaSchemaValidationAttribute_KEY(),
//   			},
//   		},
//   	}),
//   }))
//
type GlueSchemaRegistryProps struct {
	// The record format that Lambda delivers to your function after schema validation.
	//
	// - Choose JSON to have Lambda deliver the record to your function as a standard JSON object.
	//  - Choose SOURCE to have Lambda deliver the record to your function in its original source format. Lambda removes all schema metadata, such as the schema ID, before sending the record to your function.
	// Default: - none.
	//
	EventRecordFormat awslambda.EventRecordFormat `field:"required" json:"eventRecordFormat" yaml:"eventRecordFormat"`
	// An array of schema validation configuration objects, which tell Lambda the message attributes you want to validate and filter using your schema registry.
	// Default: - none.
	//
	SchemaValidationConfigs *[]*awslambda.KafkaSchemaValidationConfig `field:"required" json:"schemaValidationConfigs" yaml:"schemaValidationConfigs"`
	// The CfnRegistry reference of your glue schema registry.
	//
	// If used, schemaRegistryArn will be ignored.
	// Default: - none.
	//
	SchemaRegistry awsglue.CfnRegistry `field:"optional" json:"schemaRegistry" yaml:"schemaRegistry"`
	// The Arn of your glue schema registry.
	// Default: - none.
	//
	SchemaRegistryArn *string `field:"optional" json:"schemaRegistryArn" yaml:"schemaRegistryArn"`
}

