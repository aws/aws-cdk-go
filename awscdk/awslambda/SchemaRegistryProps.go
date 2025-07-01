package awslambda


// Properties for schema registry configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventRecordFormat eventRecordFormat
//   var kafkaSchemaValidationAttribute kafkaSchemaValidationAttribute
//
//   schemaRegistryProps := &SchemaRegistryProps{
//   	EventRecordFormat: eventRecordFormat,
//   	SchemaValidationConfigs: []kafkaSchemaValidationConfig{
//   		&kafkaSchemaValidationConfig{
//   			Attribute: kafkaSchemaValidationAttribute,
//   		},
//   	},
//   }
//
type SchemaRegistryProps struct {
	// The record format that Lambda delivers to your function after schema validation.
	//
	// - Choose JSON to have Lambda deliver the record to your function as a standard JSON object.
	//  - Choose SOURCE to have Lambda deliver the record to your function in its original source format. Lambda removes all schema metadata, such as the schema ID, before sending the record to your function.
	// Default: - none.
	//
	EventRecordFormat EventRecordFormat `field:"required" json:"eventRecordFormat" yaml:"eventRecordFormat"`
	// An array of schema validation configuration objects, which tell Lambda the message attributes you want to validate and filter using your schema registry.
	// Default: - none.
	//
	SchemaValidationConfigs *[]*KafkaSchemaValidationConfig `field:"required" json:"schemaValidationConfigs" yaml:"schemaValidationConfigs"`
}

