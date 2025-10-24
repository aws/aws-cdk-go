package awslambdaeventsources


// (Amazon MSK and self-managed Apache Kafka only) The provisioned mode configuration for the event source.
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
type ProvisionedPollerConfig struct {
	// The maximum number of pollers that can be provisioned.
	// Default: 200.
	//
	MaximumPollers *float64 `field:"required" json:"maximumPollers" yaml:"maximumPollers"`
	// The minimum number of pollers that should be provisioned.
	// Default: 1.
	//
	MinimumPollers *float64 `field:"required" json:"minimumPollers" yaml:"minimumPollers"`
}

