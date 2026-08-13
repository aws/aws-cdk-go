package awsbedrockagentcore


// Content configuration for a stream delivery resource.
//
// Defines what content type and detail level to deliver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamDeliveryContentConfiguration := &StreamDeliveryContentConfiguration{
//   	Level: awscdk.Aws_bedrockagentcore.StreamDeliveryContentLevel_METADATA_ONLY,
//   	Type: awscdk.*Aws_bedrockagentcore.StreamDeliveryContentType_MEMORY_RECORDS,
//   }
//
type StreamDeliveryContentConfiguration struct {
	// The level of content detail to deliver.
	//
	// There is no default: the level must be chosen explicitly because
	// `FULL_CONTENT` delivers complete memory record bodies, which can contain
	// personally identifiable information and other sensitive conversation
	// content. Use `METADATA_ONLY` unless the record body is required downstream.
	Level StreamDeliveryContentLevel `field:"required" json:"level" yaml:"level"`
	// The type of content to deliver.
	Type StreamDeliveryContentType `field:"required" json:"type" yaml:"type"`
}

