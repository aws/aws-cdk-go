package awslambda


// (Amazon MSK and self-managed Apache Kafka only) Specific configuration settings for a Kafka schema registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventRecordFormat EventRecordFormat
//   var kafkaSchemaRegistryAccessConfigType KafkaSchemaRegistryAccessConfigType
//   var kafkaSchemaValidationAttribute KafkaSchemaValidationAttribute
//
//   kafkaSchemaRegistryConfig := &KafkaSchemaRegistryConfig{
//   	EventRecordFormat: eventRecordFormat,
//   	SchemaRegistryUri: jsii.String("schemaRegistryUri"),
//   	SchemaValidationConfigs: []KafkaSchemaValidationConfig{
//   		&KafkaSchemaValidationConfig{
//   			Attribute: kafkaSchemaValidationAttribute,
//   		},
//   	},
//
//   	// the properties below are optional
//   	AccessConfigs: []KafkaSchemaRegistryAccessConfig{
//   		&KafkaSchemaRegistryAccessConfig{
//   			Type: kafkaSchemaRegistryAccessConfigType,
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   }
//
type KafkaSchemaRegistryConfig struct {
	// The record format that Lambda delivers to your function after schema validation.
	//
	// - Choose JSON to have Lambda deliver the record to your function as a standard JSON object.
	//  - Choose SOURCE to have Lambda deliver the record to your function in its original source format. Lambda removes all schema metadata, such as the schema ID, before sending the record to your function.
	// Default: - none.
	//
	EventRecordFormat EventRecordFormat `field:"required" json:"eventRecordFormat" yaml:"eventRecordFormat"`
	// The URI for your schema registry.
	//
	// The correct URI format depends on the type of schema registry you're using.
	// Default: - none.
	//
	SchemaRegistryUri *string `field:"required" json:"schemaRegistryUri" yaml:"schemaRegistryUri"`
	// An array of schema validation configuration objects, which tell Lambda the message attributes you want to validate and filter using your schema registry.
	// Default: - none.
	//
	SchemaValidationConfigs *[]*KafkaSchemaValidationConfig `field:"required" json:"schemaValidationConfigs" yaml:"schemaValidationConfigs"`
	// An array of access configuration objects that tell Lambda how to authenticate with your schema registry.
	// Default: - none.
	//
	AccessConfigs *[]*KafkaSchemaRegistryAccessConfig `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
}

