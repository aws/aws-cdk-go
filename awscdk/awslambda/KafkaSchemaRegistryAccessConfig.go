package awslambda


// Specific access configuration settings that tell Lambda how to authenticate with your schema registry.
//
// If you're working with an AWS Glue schema registry, don't provide authentication details in this object. Instead, ensure that your execution role has the required permissions for Lambda to access your cluster.
//
// If you're working with a Confluent schema registry, choose the authentication method in the Type field, and provide the AWS Secrets Manager secret ARN in the URI field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var kafkaSchemaRegistryAccessConfigType kafkaSchemaRegistryAccessConfigType
//
//   kafkaSchemaRegistryAccessConfig := &KafkaSchemaRegistryAccessConfig{
//   	Type: kafkaSchemaRegistryAccessConfigType,
//   	Uri: jsii.String("uri"),
//   }
//
type KafkaSchemaRegistryAccessConfig struct {
	// The type of authentication Lambda uses to access your schema registry.
	Type KafkaSchemaRegistryAccessConfigType `field:"required" json:"type" yaml:"type"`
	// The URI of the secret (Secrets Manager secret ARN) to authenticate with your schema registry.
	// See: KafkaSchemaRegistryAccessConfigType.
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

