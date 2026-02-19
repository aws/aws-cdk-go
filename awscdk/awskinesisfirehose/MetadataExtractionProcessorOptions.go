package awskinesisfirehose


// Props for MetadataExtractionProcessor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonParsingEngine JsonParsingEngine
//
//   metadataExtractionProcessorOptions := &MetadataExtractionProcessorOptions{
//   	JsonParsingEngine: jsonParsingEngine,
//   	MetadataExtractionQuery: jsii.String("metadataExtractionQuery"),
//   }
//
type MetadataExtractionProcessorOptions struct {
	// JSON parsing engine.
	JsonParsingEngine JsonParsingEngine `field:"required" json:"jsonParsingEngine" yaml:"jsonParsingEngine"`
	// Map parameter to JQ query.
	MetadataExtractionQuery *string `field:"required" json:"metadataExtractionQuery" yaml:"metadataExtractionQuery"`
}

