package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRegistry`.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-registry.html
//
type CfnRegistryProps struct {
	// The name of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-registry.html#cfn-glue-registry-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-registry.html#cfn-glue-registry-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-registry.html#cfn-glue-registry-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

