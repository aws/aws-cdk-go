package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for confluent schema registry configuration.
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
type ConfluentSchemaRegistryProps struct {
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
	// The type of authentication for schema registry credentials.
	// Default: none.
	//
	AuthenticationType awslambda.KafkaSchemaRegistryAccessConfigType `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The URI for your schema registry.
	// Default: - none.
	//
	SchemaRegistryUri *string `field:"required" json:"schemaRegistryUri" yaml:"schemaRegistryUri"`
	// The secret with the schema registry credentials.
	// Default: none.
	//
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
}

