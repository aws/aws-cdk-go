package awslambda


// Specific schema validation configuration settings that tell Lambda the message attributes you want to validate and filter using your schema registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var kafkaSchemaValidationAttribute KafkaSchemaValidationAttribute
//
//   kafkaSchemaValidationConfig := &KafkaSchemaValidationConfig{
//   	Attribute: kafkaSchemaValidationAttribute,
//   }
//
type KafkaSchemaValidationConfig struct {
	// The attributes you want your schema registry to validate and filter for.
	//
	// If you selected JSON as the EventRecordFormat, Lambda also deserializes the selected message attributes.
	Attribute KafkaSchemaValidationAttribute `field:"required" json:"attribute" yaml:"attribute"`
}

